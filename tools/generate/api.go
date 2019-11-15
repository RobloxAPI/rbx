package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/anaminus/but"
	"github.com/dave/jennifer/jen"
	"github.com/robloxapi/rbx/desc"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxapi/rbxapijson"
	"github.com/robloxapi/rbxdhist"
)

type API struct {
	Root    *rbxapijson.Root
	Version rbxdhist.Version
}

func FetchAPI(ver, dump string) (api API) {
	var v struct {
		Version             rbxdhist.Version `json:"version"`
		ClientVersionUpload string           `json:"clientVersionUpload"`
		BootstrapperVersion string           `json:"bootstrapperVersion"`
		Errors
	}

	resp, err := http.Get(ver)
	but.IfFatal(err, "get version")
	checkStatus(resp, "get version")
	err = json.NewDecoder(resp.Body).Decode(&v)
	resp.Body.Close()
	but.IfFatal(err, "decode version")
	but.IfFatal(v.Unwrap(), "version response")

	u := fmt.Sprintf(dump, v.ClientVersionUpload)
	resp, err = http.Get(u)
	but.IfFatal(err, u, ": get dump")
	checkStatus(resp, u, ": get dump")

	api.Root, err = rbxapijson.Decode(resp.Body)
	resp.Body.Close()
	but.IfFatal(err, "decode dump")

	api.Version = v.Version
	return api
}

func (api API) GenerateVersion(file string) {
	f, err := os.Create(file)
	but.IfFatal(err, "create version file")
	defer f.Close()

	buf := bufio.NewWriter(f)
	buf.WriteString(Preamble)
	fmt.Fprintf(buf, "package rbx\n\nconst Version = %q\n", api.Version.String())
	but.IfFatal(buf.Flush(), "write version file")
	but.IfFatal(f.Sync(), "sync version file")
}

func (api API) filterType(t rbxapijson.Type) string {
	switch t.Category {
	case "Class":
		return "rbx.Instance"
	case "Enum":
		return "enum." + t.Name
	}
	switch t.Name {
	case "bool":
		return "types.Bool"
	case "RBXScriptConnection":
		return "rbx.Connection"
	case "double":
		return "types.Double"
	case "float":
		return "types.Float"
	case "function":
		return "types.Function"
	case "Objects":
		return "[]rbx.Instance"
	case "int":
		return "types.Int"
	case "int64":
		return "types.Int64"
	case "RBXScriptSignal":
		return "rbx.Signal"
	case "string":
		return "types.String"
	}
	return "types." + t.Name
}
func (api API) filterDefault(t rbxapijson.Type, s string) string {
	switch t.Category {
	case "Enum":
		return "(0)"
	}
	switch s {
	case "types.Array":
		return "{}"
	case "types.Axes":
		return "{}"
	case "types.BinaryString":
		return "(\"\")"
	case "types.Bool":
		return "(false)"
	case "types.BrickColor":
		return "(0)"
	case "types.CFrame":
		return "{}"
	case "types.Color3":
		return "{}"
	case "types.ColorSequence":
		return "{}"
	case "types.ColorSequenceKeypoint":
		return "{}"
	case "rbx.Connection":
		return "(nil)"
	case "types.Content":
		return "(\"\")"
	case "types.Dictionary":
		return "{}"
	case "types.DockWidgetPluginGuiInfo":
		return "{}"
	case "types.Double":
		return "(0)"
	case "rbx.Enum":
		return "(nil)"
	case "rbx.EnumItem":
		return "(nil)"
	case "types.Faces":
		return "{}"
	case "types.Float":
		return "(0)"
	case "types.Function":
		return "(nil)"
	case "rbx.Instance":
		return "(nil)"
	case "types.Int":
		return "(0)"
	case "types.Int64":
		return "(0)"
	case "types.NumberRange":
		return "{}"
	case "types.NumberSequence":
		return "{}"
	case "types.NumberSequenceKeypoint":
		return "{}"
	case "types.PathWaypoint":
		return "{}"
	case "types.PhysicalProperties":
		return "{}"
	case "types.ProtectedString":
		return "(\"\")"
	case "types.QDir":
		return "(\"\")"
	case "types.QFont":
		return "(\"\")"
	case "types.Ray":
		return "{}"
	case "types.Rect":
		return "{}"
	case "types.Region3":
		return "{}"
	case "types.Region3int16":
		return "{}"
	case "rbx.Signal":
		return "(nil)"
	case "types.String":
		return "(\"\")"
	case "types.TweenInfo":
		return "{}"
	case "types.UDim":
		return "{}"
	case "types.UDim2":
		return "{}"
	case "types.Vector2":
		return "{}"
	case "types.Vector2int16":
		return "{}"
	case "types.Vector3":
		return "{}"
	case "types.Vector3int16":
		return "{}"
	case "[]rbx.Instance":
		return "{}"
	}
	return "(nil)"
}

