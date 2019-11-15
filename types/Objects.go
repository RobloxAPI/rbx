package types

import (
	"strings"

	"github.com/robloxapi/rbx"
)

type Objects []rbx.Instance

func (Objects) Type() string {
	return "Objects"
}
func (o Objects) String() string {
	switch len(o) {
	case 0:
		return "[]"
	case 1:
		return "[" + o[0].String() + "]"
	}
	var buf strings.Builder
	buf.WriteByte('[')
	for i, v := range o {
		if i > 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(v.String())
	}
	buf.WriteByte(']')
	return buf.String()
}
func (o Objects) Copy() rbx.Value {
	cp := make(Objects, len(o))
	for i, v := range o {
		cp[i] = v
	}
	return cp
}
