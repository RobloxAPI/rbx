package types

import (
	"github.com/robloxapi/rbx"
)

type Bool bool

func (Bool) Type() string {
	return "bool"
}
func (b Bool) String() string {
	if b {
		return "true"
	}
	return "false"
}
func (b Bool) Copy() rbx.Value {
	return b
}
