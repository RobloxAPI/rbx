package types

import (
	"github.com/robloxapi/rbx"
)

type Rect struct {
	Min, Max Vector2
}

func (Rect) Type() string {
	return "Rect"
}
func (r Rect) String() string {
	return r.Min.String() + ", " + r.Max.String()
}
func (r Rect) Copy() rbx.Value {
	return Rect{
		Min: r.Min.Copy().(Vector2),
		Max: r.Max.Copy().(Vector2),
	}
}
