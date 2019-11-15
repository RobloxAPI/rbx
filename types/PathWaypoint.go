package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
	"github.com/robloxapi/rbx/enum"
)

type PathWaypoint struct {
	Position Vector3
	Action   enum.PathWaypointAction
}

func (PathWaypoint) Type() string {
	return "PathWaypoint"
}
func (p PathWaypoint) String() string {
	b := make([]byte, 0, 32)
	b = strconv.AppendFloat(b, float64(p.Position.X), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendFloat(b, float64(p.Position.Y), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendFloat(b, float64(p.Position.Z), 'g', -1, 32)
	b = append(b, ',', ' ')
	b = strconv.AppendUint(b, uint64(p.Action), 10)
	return string(b)
}
func (p PathWaypoint) Copy() rbx.Value {
	return p
}
