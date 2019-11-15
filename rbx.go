// The rbx package is a partial implementation of the Roblox Lua API.
package rbx

//go:generate tools/generate/generate

// A Value is any value that may be passed into a Lua environment.
type Value interface {
	// Type returns the Type that the Value implements.
	Type() string
	// String returns a string describing the value.
	String() string
	// Copy returns a deep copy of the value.
	Copy() Value
}

type Instance interface {
	Proper
	ClassName() string
	IsA(string) bool
}

// Proper is any value that is capable of having properties.
type Proper interface {
	Value
	Get(string) (Value, error)
	Set(string, Value) error
}

// Methoder is any value that is capable of having methods.
type Methoder interface {
	Value
	Method(string) (Function, error)
}

// Signaler is any value that is capable of having signals.
type Signaler interface {
	Value
	Signal(string) (Signal, error)
}

type Enum interface {
	Value
	// Name returns the name of the enum.
	Name() string
	// Item returns the item corresponding to the given name, or nil if no such
	// item exists.
	Item(name string) EnumItem
	// Items returns a list of the enum's items.
	Items() []EnumItem
}

type EnumItem interface {
	Value
	// Enum returns the enum to which the item belongs.
	Enum() Enum
	// Name returns the name of the item.
	Name() string
	// Value returns the value of the item.
	Value() int
}

// Event is implemented by types that fire a Signal.
type Event interface {
	// Signal returns the Signal that is fired by the event.
	Signal() Signal
	// Fire invokes the Signal.
	Fire(parameters ...Value)
}

type Signal interface {
	Value
	Connect(Function) Connection
	Wait() []Value
}

type Connection interface {
	Value
	Connected() bool
	Disconnect()
}

// Function represents a function that can receive and return Values. If the
// function returns an error, then returns must be nil.
type Function func(parameters ...Value) (returns []Value, err error)
