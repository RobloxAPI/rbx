package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
)

type Vector2 struct {
	X, Y float32
}

func (Vector2) Type() string {
	return "Vector2"
}
func (v Vector2) String() string {
	b := make([]byte, 0, 16)
	b = strconv.AppendFloat(b, float64(v.X), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendFloat(b, float64(v.Y), 'g', -1, 32)
	return string(b)
}
func (v Vector2) Copy() rbx.Value {
	return v
}
