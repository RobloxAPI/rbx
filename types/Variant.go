package types

import (
	"github.com/robloxapi/rbx"
)

type Variant struct {
	rbx.Value
}

func (Variant) Type() string {
	return "Variant"
}
func (v Variant) String() string {
	return v.Value.String()
}
func (v Variant) Copy() rbx.Value {
	return Variant{Value: v.Value.Copy()}
}
