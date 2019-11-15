package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
)

type Color3 struct {
	R, G, B float32
}

func (Color3) Type() string {
	return "Color3"
}
func (c Color3) String() string {
	b := make([]byte, 0, 24)
	b = strconv.AppendFloat(b, float64(c.R), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendFloat(b, float64(c.G), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendFloat(b, float64(c.B), 'g', -1, 32)
	return string(b)
}
func (c Color3) Copy() rbx.Value {
	return c
}
