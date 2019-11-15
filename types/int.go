package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
)

type Int int32

func (Int) Type() string {
	return "int"
}
func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}
func (i Int) Copy() rbx.Value {
	return i
}
