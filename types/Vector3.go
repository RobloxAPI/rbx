package types

import (
	"math"
	"strconv"

	"github.com/robloxapi/rbx"
	"github.com/robloxapi/rbx/enum"
)

type Vector3 struct {
	X, Y, Z float32
}

func (Vector3) Type() string {
	return "Vector3"
}
func (v Vector3) String() string {
	b := make([]byte, 0, 24)
	b = strconv.AppendFloat(b, float64(v.X), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendFloat(b, float64(v.Y), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendFloat(b, float64(v.Z), 'g', -1, 32)
	return string(b)
}
func (v Vector3) Copy() rbx.Value {
	return v
}

//// Fields

func (v Vector3) Magnitude() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}
func (v Vector3) Unit() Vector3 {
	m := float32(v.Magnitude())
	return Vector3{X: v.X / m, Y: v.Y / m, Z: v.Z / m}
}

//// Methods

func (v Vector3) Lerp(goal Vector3, alpha float64) Vector3 {
	a := float32(alpha)
	na := 1 - a
	return Vector3{
		X: na*v.X + a*goal.X,
		Y: na*v.Y + a*goal.Y,
		Z: na*v.Z + a*goal.Z,
	}
}
func (v Vector3) Dot(op Vector3) float64 {
	return float64(v.X*op.X + v.Y*op.Y + v.Z*op.Z)
}
func (v Vector3) Cross(op Vector3) Vector3 {
	return Vector3{
		v.Y*op.Z - v.Z*op.Y,
		v.Z*op.X - v.X*op.Z,
		v.X*op.Y - v.Y*op.X,
	}
}
func (v Vector3) FuzzyEq(op Vector3, epsilon float64) bool {
	switch {
	case epsilon == 0:
		return v == op
	case epsilon < 0:
		// Default.
		epsilon = 1e-5
	}
	x := v.X - op.X
	y := v.Y - op.Y
	z := v.Z - op.Z
	return x*x+y*y+z*z <= float32(epsilon)
}

//// Operators

func (v Vector3) Add(op Vector3) Vector3 {
	return Vector3{X: v.X + op.X, Y: v.Y + op.Y, Z: v.Z + op.Z}
}
func (v Vector3) Sub(op Vector3) Vector3 {
	return Vector3{X: v.X - op.X, Y: v.Y - op.Y, Z: v.Z - op.Z}
}
func (v Vector3) Mul(op Vector3) Vector3 {
	return Vector3{X: v.X * op.X, Y: v.Y * op.Y, Z: v.Z * op.Z}
}
func (v Vector3) Div(op Vector3) Vector3 {
	return Vector3{X: v.X / op.X, Y: v.Y / op.Y, Z: v.Z / op.Z}
}
func (v Vector3) MulN(op float64) Vector3 {
	return Vector3{X: v.X * float32(op), Y: v.Y * float32(op), Z: v.Z * float32(op)}
}
func (v Vector3) DivN(op float64) Vector3 {
	return Vector3{X: v.X / float32(op), Y: v.Y / float32(op), Z: v.Z / float32(op)}
}

//// Constructors

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{X: float32(x), Y: float32(y), Z: float32(z)}
}
func NewVector3FromNormalID(normal enum.NormalId) Vector3 {
	switch normal {
	case 0:
		return Vector3{X: 1, Y: 0, Z: 0}
	case 1:
		return Vector3{X: 0, Y: 1, Z: 0}
	case 2:
		return Vector3{X: 0, Y: 0, Z: 1}
	case 3:
		return Vector3{X: -1, Y: 0, Z: 0}
	case 4:
		return Vector3{X: 0, Y: -1, Z: 0}
	case 5:
		return Vector3{X: 0, Y: 0, Z: -1}
	}
	panic("expected Enum.NormalId input")
}
func NewVector3FromAxis(axis enum.Axis) Vector3 {
	switch axis {
	case 0:
		return Vector3{X: 1, Y: 0, Z: 0}
	case 1:
		return Vector3{X: 0, Y: 1, Z: 0}
	case 2:
		return Vector3{X: 0, Y: 0, Z: 1}
	}
	panic("expected Enum.Axis input")
}
