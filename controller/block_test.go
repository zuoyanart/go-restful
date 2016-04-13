package controller

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	. "github.com/smartystreets/goconvey/convey"
	"pizzaCmsApi/model"
	"testing"
)

var url = "http://127.0.0.1:3000/v1/"
var request = gorequest.New()

type jsonBlock struct {
	State bool
	Msg   model.Block
}

func TestBlockGet(t *testing.T) {
	Convey("获取模块", t, func() {
		_, body, _ := request.Get(url + "block/1").End()
		var jb jsonBlock
		json.Unmarshal([]byte(body), &jb)
		So(jb.State, ShouldEqual, true)
		So(jb.Msg.ID, ShouldEqual, 1)
	})
}

func TestBlockUpdate(t *testing.T) {
	Convey("更新模块", t, func() {
		_, body, _ := request.Put(url + "block").Send(`{"id":1, "title": "ces","content":"content"}`).End()
		var jb jsonBlock
		json.Unmarshal([]byte(body), &jb)
		So(jb.State, ShouldEqual, true)
		Convey("更新后验证", func() {
			_, body, _ := request.Get(url + "block/1").End()
			var jb jsonBlock
			json.Unmarshal([]byte(body), &jb)
			So(jb.State, ShouldEqual, true)
			So(jb.Msg.ID, ShouldEqual, 1)
			So(jb.Msg.Title, ShouldEqual, "ces")
			So(jb.Msg.Content, ShouldEqual, "content")
		})
	})
}

func TestBlockCreate(t *testing.T) {
	Convey("创建模块", t, func() {
		_, body, _ := request.Post(url + "block").Send(`{ "title": "ces","content":"content"}`).End()
		var jb jsonBlock
		json.Unmarshal([]byte(body), &jb)
		So(jb.State, ShouldEqual, true)
		Convey("创建后验证", func() {
			_, body, _ := request.Get(url + "block/1").End()
			var jb jsonBlock
			json.Unmarshal([]byte(body), &jb)
			So(jb.State, ShouldEqual, true)
			So(jb.Msg.Title, ShouldEqual, "ces")
			So(jb.Msg.Content, ShouldEqual, "content")
		})
	})
}
