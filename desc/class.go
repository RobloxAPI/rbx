package desc

import (
	"errors"

	"github.com/robloxapi/rbx"
)

// Flags represents a descriptor's flag bits.
type Flags uint

const (
	_ Flags = (1 << iota) / 2
	NotCreatable
	NotReplicated
	PlayerReplicated
	Service
	Settings
	Hidden
	ReadOnly
	CanLoad
	CanSave
	CanYield
	CustomLuaState
	Yields

	// DescriptorFlags are the Flags that apply to all descriptors.
	DescriptorFlags = 0
	// ClassFlags are the Flags that apply to Class descriptors.
	ClassFlags = DescriptorFlags | NotCreatable | NotReplicated | PlayerReplicated | Service | Settings
	// MemberFlags are the Flags that apply to all member descriptors.
	MemberFlags = DescriptorFlags
	// PropertyFlags are the Flags that apply to all Property descriptors.
	PropertyFlags = MemberFlags | NotReplicated | Hidden | ReadOnly | CanLoad | CanSave
	// FunctionFlags are the Flags that apply to all Function descriptors.
	FunctionFlags = MemberFlags | CanYield | CustomLuaState | Yields
	// EventFlags are the Flags that apply to all Event descriptors.
	EventFlags = MemberFlags
	// CallbackFlags are the Flags that apply to all Callback descriptors.
	CallbackFlags = MemberFlags
)

var flagNames = []string{
	"NotCreatable",
	"NotReplicated",
	"PlayerReplicated",
	"Service",
	"Settings",
	"Hidden",
	"ReadOnly",
	"CanLoad",
	"CanSave",
	"CanYield",
	"CustomLuaState",
	"Yields",
}

func FlagsFromString(s ...string) Flags {
	var f Flags
	for _, s := range s {
		switch s {
		case "NotCreatable":
			f |= NotCreatable
		case "NotReplicated":
			f |= NotReplicated
		case "PlayerReplicated":
			f |= PlayerReplicated
		case "Service":
			f |= Service
		case "Settings":
			f |= Settings
		case "Hidden":
			f |= Hidden
		case "ReadOnly":
			f |= ReadOnly
		case "CanLoad":
			f |= CanLoad
		case "CanSave":
			f |= CanSave
		case "CanYield":
			f |= CanYield
		case "CustomLuaState":
			f |= CustomLuaState
		case "Yields":
			f |= Yields
		}
	}
	return f
}

func (f Flags) String() string {
	var s string
	for i, name := range flagNames {
		if f&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += name
		}
	}
	if s == "" {
		s = "0"
	}
	return s
}

// Security indicates a context from which a member descriptor may be accessed.
type Security uint

const (
	NoSecurity Security = iota
	RobloxScriptSecurity
	LocalUserSecurity
	PluginSecurity
	RobloxSecurity
	NotAccessibleSecurity
	maxSecurity
)

func SecurityFromString(s string) Security {
	switch s {
	default:
		fallthrough
	case "NoSecurity", "None", "":
		return NoSecurity
	case "RobloxScriptSecurity":
		return RobloxScriptSecurity
	case "LocalUserSecurity":
		return LocalUserSecurity
	case "PluginSecurity":
		return PluginSecurity
	case "RobloxSecurity":
		return RobloxSecurity
	case "NotAccessibleSecurity":
		return NotAccessibleSecurity
	}
}

func (s Security) String() string {
	switch s {
	case NoSecurity:
		return "None"
	case RobloxScriptSecurity:
		return "RobloxScriptSecurity"
	case LocalUserSecurity:
		return "LocalUserSecurity"
	case PluginSecurity:
		return "PluginSecurity"
	case RobloxSecurity:
		return "RobloxSecurity"
	case NotAccessibleSecurity:
		return "NotAccessibleSecurity"
	}
	return "None"
}

// Alias marks a member called Name as an alias of a member called Of. Members
// of differing types are ignored. If one of the member does not exist, it is
// defined using the information from the other member as needed.
type Alias struct {
	Name   string
	Target string
}

