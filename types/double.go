package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
)

type Double float64

func (Double) Type() string {
	return "Double"
}
func (d Double) String() string {
	return strconv.FormatFloat(float64(d), 'g', -1, 64)
}
func (d Double) Copy() rbx.Value {
	return d
}
