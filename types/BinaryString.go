package types

import (
	"github.com/robloxapi/rbx"
)

type BinaryString []byte

func (BinaryString) Type() string {
	return "BinaryString"
}
func (b BinaryString) String() string {
	return string(b)
}
func (b BinaryString) Copy() rbx.Value {
	v := make(BinaryString, len(b))
	copy(v, b)
	return v
}