// Class describes a Roblox class.
type Class struct {
	// Name is the name of the class.
	Name string
	// Flags are the flags associated with the descriptor, masked by ClassFlags.
	Flags Flags

	inited bool

	superclass *Class
	subclasses []*Class
	inheritMin int
	inheritMax int

	members   []Member
	memberMap map[string]Member

	aliases   map[string]string
	props     []*Property
	propIndex map[string]int
}

// Define sets an option on the class.
//
// When an option is a *Class, it will be set as the class's superclass. This
// can only be set once.
//
// When an option is a member descriptor (non-pointer), the member is appended
// to the class. When retrieving a member by name, the first member defined by
// the name is returned.
//
// When an option is an Alias, the member matching Alias.Name becomes an alias
// of the member matching Alias.Of. If the member types do not match, or either
// member does not exist, then the alias is ignored.
//
// When an option is a Flags, the set bits are set in the flags of the class,
// masked by ClassFlags.
//
// Other types are a no-op.
//
// Define panics if the class is initialized.
func (c *Class) Define(options ...Option) {
	if c.inited {
		panic("class is initialized")
	}

	for _, opt := range options {
		switch opt := opt.(type) {
		case *Class:
			if opt == nil {
				continue
			}
			if c.superclass == nil {
				c.superclass = opt
				opt.subclasses = append(opt.subclasses, c)
			}
		case Flags:
			c.Flags |= opt & ClassFlags
		case Alias:
			if c.aliases == nil {
				c.aliases = make(map[string]string)
			}
			c.aliases[opt.Name] = opt.Target
		case Property:
			opt.class = c
			opt.Flags &= PropertyFlags
			if c.memberMap == nil {
				c.memberMap = make(map[string]Member)
			}
			c.memberMap[opt.Name] = &opt
			c.members = append(c.members, &opt)
		case Function:
			opt.class = c
			opt.Flags &= FunctionFlags
			if c.memberMap == nil {
				c.memberMap = make(map[string]Member)
			}
			c.memberMap[opt.Name] = &opt
			c.members = append(c.members, &opt)
		case Event:
			opt.class = c
			opt.Flags &= EventFlags
			if c.memberMap == nil {
				c.memberMap = make(map[string]Member)
			}
			c.memberMap[opt.Name] = &opt
			c.members = append(c.members, &opt)
		case Callback:
			opt.class = c
			opt.Flags &= CallbackFlags
			if c.memberMap == nil {
				c.memberMap = make(map[string]Member)
			}
			c.memberMap[opt.Name] = &opt
			c.members = append(c.members, &opt)
		}
	}
}

func (c *Class) init(i int, ancestors []*Class) int {
	c.propIndex = make(map[string]int)
	for _, ancestor := range ancestors {
		// Build property tables.
		for _, member := range ancestor.members {
			prop, ok := member.(*Property)
			if !ok {
				continue
			}
			if i, ok := c.propIndex[prop.Name]; ok {
				// Override.
				c.props[i] = prop
				continue
			}
			if target, ok := ancestor.aliases[prop.Name]; ok {
				if _, ok = ancestor.memberMap[target].(*Property); ok {
					// Skip, for now.
					continue
				}
			}
			// Map property to index.
			c.propIndex[prop.Name] = len(c.props)
			c.props = append(c.props, prop)
		}
		for alias, target := range ancestor.aliases {
			if i, ok := c.propIndex[target]; ok {
				// Map alias to target.
				c.propIndex[alias] = i
				continue
			}
		}
	}
	c.inheritMin = i
	for _, sub := range c.subclasses {
		ancestors = append(ancestors, sub)
		i = sub.init(i+1, ancestors)
		ancestors = ancestors[:len(ancestors)-1]
	}
	c.inheritMax = i
	return c.inheritMax
}

// Init recursively initializes the class and all of it's descendant subclasses.
//
// Init rebuilds property lookup tables and inheritance index tables. Once
// initialized, classes cannot be modified.
func (c *Class) Init() {
	if c.inited {
		panic("class is already initialized")
	}
	c.inited = true

	inherited := make([]*Class, 1, 8)
	inherited[0] = c
	c.init(0, inherited)
}

