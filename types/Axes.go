package types

import (
	"github.com/robloxapi/rbx"
	"github.com/robloxapi/rbx/enum"
)

type Axes struct {
	X, Y, Z bool
}

func (Axes) Type() string {
	return "Axes"
}
func (a Axes) String() string {
	if a == (Axes{}) {
		return "(none)"
	}
	b := make([]byte, 0, 7)
	if a.X {
		b = append(b, 'X')
	}
	if a.Y {
		if a.X {
			b = append(b, ',', ' ')
		}
		b = append(b, 'Y')
	}
	if a.Z {
		if a.X || a.Y {
			b = append(b, ',', ' ')
		}
		b = append(b, 'Z')
	}
	return string(b)
}
func (a Axes) Copy() rbx.Value {
	return a
}

//// Fields

func (a Axes) Top() bool    { return a.Y }
func (a Axes) Bottom() bool { return a.Y }
func (a Axes) Back() bool   { return a.Z }
func (a Axes) Front() bool  { return a.Z }
func (a Axes) Right() bool  { return a.X }
func (a Axes) Left() bool   { return a.X }

//// Constructors

func NewAxes(x, y, z bool) Axes {
	return Axes{X: x, Y: y, Z: z}
}

func NewAxesFromEnum(axes ...rbx.EnumItem) Axes {
	a := Axes{}
	for _, item := range axes {
		switch item := item.(type) {
		case enum.Axis:
			switch item {
			case 0:
				a.X = true
				continue
			case 1:
				a.Y = true
				continue
			case 2:
				a.Z = true
				continue
			}
		case enum.NormalId:
			switch item {
			case 0, 3:
				a.X = true
				continue
			case 1, 4:
				a.Y = true
				continue
			case 2, 5:
				a.Z = true
				continue
			}
		}
		panic("expected Enum.Axis or Enum.NormalId inputs")
	}
	return a
}
