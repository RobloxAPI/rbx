package types

import (
	"github.com/robloxapi/rbx"
)

type Function rbx.Function

func (Function) Type() string {
	return "Function"
}
func (f Function) String() string {
	return "Function"
}
func (f Function) Copy() rbx.Value {
	return f
}