func (c *Class) Superclass() *Class {
	return c.superclass
}

func (c *Class) Subclasses() []*Class {
	l := make([]*Class, len(c.subclasses))
	copy(l, c.subclasses)
	return l
}

// IsA returns whether the class is or inherits from the given class.
//
// The implementation assumes that the two classes are a part of the same
// inheritance tree. Results from separate trees are undefined.
func (c *Class) IsA(super *Class) bool {
	return super.inheritMin <= c.inheritMin && c.inheritMin <= super.inheritMax
}

func (c *Class) Member(name string) Member {
	return c.memberMap[name]
}
func (c *Class) Property(name string) *Property {
	m, _ := c.memberMap[name].(*Property)
	return m
}
func (c *Class) Function(name string) *Function {
	m, _ := c.memberMap[name].(*Function)
	return m
}
func (c *Class) Event(name string) *Event {
	m, _ := c.memberMap[name].(*Event)
	return m
}
func (c *Class) Callback(name string) *Callback {
	m, _ := c.memberMap[name].(*Callback)
	return m
}

// ErrNilValue indicates an attempt to set a property to a nil rbx.Value.
var ErrNilValue = errors.New("value is nil")

// ErrInvalidMember indicates that a member does not exist.
var ErrInvalidMember = errors.New("not a valid member")

// ErrReadOnly indicates that a property cannot be set because it is read-only.
var ErrReadOnly = errors.New("cannot set value")

// TypeMismatchError indicates a mismatch between an expected type and a given
// type.
type TypeMismatchError struct {
	Expected string
	Got      string
}

func (e TypeMismatchError) Error() string {
	return e.Expected + " expected, got " + e.Got
}

func (c *Class) GetProperty(list []rbx.Value, name string) (value rbx.Value, ok bool) {
	index, ok := c.propIndex[name]
	if !ok {
		return nil, false
	}
	return list[index], true
}

// SetProperty attempts to set the entry in list to value, indexed according to
// the mapping of name in the descriptor's property lookup table. May return
// ErrNilValue, ErrInvalidMember, ErrReadOnly, or TypeMismatchError.
func (c *Class) SetProperty(list []rbx.Value, name string, value rbx.Value) error {
	if value == nil {
		return ErrNilValue
	}

	index, ok := c.propIndex[name]
	if !ok {
		return ErrInvalidMember
	}
	desc := c.props[index]
	if desc.Flags&ReadOnly != 0 {
		return ErrReadOnly
	}
	if exp, got := desc.Value.Type(), value.Type(); exp != got {
		return TypeMismatchError{Expected: exp, Got: got}
	}
	list[index] = value
	return nil
}

func (c *Class) Members() []Member {
	l := make([]Member, len(c.members))
	copy(l, c.members)
	return l
}

func (c *Class) Properties() []*Property {
	n := 0
	for _, m := range c.members {
		if _, ok := m.(*Property); ok {
			n++
		}
	}
	l := make([]*Property, 0, n)
	for _, m := range c.members {
		if m, ok := m.(*Property); ok {
			l = append(l, m)
		}
	}
	return l
}

func (c *Class) Functions() []*Function {
	n := 0
	for _, m := range c.members {
		if _, ok := m.(*Function); ok {
			n++
		}
	}
	l := make([]*Function, 0, n)
	for _, m := range c.members {
		if m, ok := m.(*Function); ok {
			l = append(l, m)
		}
	}
	return l
}

func (c *Class) Events() []*Event {
	n := 0
	for _, m := range c.members {
		if _, ok := m.(*Event); ok {
			n++
		}
	}
	l := make([]*Event, 0, n)
	for _, m := range c.members {
		if m, ok := m.(*Event); ok {
			l = append(l, m)
		}
	}
	return l
}

func (c *Class) Callbacks() []*Callback {
	n := 0
	for _, m := range c.members {
		if _, ok := m.(*Callback); ok {
			n++
		}
	}
	l := make([]*Callback, 0, n)
	for _, m := range c.members {
		if m, ok := m.(*Callback); ok {
			l = append(l, m)
		}
	}
	return l
}

