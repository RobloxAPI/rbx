package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
	"github.com/robloxapi/rbx/enum"
)

type TweenInfo struct {
	Time            float32
	DelayTime       float32
	RepeatCount     uint32
	Reverses        bool
	EasingStyle     enum.EasingStyle
	EasingDirection enum.EasingDirection
}

func (TweenInfo) Type() string {
	return "TweenInfo"
}
func (t TweenInfo) String() string {
	b := make([]byte, 0, 32)
	b = strconv.AppendFloat(b, float64(t.Time), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendFloat(b, float64(t.DelayTime), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendUint(b, uint64(t.RepeatCount), 10)
	if t.Reverses {
		b = append(b, "true"...)
	} else {
		b = append(b, "false"...)
	}
	b = append(b, ',', ' ')
	b = strconv.AppendUint(b, uint64(t.EasingStyle), 10)
	b = append(b, ',', ' ')
	b = strconv.AppendUint(b, uint64(t.EasingDirection), 10)
	return string(b)
}
func (t TweenInfo) Copy() rbx.Value {
	return t
}

//// Constructors

func NewTweenInfo() TweenInfo {
	return TweenInfo{
		Time:            1,
		EasingStyle:     3,
		EasingDirection: 1,
		RepeatCount:     0,
		Reverses:        false,
		DelayTime:       0,
	}
}

func NewTweenInfoFields(
	time float64,
	easingStyle enum.EasingStyle,
	easingDirection enum.EasingDirection,
	repeatCount float64,
	reverses bool,
	delayTime float64,
) TweenInfo {
	var t TweenInfo
	t.Time = float32(time)
	t.DelayTime = float32(delayTime)
	t.RepeatCount = uint32(repeatCount)
	t.Reverses = reverses
	t.EasingStyle = easingStyle
	t.EasingDirection = easingDirection
	return t
}
