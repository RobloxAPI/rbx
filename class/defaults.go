package class

import (
	"github.com/robloxapi/rbx/types"
)

func init() {
	Instance.Property("Archivable").Value = types.Bool(true)
	Instance.Property("ClassName").Value = types.String("")
	Instance.Property("DataCost").Value = types.Int(0)
	Instance.Property("Name").Value = types.String("Instance")
	Instance.Property("Parent").Value = types.NilInstance
	Instance.Property("RobloxLocked").Value = types.Bool(false)
	Instance.Property("archivable").Value = types.Bool(true)
	Instance.Property("className").Value = types.String("")
}
