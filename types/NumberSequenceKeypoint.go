package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
)

type NumberSequenceKeypoint struct {
	Time     float32
	Value    float32
	Envelope float32
}

func (NumberSequenceKeypoint) Type() string {
	return "NumberSequenceKeypoint"
}
func (k NumberSequenceKeypoint) String() string {
	b := make([]byte, 0, 32)
	b = strconv.AppendFloat(b, float64(k.Time), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendFloat(b, float64(k.Value), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendFloat(b, float64(k.Envelope), 'g', -1, 32)
	return string(b)
}
func (n NumberSequenceKeypoint) Copy() rbx.Value {
	return n
}
