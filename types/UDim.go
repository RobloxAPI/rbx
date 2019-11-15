package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
)

type UDim struct {
	Scale  float32
	Offset int32
}

func (UDim) Type() string {
	return "UDim"
}
func (u UDim) String() string {
	b := make([]byte, 0, 16)
	b = strconv.AppendFloat(b, float64(u.Scale), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendInt(b, int64(u.Offset), 10)
	return string(b)
}
func (u UDim) Copy() rbx.Value {
	return u
}
