package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
	"github.com/robloxapi/rbx/enum"
)

type DockWidgetPluginGuiInfo struct {
	InitialDockState                    enum.InitialDockState
	InitialEnabled                      bool
	InitialEnabledShouldOverrideRestore bool
	FloatingXSize                       uint32
	FloatingYSize                       uint32
	MinWidth                            uint32
	MinHeight                           uint32
}

func (DockWidgetPluginGuiInfo) Type() string {
	return "DockWidgetPluginGuiInfo"
}
func (d DockWidgetPluginGuiInfo) String() string {
	b := make([]byte, 0, 48)
	b = strconv.AppendUint(b, uint64(d.InitialDockState), 10)
	if d.InitialEnabled {
		b = append(b, "true"...)
	} else {
		b = append(b, "false"...)
	}
	b = append(b, ',', ' ')
	if d.InitialEnabledShouldOverrideRestore {
		b = append(b, "true"...)
	} else {
		b = append(b, "false"...)
	}
	b = append(b, ',', ' ')
	b = strconv.AppendUint(b, uint64(d.FloatingXSize), 10)
	b = append(b, ',', ' ')
	b = strconv.AppendUint(b, uint64(d.FloatingYSize), 10)
	b = append(b, ',', ' ')
	b = strconv.AppendUint(b, uint64(d.MinWidth), 10)
	b = append(b, ',', ' ')
	b = strconv.AppendUint(b, uint64(d.MinHeight), 10)
	return string(b)
}
func (d DockWidgetPluginGuiInfo) Copy() rbx.Value {
	return d
}

//// Constructors

func NewDockWidgetPluginGuiInfo() DockWidgetPluginGuiInfo {
	return DockWidgetPluginGuiInfo{
		InitialDockState:                    3,
		InitialEnabled:                      false,
		InitialEnabledShouldOverrideRestore: false,
		FloatingXSize:                       0,
		FloatingYSize:                       0,
		MinWidth:                            0,
		MinHeight:                           0,
	}
}
