package restest

import (
	"github.com/parnurzeal/gorequest"
	"pizzaCmsApi/tools"
)

var (
	Tools   *tools.Tools
	Url     string
	id      int
	request *gorequest.SuperAgent
)

func init() {
	Tools = tools.New()
	Url = "http://127.0.0.1:3000/v1/"
	request = gorequest.New()

}
