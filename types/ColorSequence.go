package types

import (
	"strings"

	"github.com/robloxapi/rbx"
)

type ColorSequence []ColorSequenceKeypoint

func (ColorSequence) Type() string {
	return "ColorSequence"
}
func (s ColorSequence) String() string {
	switch len(s) {
	case 0:
		return "[]"
	case 1:
		return "[" + s[0].String() + "]"
	}
	var buf strings.Builder
	buf.WriteByte('[')
	for i, k := range s {
		if i > 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(k.String())
	}
	buf.WriteByte(']')
	return buf.String()
}
func (c ColorSequence) Copy() rbx.Value {
	cp := make(ColorSequence, len(c))
	for i, v := range c {
		cp[i] = v.Copy().(ColorSequenceKeypoint)
	}
	return cp
}
