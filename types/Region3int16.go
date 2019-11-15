package types

import (
	"github.com/robloxapi/rbx"
)

type Region3int16 struct {
	Min, Max Vector3int16
}

func (Region3int16) Type() string {
	return "Region3int16"
}
func (r Region3int16) String() string {
	return r.Min.String() + ", " + r.Max.String()
}
func (r Region3int16) Copy() rbx.Value {
	return Region3int16{
		Min: r.Min.Copy().(Vector3int16),
		Max: r.Max.Copy().(Vector3int16),
	}
}
