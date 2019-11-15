// The class package contains class descriptors generated automatically from the
// Roblox API.
package class

import (
	"github.com/robloxapi/rbx/desc"
)

// Get returns the class descriptor matching the given name, or nil if no match
// was found.
func Get(name string) *desc.Class {
	return classes[name]
}

// List returns all of the classes in the package.
func List() []*desc.Class {
	l := make([]*desc.Class, len(list))
	copy(l, list)
	return l
}
