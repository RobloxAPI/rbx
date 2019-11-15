package types

import (
	"github.com/robloxapi/rbx"
)

type UDim2 struct {
	X, Y UDim
}

func (UDim2) Type() string {
	return "UDim2"
}
func (u UDim2) String() string {
	return u.X.String() + ", " + u.Y.String()
}
func (u UDim2) Copy() rbx.Value {
	return UDim2{
		X: u.X.Copy().(UDim),
		Y: u.Y.Copy().(UDim),
	}
}
