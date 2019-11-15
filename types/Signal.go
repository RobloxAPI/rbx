package types

import (
	"github.com/robloxapi/rbx"
)

// NilSignal implements an rbx.Signal, but represents the non-presence of a
// value.
const NilSignal = nil_signal

// Some indirection to hide the implementation.
const nil_signal = _NilSignal(false)

type _NilSignal bool

func (n _NilSignal) Type() string {
	return "RBXScriptSignal"
}
func (n _NilSignal) String() string {
	return "nil"
}
func (n _NilSignal) Copy() rbx.Value {
	return n
}
func (n _NilSignal) Connect(Function) rbx.Connection {
	return nil
}
func (n _NilSignal) Wait() []rbx.Value {
	return nil
}
