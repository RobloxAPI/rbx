package types

import (
	"strings"

	"github.com/robloxapi/rbx"
)

type NumberSequence []NumberSequenceKeypoint

func (NumberSequence) Type() string {
	return "NumberSequence"
}
func (s NumberSequence) String() string {
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
func (n NumberSequence) Copy() rbx.Value {
	cp := make(NumberSequence, len(n))
	for i, v := range n {
		cp[i] = v.Copy().(NumberSequenceKeypoint)
	}
	return cp
}
