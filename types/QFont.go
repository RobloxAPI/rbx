package types

import (
	"github.com/robloxapi/rbx"
)

type QFont string

func (QFont) Type() string {
	return "QFont"
}
func (s QFont) String() string {
	return string(s)
}
func (s QFont) Copy() rbx.Value {
	return s
}
