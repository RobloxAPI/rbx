package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
)

type CFrame struct {
	Position Vector3
	Rotation [9]float32
}

func (CFrame) Type() string {
	return "CFrame"
}
func (c CFrame) String() string {
	b := make([]byte, 0, 128)
	b = strconv.AppendFloat(b, float64(c.Position.X), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendFloat(b, float64(c.Position.Y), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendFloat(b, float64(c.Position.Z), 'g', -1, 32)
	for _, v := range c.Rotation {
		b = append(b, ',', ' ')
		b = strconv.AppendFloat(b, float64(v), 'g', -1, 32)
	}
	return string(b)
}
func (c CFrame) Copy() rbx.Value {
	return c
}

var cframeID = CFrame{
	Position: Vector3{X: 0, Y: 0, Z: 0},
	Rotation: [9]float32{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	},
}

//// Fields

func (c CFrame) X() float64 {
	return float64(c.Position.X)
}
func (c CFrame) Y() float64 {
	return float64(c.Position.Y)
}
func (c CFrame) Z() float64 {
	return float64(c.Position.Z)
}
func (c CFrame) RightVector() Vector3 {
	return Vector3{c.Rotation[0], c.Rotation[3], c.Rotation[6]}
}
func (c CFrame) UpVector() Vector3 {
	return Vector3{c.Rotation[1], c.Rotation[4], c.Rotation[7]}
}
func (c CFrame) LookVector() Vector3 {
	return Vector3{-c.Rotation[2], -c.Rotation[5], -c.Rotation[8]}
}

//// Methods

func (c CFrame) Inverse() CFrame {
	return CFrame{}
}
func (c CFrame) Lerp(goal CFrame, alpha float64) CFrame {
	return CFrame{}
}
func (c CFrame) ToWorldSpace(cf CFrame) CFrame {
	return CFrame{}
}
func (c CFrame) ToObjectSpace(cf CFrame) CFrame {
	return CFrame{}
}
func (c CFrame) PointToWorldSpace(v Vector3) Vector3 {
	return Vector3{}
}
func (c CFrame) PointToObjectSpace(v Vector3) Vector3 {
	return Vector3{}
}
func (c CFrame) VectorToWorldSpace(v Vector3) Vector3 {
	return Vector3{}
}
func (c CFrame) VectorToObjectSpace(v Vector3) Vector3 {
	return Vector3{}
}
func (c CFrame) Components() (x, y, z, r00, r01, r02, r10, r11, r12, r20, r21, r22 float64) {
	return
}
func (c CFrame) Angles() (rx, ry, rz float64) {
	return
}
func (c CFrame) Orientation() (rx, ry, rz float64) {
	return
}
func (c CFrame) AxisAngle() (axis Vector3, rotation float64) {
	return
}

//// Operators

func (c CFrame) Mul(op CFrame) CFrame {
	return CFrame{}
}
func (c CFrame) MulVec(op Vector3) Vector3 {
	return Vector3{}
}
func (c CFrame) AddVec(op Vector3) CFrame {
	return CFrame{}
}
func (c CFrame) SubVec(op Vector3) CFrame {
	return CFrame{}
}

//// Constructors

func NewCFrame() CFrame {
	return cframeID
}
func NewCFrameFromVector3(v Vector3) CFrame {
	return CFrame{
		Position: v,
		Rotation: cframeID.Rotation,
	}
}
func NewCFrameFromLook(pos, lookAt Vector3) CFrame {
	return CFrame{}
}
func NewCFrameFromPosition(x, y, z float64) CFrame {
	return CFrame{
		Position: Vector3{X: float32(x), Y: float32(y), Z: float32(z)},
		Rotation: cframeID.Rotation,
	}
}
func NewCFrameFromQuat(x, y, z, qx, qy, qz, qw float64) CFrame {
	return CFrame{}
}
func NewCFrameFromComponents(x, y, z, r00, r01, r02, r10, r11, r12, r20, r21, r22 float64) CFrame {
	return CFrame{
		Position: Vector3{X: float32(x), Y: float32(y), Z: float32(z)},
		Rotation: [9]float32{
			float32(r00), float32(r01), float32(r02),
			float32(r10), float32(r11), float32(r12),
			float32(r20), float32(r21), float32(r22),
		},
	}
}
func NewCFrameFromAngles(rx, ry, rz float64) CFrame {
	return CFrame{}
}
func NewCFrameFromOrientation(rx, ry, rz float64) CFrame {
	return CFrame{}
}
func NewCFrameFromAxisAngle(axis Vector3, rotation float64) CFrame {
	return CFrame{}
}
func NewCFrameFromMatrix(p, vx, vy, vz Vector3) CFrame {
	return CFrame{}
}
