package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
)

type NumberRange struct {
	Min, Max float32
}

func (NumberRange) Type() string {
	return "NumberRange"
}
func (n NumberRange) String() string {
	b := make([]byte, 0, 16)
	b = strconv.AppendFloat(b, float64(n.Min), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendFloat(b, float64(n.Max), 'g', -1, 32)
	return string(b)
}
func (n NumberRange) Copy() rbx.Value {
	return n
}
