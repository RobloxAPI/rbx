package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/anaminus/but"
)

func main() {
	api := FetchAPI(
		`https://clientsettings.roblox.com/v1/client-version/WindowsStudio64`,
		`https://setup.rbxcdn.com/%s-API-Dump.json`,
	)
	api.GenerateVersion(`version_gen.go`)
	api.GenerateClasses(`class/classes_gen.go`)
	// api.GenerateRegisterClasses(`classes_gen.go`)
	api.GenerateEnums(`enum/enums_gen.go`)
	// api.GenerateRegisterEnums(`enums_gen.go`)
	GenerateBrickColors("tools/generate/brickcolors.json", "types/BrickColor_gen.go")
}

const Preamble = "// This file was generated by tools/generate. DO NOT EDIT!\n\n"

type Error struct {
	Code              int    `json:"code"`
	Message           string `json:"message"`
	UserFacingMessage string `json:"userFacingMessage"`
}

func (err Error) Error() string {
	return fmt.Sprintf("%d: %s", err.Code, err.Message)
}

type Errors struct {
	Errors []Error `json:"errors"`
}

func (err Errors) Error() string {
	if len(err.Errors) == 0 {
		return "an error occurred"
	}
	return err.Errors[0].Error()
}

func (err Errors) Unwrap() error {
	if len(err.Errors) == 0 {
		return nil
	}
	return err.Errors[0]
}

func checkStatus(resp *http.Response, args ...interface{}) {
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if resp.Body != nil {
			resp.Body.Close()
		}
		but.IfFatal(errors.New(resp.Status), args...)
	}
}

func wrapList(buf *strings.Builder, wrap int, length int, cb func(*strings.Builder, int) bool) {
	if wrap <= 0 {
		wrap = 1
	}
	var b strings.Builder
	for i, j := 0, 0; i < length; i, j = i+1, j+1 {
		b.Reset()
		if !cb(&b, i) {
			continue
		}
		if j%wrap == 0 {
			buf.WriteString("\n\t")
			buf.WriteString(b.String())
			buf.WriteString(",")
		} else {
			buf.WriteString(" ")
			buf.WriteString(b.String())
			buf.WriteString(",")
		}
	}
	buf.WriteString("\n")
}

func wrapString(wrap, indent int, typ string, b []byte) string {
	if typ == "string" {
		typ = ""
	}
	if len(b) == 0 {
		if typ == "" {
			return "\"\"\n"
		}
		return typ + "(\"\")\n"
	}
	var s strings.Builder
	if typ != "" {
		s.WriteString(typ)
		s.WriteByte('(')
	}
	if wrap > 0 && len(b) > wrap {
		s.WriteString("\"\" +\n")
	} else {
		s.WriteByte('"')
	}
	for i := 0; i < len(b); i++ {
		if len(b) > wrap && wrap > 0 && i%wrap == 0 {
			for i := 0; i < indent; i++ {
				s.WriteString("\t")
			}
			s.WriteString("\"")
		}
		s.WriteByte(b[i])
		if i == len(b)-1 {
			s.WriteByte('"')
			if typ != "" {
				s.WriteByte(')')
			}
			s.WriteByte('\n')
		} else if wrap > 0 && i%wrap == wrap-1 {
			s.WriteString("\" +\n")
		}
	}
	return s.String()
}
