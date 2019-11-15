package types

import (
	"github.com/robloxapi/rbx"
	"github.com/robloxapi/rbx/enum"
)

type Faces struct {
	Right, Top, Back, Left, Bottom, Front bool
}

func (Faces) Type() string {
	return "Faces"
}
func (f Faces) String() string {
	if f == (Faces{}) {
		return "(none)"
	}
	b := make([]byte, 0, 7)
	if f.Right {
		b = append(b, "Right"...)
	}
	if f.Top {
		if len(b) > 0 {
			b = append(b, ',', ' ')
		}
		b = append(b, "Top"...)
	}
	if f.Back {
		if len(b) > 0 {
			b = append(b, ',', ' ')
		}
		b = append(b, "Back"...)
	}
	if f.Left {
		if len(b) > 0 {
			b = append(b, ',', ' ')
		}
		b = append(b, "Left"...)
	}
	if f.Bottom {
		if len(b) > 0 {
			b = append(b, ',', ' ')
		}
		b = append(b, "Bottom"...)
	}
	if f.Front {
		if len(b) > 0 {
			b = append(b, ',', ' ')
		}
		b = append(b, "Front"...)
	}
	return string(b)
}
func (f Faces) Copy() rbx.Value {
	return f
}

func NewFacesFromNormalId(faces ...enum.NormalId) Faces {
	f := Faces{}
	for _, face := range faces {
		switch face {
		case 1:
			f.Top = true
			continue
		case 4:
			f.Bottom = true
			continue
		case 2:
			f.Back = true
			continue
		case 5:
			f.Front = true
			continue
		case 0:
			f.Right = true
			continue
		case 3:
			f.Left = true
			continue
		}
	}
	return f
}
