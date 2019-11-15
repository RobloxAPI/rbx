package types

import (
	"strconv"

	"github.com/robloxapi/rbx"
)

type Int64 int64

func (Int64) Type() string {
	return "int64"
}
func (i Int64) String() string {
	return strconv.FormatInt(int64(i), 10)
}
func (i Int64) Copy() rbx.Value {
	return i
}
