package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
)

type ColorSequenceKeypoint struct {
	Time     float32
	Value    Color3
	Envelope float32
}

func (ColorSequenceKeypoint) Type() string {
	return "ColorSequenceKeypoint"
}
func (k ColorSequenceKeypoint) String() string {
	v := k.Value.String()
	b := make([]byte, 0, len(v)+20)
	b = strconv.AppendFloat(b, float64(k.Time), 'g', -1, 32)
	b = append(b, ',', ' ', '[')
	b = append(b, v...)
	b = append(b, ']', ',', ' ')
	b = strconv.AppendFloat(b, float64(k.Envelope), 'g', -1, 32)
	return string(b)
}
func (c ColorSequenceKeypoint) Copy() rbx.Value {
	return c
}