const (
	rbxdesc  = "github.com/robloxapi/rbx/desc"
	rbxtypes = "github.com/robloxapi/rbx/types"
	rbxenum  = "github.com/robloxapi/rbx/enum"
)

func formatTypeString(t rbxapi.Type) jen.Code {
	switch t.GetCategory() {
	case "Class", "Enum":
		return jen.Lit(t.GetCategory() + "." + t.GetName())
	}
	return jen.Lit(t.GetName())
}

func formatType(root *rbxapijson.Root, t rbxapi.Type, def string, hasdef bool) jen.Code {
	switch t.GetCategory() {
	case "Enum":
		if !hasdef {
			return jen.Op("*").Id("new").Call(jen.Qual(rbxenum, t.GetName()))
		}
		for _, enum := range root.Enums {
			if enum.Name != t.GetName() {
				continue
			}
			for _, item := range enum.Items {
				if item.Name != def {
					continue
				}
				return jen.Qual(rbxenum, t.GetName()).Call(jen.Lit(item.Value))
			}
		}
		return jen.Qual(rbxenum, t.GetName()).Op(".").Id("Enum").Call(jen.Lit(0)).Op(".").Id("Item").Call(jen.Lit(def))
	case "Class":
		return jen.Qual(rbxtypes, "NilInstance")
	}
	typ := t.GetName()
	if !hasdef {
		switch typ {
		case "bool", "double", "float", "int", "int64", "string":
			typ = strings.Title(typ)
		case "void":
			return jen.Nil()
		case "RBXScriptSignal":
			return jen.Qual(rbxtypes, "NilSignal")
		case "RBXScriptConnection":
			return jen.Qual(rbxtypes, "NilConnection")
		}
		return jen.Op("*").Id("new").Call(jen.Qual(rbxtypes, typ))
	}
	switch typ {
	case "void":
		return jen.Nil()
	case "BinaryString", "Content", "ProtectedString", "QDir", "QFont", "string":
		return jen.Qual(rbxtypes, "String").Call(jen.Lit(def))
	case "bool":
		if def == "true" {
			return jen.Qual(rbxtypes, "Bool").Call(jen.Lit(true))
		}
		return jen.Qual(rbxtypes, "Bool").Call(jen.Lit(false))
	case "CFrame":
		// def == "Identity"
		return jen.Qual(rbxtypes, "NewCFrame").Call()
	case "double":
		n, _ := strconv.ParseFloat(def, 64)
		return jen.Qual(rbxtypes, "Double").Call(jen.Lit(n))
	case "float":
		n, _ := strconv.ParseFloat(def, 32)
		return jen.Qual(rbxtypes, "Float").Call(jen.Lit(float32(n)))
	case "int":
		n, _ := strconv.ParseInt(def, 10, 32)
		return jen.Qual(rbxtypes, "Int").Call(jen.Lit(int32(n)))
	case "int64":
		n, _ := strconv.ParseInt(def, 10, 64)
		return jen.Qual(rbxtypes, "Int64").Call(jen.Lit(int64(n)))
	case "RBXScriptSignal":
		return jen.Qual(rbxtypes, "NilSignal")
	case "RBXScriptConnection":
		return jen.Qual(rbxtypes, "NilConnection")
	default:
		return jen.Op("*").Id("new").Call(jen.Qual(rbxtypes, typ))
	}
}

