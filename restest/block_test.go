package restest

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"pizzaCmsApi/model"
	"testing"
)

type jsonBlock struct {
	State bool
	Msg   model.Block
}

type jsonBlockGet struct {
	State bool
	Msg   int
}

type jsonBlockList struct {
	State bool
	Msg   []model.Block
	Count int
}

func TestBlockCreate(t *testing.T) {
	Convey("创建模块", t, func() {
		_, body, _ := request.Post(Url + "block").Send(`{"title": "title","content": "content","1":"1"}`).End()
		var jsonEnt jsonBlockGet
		json.Unmarshal([]byte(body), &jsonEnt) //转换成struct
		id = jsonEnt.Msg                       //赋值id
		So(jsonEnt.State, ShouldEqual, true)
		// So(body, ShouldEqual, true)
		Convey("创建后验证", func() {
			_, body, _ := request.Get(Url + "block/" + Tools.ParseString(id)).End()
			var jsonEnt jsonBlock
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, true)
			So(jsonEnt.Msg.Title, ShouldEqual, "title")
			So(jsonEnt.Msg.Content, ShouldEqual, "content")

		})
	})
}

func TestBlockGet(t *testing.T) {
	Convey("获取模块", t, func() {
		_, body, _ := request.Get(Url + "block/" + Tools.ParseString(id)).End()
		var jsonEnt jsonBlock
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		So(jsonEnt.Msg.Id, ShouldEqual, id)
	})
}

func TestBlockUpdate(t *testing.T) {
	Convey("更新模块", t, func() {
		_, body, _ := request.Put(Url + "block").Send(`{"id":`+Tools.ParseString(id)+`,"title": "title","content": "content","1":"1"}}`).End()
		var jsonEnt jsonBlock
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		Convey("更新后验证", func() {
			_, body, _ := request.Get(Url + "block/" + Tools.ParseString(id)).End()
			var jsonEnt jsonBlock
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, true)
			So(jsonEnt.Msg.Title, ShouldEqual, "title")
			So(jsonEnt.Msg.Content, ShouldEqual, "content")

		})
	})
}

func TestBlockDele(t *testing.T) {
	Convey("删除模块", t, func() {
		_, body, _ := request.Delete(Url + "block").Query("id=" + Tools.ParseString(id)).End()
		var jsonEnt jsonBlock
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		Convey("获取验证", func() {
			_, body, _ := request.Get(Url + "block/" + Tools.ParseString(id)).End()
			var jsonEnt jsonBlock
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, false)
		})
	})
}

func TestBlockPage(t *testing.T) {
	Convey("获取列表模块", t, func() {
		_, body, _ := request.Post(Url + "block/page").Send("cp=1&mp=10&kw=").End()
		var jsonEnt jsonBlockList
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		So(jsonEnt.Count, ShouldNotEqual, 0)
	})
}
