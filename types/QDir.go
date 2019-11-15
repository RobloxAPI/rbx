package types

import (
	"github.com/robloxapi/rbx"
)

type QDir string

func (QDir) Type() string {
	return "QDir"
}
func (s QDir) String() string {
	return string(s)
}
func (s QDir) Copy() rbx.Value {
	return s
}
