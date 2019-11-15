package types

import (
	"strings"

	"github.com/robloxapi/rbx"
)

type Array []rbx.Value

func (Array) Type() string {
	return "Array"
}
func (a Array) String() string {
	switch len(a) {
	case 0:
		return "[]"
	case 1:
		return "[" + a[0].String() + "]"
	}
	var buf strings.Builder
	buf.WriteByte('[')
	for i, v := range a {
		if i > 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(v.String())
	}
	buf.WriteByte(']')
	return buf.String()
}
func (a Array) Copy() rbx.Value {
	cp := make(Array, len(a))
	for i, v := range a {
		cp[i] = v.Copy()
	}
	return cp
}
