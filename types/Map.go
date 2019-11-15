package types

import (
	"strings"

	"github.com/robloxapi/rbx"
)

type Map map[rbx.Value]rbx.Value

func (Map) Type() string {
	return "Map"
}
func (m Map) String() string {
	switch len(m) {
	case 0:
		return "[]"
	}

	var buf strings.Builder
	buf.WriteByte('[')
	var mult bool
	for k, v := range m {
		if mult {
			buf.WriteString("; ")
		}
		buf.WriteString(k.String())
		buf.WriteString(": ")
		buf.WriteString(v.String())
		mult = true
	}
	buf.WriteByte(']')
	return buf.String()
}
func (m Map) Copy() rbx.Value {
	cp := make(Map, len(m))
	for k, v := range m {
		cp[k] = v.Copy()
	}
	return cp
}
