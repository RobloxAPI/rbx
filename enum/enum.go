// The enum package contains implementations of the types.Enum and
// types.EnumItem interfaces, generated automatically from the Roblox API.
package enum

import (
	"github.com/robloxapi/rbx"
)

// Get returns the Enum matching the given name, or nil if no match was found.
func Get(name string) rbx.Enum {
	return enums[name]
}

// List returns all of the enums in the package.
func List() []rbx.Enum {
	l := make([]rbx.Enum, len(list))
	copy(l, list)
	return l
}
