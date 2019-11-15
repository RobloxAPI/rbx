package types

import (
	"github.com/robloxapi/rbx"
)

type Content string

func (Content) Type() string {
	return "Content"
}
func (c Content) String() string {
	return string(c)
}
func (c Content) Copy() rbx.Value {
	return c
}