func (api API) GenerateClasses(file string) {
	j := jen.NewFile("class")
	j.HeaderComment(strings.TrimSpace(Preamble))
	j.ImportName(rbxdesc, "desc")
	j.ImportName(rbxtypes, "types")
	j.ImportName(rbxenum, "enum")

	var list []jen.Code
	for _, class := range api.Root.Classes {
		list = append(list, jen.Line().Id(class.Name))
	}
	list = append(list, jen.Line())
	j.Var().Id("list").Op("=").Index().Op("*").Qual(rbxdesc, "Class").Values(list...).Line()

	dict := jen.Dict{}
	for _, class := range api.Root.Classes {
		dict[jen.Lit(class.Name)] = jen.Id(class.Name)
	}
	j.Var().Id("classes").Op("=").Map(jen.String()).Op("*").Qual(rbxdesc, "Class").Values(dict).Line()

	for _, class := range api.Root.Classes {
		var params []jen.Code
		params = append(params, jen.Lit(class.Name))
		s := class.Superclass
		if s == "" || s == "<<<ROOT>>>" {
			s = "_Root"
		}
		params = append(params, jen.Id(s))
		var flags = &jen.Statement{}
		for _, tag := range class.GetTags() {
			if desc.FlagsFromString(tag) != 0 {
				if len(*flags) > 0 {
					flags.Op("|")
				}
				flags.Qual(rbxdesc, tag)
			}
		}
		if len(*flags) > 0 {
			params = append(params, jen.Line().Add(flags))
		}
		for _, member := range class.Members {
			fields := jen.Dict{}

			// Name.
			fields[jen.Id("Name")] = jen.Lit(member.GetName())

			// Flags.
			var flags = &jen.Statement{}
			for _, tag := range member.GetTags() {
				if desc.FlagsFromString(tag) != 0 {
					if len(*flags) > 0 {
						flags.Op("|")
					}
					flags.Qual(rbxdesc, tag)
				}
			}

			// Parameters.
			if member, ok := member.(interface{ GetParameters() rbxapi.Parameters }); ok {
				var params []jen.Code
				for _, param := range member.GetParameters().GetParameters() {
					f := jen.Dict{
						jen.Id("Name"): jen.Lit(param.GetName()),
					}
					if def, ok := param.GetDefault(); ok {
						f[jen.Id("Optional")] = jen.Lit(true)
						f[jen.Id("Type")] = formatType(api.Root, param.GetType(), def, true)
					} else {
						f[jen.Id("Type")] = formatType(api.Root, param.GetType(), "", false)
					}
					params = append(params, jen.Line().Values(f))
				}
				if len(params) > 0 {
					params = append(params, jen.Line())
					fields[jen.Id("Parameters")] = jen.Index().Qual(rbxdesc, "Parameter").Values(params...)
				}
			}

			// ReturnType.
			if member, ok := member.(interface{ GetReturnType() rbxapi.Type }); ok {
				if ret := member.GetReturnType(); ret.GetName() != "void" {
					fields[jen.Id("Returns")] = formatType(api.Root, ret, "", false)
				}
			}

			// Security.
			if member, ok := member.(interface{ GetSecurity() string }); ok {
				if s := member.GetSecurity(); s != "" && s != "None" {
					fields[jen.Id("Security")] = jen.Qual(rbxdesc, s)
				}
			}

			// Member-specific.
			switch member := member.(type) {
			case *rbxapijson.Property:
				fields[jen.Id("Value")] = formatType(api.Root, member.ValueType, "", false)
				if s := member.ReadSecurity; s != "" && s != "None" {
					fields[jen.Id("ReadSecurity")] = jen.Qual(rbxdesc, s)
				}
				if s := member.WriteSecurity; s != "" && s != "None" {
					fields[jen.Id("WriteSecurity")] = jen.Qual(rbxdesc, s)
				}
				if member.CanLoad {
					if len(*flags) > 0 {
						flags.Op("|")
					}
					flags.Qual(rbxdesc, "CanLoad")
				}
				if member.CanSave {
					if len(*flags) > 0 {
						flags.Op("|")
					}
					flags.Qual(rbxdesc, "CanSave")
				}
			case *rbxapijson.Function:
			case *rbxapijson.Event:
			case *rbxapijson.Callback:
			}

			// Finish flags.
			if len(*flags) > 0 {
				fields[jen.Id("Flags")] = flags
			}

			params = append(params, jen.Line().Qual(rbxdesc, member.GetMemberType()).Values(fields))
		}

		if len(params) > 1 {
			params = append(params, jen.Line())
		}
		j.Var().Id(class.Name).Op("=").Qual(rbxdesc, "NewClass").Call(params...).Line()
	}

	f, err := os.Create(file)
	but.IfFatal(err, "create classes file")
	defer f.Close()
	buf := bufio.NewWriter(f)
	but.IfFatal(j.Render(buf), "render")
	but.IfFatal(buf.Flush(), "write classes file")
	but.IfFatal(f.Sync(), "sync classes file")
}

func (api API) GenerateRegisterClasses(file string) {
	f, err := os.Create(file)
	but.IfFatal(err, "create class registration file")
	defer f.Close()

	buf := bufio.NewWriter(f)
	buf.WriteString(Preamble)

	buf.WriteString("package rbx\n\nimport \"github.com/robloxapi/rbx/class\"\n\nfunc init() {\n")
	for _, class := range api.Root.Classes {
		fmt.Fprintf(buf, "\tclasses[%[1]q] = class.%[1]s\n", class.Name)
	}
	buf.WriteString("}\n")
	but.IfFatal(buf.Flush(), "write class registration file")
	but.IfFatal(f.Sync(), "sync class registration file")
}

