package types

import (
	"github.com/robloxapi/rbx"
)

type String string

func (String) Type() string {
	return "string"
}
func (s String) String() string {
	return string(s)
}
func (s String) Copy() rbx.Value {
	return s
}
