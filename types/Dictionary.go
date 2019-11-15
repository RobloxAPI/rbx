package types

import (
	"sort"
	"strings"

	"github.com/robloxapi/rbx"
)

type Dictionary map[string]rbx.Value

func (Dictionary) Type() string {
	return "Dictionary"
}
func (d Dictionary) String() string {
	switch len(d) {
	case 0:
		return "[]"
	}

	keys := make([]string, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf strings.Builder
	buf.WriteByte('[')
	for i, k := range keys {
		if i > 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(k)
		buf.WriteString(": ")
		buf.WriteString(d[k].String())
	}
	buf.WriteByte(']')
	return buf.String()
}
func (d Dictionary) Copy() rbx.Value {
	cp := make(Dictionary, len(d))
	for k, v := range d {
		cp[k] = v.Copy()
	}
	return cp
}