func (api API) GenerateEnums(file string) {
	f, err := os.Create(file)
	but.IfFatal(err, "create enums file")
	defer f.Close()

	buf := bufio.NewWriter(f)
	buf.WriteString(Preamble)
	buf.WriteString("package enum\n\nimport (\n\t\"github.com/robloxapi/rbx\"\n)\n")
	buf.WriteString("\nvar list = []rbx.Enum{\n")
	var max int
	for _, enum := range api.Root.Enums {
		fmt.Fprintf(buf, "\t_%s{},\n", enum.Name)
		if n := len(enum.Name); n > max {
			max = n
		}
	}
	buf.WriteString("}\n")

	buf.WriteString("\nvar enums = map[string]rbx.Enum{\n")
	for _, enum := range api.Root.Enums {
		fmt.Fprintf(buf, "\t%*s _%s{},\n", -max-3, "\""+enum.Name+"\":", enum.Name)
	}
	buf.WriteString("}\n")

	tmpl := template.Must(template.New("").Parse(`
////////////////////////////////////////////////////////////////////////////////

type _{{.Name}} struct{}

func (e _{{.Name}}) Type() string {
	return "Enum"
}
func (e _{{.Name}}) String() string {
	return "Enum.{{.Name}}"
}
func (e _{{.Name}}) Copy() rbx.Value {
	return e
}
func (e _{{.Name}}) Name() string {
	return "{{.Name}}"
}
func (e _{{.Name}}) Items() []rbx.EnumItem {
	return []rbx.EnumItem{
{{- $enum := .}}
{{- range .Items}}
		{{$enum.Name}}({{.Value}}),
{{- end}}
	}
}
func (e _{{.Name}}) Item(name string) rbx.EnumItem {
	switch name {
{{- range .Items}}
	case "{{.Name}}":
		return {{$enum.Name}}({{.Value}})
{{- end}}
	}
	return nil
}

type {{.Name}} uint

func (e {{.Name}}) Type() string {
	return "Enum.{{.Name}}"
}
func (e {{.Name}}) String() string {
	return "Enum.{{.Name}}." + e.Name()
}
func (e {{.Name}}) Copy() rbx.Value {
	return e
}
func (e {{.Name}}) Enum() rbx.Enum {
	return _{{.Name}}{}
}
func (e {{.Name}}) Value() int {
	return int(e)
}
func (e {{.Name}}) Name() string {
	switch e {
{{- range .FItems}}
	case {{.Value}}:
		return "{{.Name}}"
{{- end}}
	}
	return ""
}
`))
	// Filter items by value; a duplicate value selects the latest item.
	type Enum struct {
		Name   string
		Items  []*rbxapijson.EnumItem
		FItems []*rbxapijson.EnumItem
	}
	enums := make([]Enum, len(api.Root.Enums))
	for i, apiEnum := range api.Root.Enums {
		enum := Enum{
			Name:   apiEnum.Name,
			Items:  apiEnum.Items,
			FItems: make([]*rbxapijson.EnumItem, 0, len(apiEnum.Items)),
		}
	loop:
		for i, item := range apiEnum.Items {
			for j := i + 1; j < len(apiEnum.Items); j++ {
				if apiEnum.Items[j].Value == item.Value {
					continue loop
				}
			}
			enum.FItems = append(enum.FItems, item)
		}
		enums[i] = enum
	}
	for _, enum := range enums {
		tmpl.Execute(buf, enum)
	}
	but.IfFatal(buf.Flush(), "write enums file")
	but.IfFatal(f.Sync(), "sync enums file")
}

func (api API) GenerateRegisterEnums(file string) {
	f, err := os.Create(file)
	but.IfFatal(err, "create enum registration file")
	defer f.Close()

	buf := bufio.NewWriter(f)
	buf.WriteString(Preamble)
	buf.WriteString("package rbx\n\nimport \"github.com/robloxapi/rbx/enum\"\n\nfunc init() {\n")
	for _, enum := range api.Root.Enums {
		fmt.Fprintf(buf, "\tenums[%[1]q] = enum.%[1]s(0).Enum()\n", enum.Name)
	}
	buf.WriteString("}\n")
	but.IfFatal(buf.Flush(), "write enum registration file")
	but.IfFatal(f.Sync(), "sync enum registration file")
}
