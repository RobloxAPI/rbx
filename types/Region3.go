package types

import (
	"github.com/robloxapi/rbx"
)

type Region3 struct {
	Min, Max Vector3
}

func (Region3) Type() string {
	return "Region3"
}
func (r Region3) String() string {
	return r.Min.String() + ", " + r.Max.String()
}
func (r Region3) Copy() rbx.Value {
	return Region3{
		Min: r.Min.Copy().(Vector3),
		Max: r.Max.Copy().(Vector3),
	}
}
