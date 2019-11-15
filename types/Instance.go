package types

import (
	"github.com/robloxapi/rbx"
)

// NilInstance implements an rbx.Instance, but represents the non-presence of a
// value.
const NilInstance = nil_instance

// Some indirection to hide the implementation.
const nil_instance = _NilInstance(false)

type _NilInstance bool

func (n _NilInstance) Type() string {
	return "Instance"
}
func (n _NilInstance) String() string {
	return "nil"
}
func (n _NilInstance) Copy() rbx.Value {
	return n
}
func (n _NilInstance) Get(string) (rbx.Value, error) {
	return nil, nil
}
func (n _NilInstance) Set(string, rbx.Value) error {
	return nil
}
func (n _NilInstance) ClassName() string {
	return "nil"
}
func (n _NilInstance) IsA(string) bool {
	return false
}
