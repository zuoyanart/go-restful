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
		_, body, _ := request.Post(Url + "block").Send(`{ "title": "ces","content":"content"}`).End()
		var jb jsonBlockGet
		json.Unmarshal([]byte(body), &jb) //转换成struct
		id = jb.Msg                       //赋值id
		So(jb.State, ShouldEqual, true)
		Convey("创建后验证", func() {
			_, body, _ := request.Get(Url + "block/" + Tools.ParseString(id)).End()
			var jb jsonBlock
			json.Unmarshal([]byte(body), &jb)
			So(jb.State, ShouldEqual, true)
			So(jb.Msg.Title, ShouldEqual, "ces")
			So(jb.Msg.Content, ShouldEqual, "content")
		})
	})
}

func TestBlockGet(t *testing.T) {
	Convey("获取模块", t, func() {
		_, body, _ := request.Get(Url + "block/" + Tools.ParseString(id)).End()
		var jb jsonBlock
		json.Unmarshal([]byte(body), &jb)
		So(jb.State, ShouldEqual, true)
		So(jb.Msg.ID, ShouldEqual, id)
	})
}

func TestBlockUpdate(t *testing.T) {
	Convey("更新模块", t, func() {
		_, body, _ := request.Put(Url + "block").Send(`{"id":` + Tools.ParseString(id) + `, "title": "ces","content":"content"}`).End()
		var jb jsonBlock
		json.Unmarshal([]byte(body), &jb)
		So(jb.State, ShouldEqual, true)
		Convey("更新后验证", func() {
			_, body, _ := request.Get(Url + "block/" + Tools.ParseString(id)).End()
			var jb jsonBlock
			json.Unmarshal([]byte(body), &jb)
			So(jb.State, ShouldEqual, true)
			So(jb.Msg.ID, ShouldEqual, id)
			So(jb.Msg.Title, ShouldEqual, "ces")
			So(jb.Msg.Content, ShouldEqual, "content")
		})
	})
}

func TestBlockDele(t *testing.T) {
	Convey("删除模块", t, func() {
		_, body, _ := request.Delete(Url + "block").Query("id=" + Tools.ParseString(id)).End()
		var jb jsonBlock
		json.Unmarshal([]byte(body), &jb)
		So(jb.State, ShouldEqual, true)
		Convey("获取验证", func() {
			_, body, _ := request.Get(Url + "block/" + Tools.ParseString(id)).End()
			var jb jsonBlock
			json.Unmarshal([]byte(body), &jb)
			So(jb.State, ShouldEqual, false)
		})
	})
}
