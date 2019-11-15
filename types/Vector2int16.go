package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
)

type Vector2int16 struct {
	X, Y int16
}

func (Vector2int16) Type() string {
	return "Vector2int16"
}
func (v Vector2int16) String() string {
	b := make([]byte, 0, 16)
	strconv.AppendInt(b, int64(v.X), 10)
	b = append(b, ',', ' ')
	strconv.AppendInt(b, int64(v.Y), 10)
	return string(b)
}
func (v Vector2int16) Copy() rbx.Value {
	return v
}