// Values returns a copy of the default property values of the class.
func (c *Class) Values() []rbx.Value {
	l := make([]rbx.Value, len(c.props))
	for i, prop := range c.props {
		l[i] = prop.Value.Copy()
	}
	return l
}

////////////////////////////////////////////////////////////////////////////////

// Member is implemented by any descriptor that is a member of a Class
// descriptor.
type Member interface {
	// member prevents external types from implementing this interface.
	member()
	// MemberType returns the kind of member.
	MemberType() string
	// Class returns the class descriptor to which the member belongs. Returns
	// nil if the member has not been associated with a class.
	Class() *Class
}

// Property describes a property member of a Class descriptor.
type Property struct {
	// Name is the name of the property.
	Name string
	// Value indicates the value type of the property, as well as the default
	// value. Must not be nil.
	Value rbx.Value
	// Flags are the flags associated with the descriptor, masked by
	// PropertyFlags.
	Flags Flags
	// ReadSecurity indicates the context from which the property may be read.
	ReadSecurity Security
	// WriteSecurity indicates the context from which the property member may be
	// written.
	WriteSecurity Security

	// Get is an optional getter implementation of the property.
	Get func() (rbx.Value, error)
	// Set is an optional setter implementation of the property.
	Set func(rbx.Value) error

	// class is the class descriptor to which the member belongs.
	class *Class
}

func (Property) member()            {}
func (Property) MemberType() string { return "Property" }
func (m *Property) Class() *Class   { return m.class }

// Parameter describes a parameter of a Function descriptor.
type Parameter struct {
	Name     string
	Optional bool
	Type     rbx.Value // If Optional, value is default.
}

// Function describes a method member of a Class descriptor.
type Function struct {
	// Name is the name of the function.
	Name string
	// Flags are the flags associated with the descriptor, masked by
	// FunctionFlags.
	Flags Flags
	// Security indicates the context from which the member may be accessed.
	Security Security
	// Parameters lists the values passed to the function.
	Parameters []Parameter
	// Returns is the value returned by the function. Nil indicates no value is
	// returned.
	Returns rbx.Value

	// Call is an optional implementation of the function.
	Call rbx.Function

	// Class is the class descriptor to which the member belongs.
	class *Class
}

func (Function) member()            {}
func (Function) MemberType() string { return "Function" }
func (m *Function) Class() *Class   { return m.class }

// Event describes an event member of a Class descriptor.
type Event struct {
	// Name is the name of the event.
	Name string
	// Flags are the flags associated with the descriptor, masked by EventFlags.
	Flags Flags
	// Security indicates the context from which the member may be accessed.
	Security Security
	// Parameters lists the values returned by the event.
	Parameters []Parameter

	// Event is an optional implementation of the event.
	Event rbx.Event

	// class is the class descriptor to which the member belongs.
	class *Class
}

func (Event) member()            {}
func (Event) MemberType() string { return "Event" }
func (m *Event) Class() *Class   { return m.class }

// Callback describes a callback member of a Class descriptor.
type Callback struct {
	// Name is the name of the callback.
	Name string
	// Flags are the flags associated with the descriptor, masked by
	// CallbackFlags.
	Flags Flags
	// Security indicates the context from which the member may be accessed.
	Security Security
	// Parameters lists the values passed to the function.
	Parameters []Parameter
	// Returns is the value returned by the function.
	Returns rbx.Value
	// Set is an optional implementation of the callback.
	Set func(rbx.Function) error

	// class is the class descriptor to which the member belongs.
	class *Class
}

func (Callback) member()            {}
func (Callback) MemberType() string { return "Callback" }
func (m *Callback) Class() *Class   { return m.class }

////////////////////////////////////////////////////////////////////////////////

// Option holds one of the option types: Property, Function, Event, Callback,
// Flags.
type Option interface{}

// NewClass returns a new class descriptor.
func NewClass(name string, options ...Option) *Class {
	class := Class{Name: name}
	class.Define(options...)
	return &class
}
