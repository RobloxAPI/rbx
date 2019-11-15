package types

import (
	"github.com/robloxapi/rbx"
)

type ProtectedString string

func (ProtectedString) Type() string {
	return "ProtectedString"
}
func (p ProtectedString) String() string {
	return string(p)
}
func (p ProtectedString) Copy() rbx.Value {
	return p
}
