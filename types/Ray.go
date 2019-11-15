package types

import (
	"github.com/robloxapi/rbx"
)

type Ray struct {
	Origin    Vector3
	Direction Vector3
}

func (Ray) Type() string {
	return "Ray"
}
func (r Ray) String() string {
	return "{" + r.Origin.String() + "}, {" + r.Direction.String() + "}"
}
func (r Ray) Copy() rbx.Value {
	return Ray{
		Origin:    r.Origin.Copy().(Vector3),
		Direction: r.Direction.Copy().(Vector3),
	}
}
