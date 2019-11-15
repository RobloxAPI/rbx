package desc

import (
	"github.com/robloxapi/rbx"
)

////////////////////////////////////////////////////////////////////////////////

// enum implements rbx.Enum.
type enum struct {
	name  string
	items []*enumItem
}

func (e enum) Type() string {
	return "Enum"
}
func (e enum) String() string {
	return "Enum." + e.name
}
func (e *enum) Copy() rbx.Value {
	cp := enum{
		name:  e.name,
		items: make([]*enumItem, len(e.items)),
	}
	for i, item := range e.items {
		cp.items[i] = item.Copy().(*enumItem)
		cp.items[i].enum = e
	}
	return &cp
}
func (e enum) Name() string {
	return e.name
}
func (e enum) Item(name string) rbx.EnumItem {
	for _, item := range e.items {
		if item.name == name {
			return item
		}
	}
	return nil
}
func (e enum) Items() []rbx.EnumItem {
	items := make([]rbx.EnumItem, len(e.items))
	for i, item := range e.items {
		items[i] = item
	}
	return items
}

////////////////////////////////////////////////////////////////////////////////

// enumItem implements rbx.EnumItem.
type enumItem struct {
	enum  *enum
	name  string
	value int
}

func (e enumItem) Type() string {
	return "Enum." + e.enum.name
}
func (e enumItem) String() string {
	return "Enum." + e.enum.name + "." + e.name
}
func (e *enumItem) Copy() rbx.Value {
	return &enumItem{
		enum:  e.enum,
		name:  e.name,
		value: e.value,
	}
}
func (e enumItem) Enum() rbx.Enum {
	return e.enum
}
func (e enumItem) Name() string {
	return e.name
}
func (e enumItem) Value() int {
	return e.value
}

////////////////////////////////////////////////////////////////////////////////

// Item is passed to NewEnum to describe the implementation of an enum item.
type Item struct {
	Name  string
	Value int
}

// Enum creates a new implementation of the rbx.Enum interface, according to the
// given enum name and items.
func Enum(name string, items ...Item) rbx.Enum {
	e := enum{
		name:  name,
		items: make([]*enumItem, len(items)),
	}
	for i, item := range items {
		e.items[i] = &enumItem{
			enum:  &e,
			name:  item.Name,
			value: item.Value,
		}
	}
	return &e
}
