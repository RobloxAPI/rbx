package types

import (
	"github.com/robloxapi/rbx"
)

// NilConnection implements an rbx.Instance, but represents the non-presence of a
// value.
const NilConnection = nil_connection

// Some indirection to hide the implementation.
const nil_connection = _NilConnection(false)

type _NilConnection bool

func (n _NilConnection) Type() string {
	return "RBXScriptConnection"
}
func (n _NilConnection) String() string {
	return "nil"
}
func (n _NilConnection) Copy() rbx.Value {
	return n
}
func (n _NilConnection) Connected() bool {
	return false
}
func (n _NilConnection) Disconnect() {}
