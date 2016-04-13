package controller

import (
	"github.com/parnurzeal/gorequest"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var url = "http://127.0.0.1:3000/v1/"
var request = gorequest.New()

func TestBlockGet(t *testing.T) {
	Convey("获取模块", t, func() {
		resp, body, errs := request.Get(url + "block/1").End()
		So(body, ShouldEqual, 2)
	})
}
