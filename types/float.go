package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
)

type Float float32

func (Float) Type() string {
	return "float"
}
func (f Float) String() string {
	return strconv.FormatFloat(float64(f), 'g', -1, 64)

}
func (f Float) Copy() rbx.Value {
	return f
}
