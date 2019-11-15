package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
)

type Vector3int16 struct {
	X, Y, Z int16
}

func (Vector3int16) Type() string {
	return "Vector3int16"
}
func (v Vector3int16) String() string {
	b := make([]byte, 0, 16)
	strconv.AppendInt(b, int64(v.X), 10)
	b = append(b, ',', ' ')
	strconv.AppendInt(b, int64(v.Y), 10)
	b = append(b, ',', ' ')
	strconv.AppendInt(b, int64(v.Z), 10)
	return string(b)
}
func (v Vector3int16) Copy() rbx.Value {
	return v
}
