package logic

import (
	"pizzaCmsApi/tools"
)

var (
	Tools *tools.Tools
)

func init() {
	Tools = tools.New()
}
