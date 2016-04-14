package restest

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"pizzaCmsApi/model"
	"testing"
)

type jsonComment struct {
	State bool
	Msg   model.Comment
}

type jsonCommentGet struct {
	State bool
	Msg   int
}

type jsonCommentList struct {
	State bool
	Msg   []model.Comment
	Count int
}

func TestCommentCreate(t *testing.T) {
	Convey("创建文章评论", t, func() {
		_, body, _ := request.Post(Url + "comment").Send(`{"articleid": 1,"addtime": 1,"content": "content","uid": 1,"username": "username"}`).End()
		var jsonEnt jsonCommentGet
		json.Unmarshal([]byte(body), &jsonEnt) //转换成struct
		id = jsonEnt.Msg                       //赋值id
		So(jsonEnt.State, ShouldEqual, true)
		Convey("创建后验证", func() {
			_, body, _ := request.Get(Url + "comment/" + Tools.ParseString(id)).End()
			var jsonEnt jsonComment
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, true)
			So(jsonEnt.Msg.Articleid, ShouldEqual, 1)
			So(jsonEnt.Msg.Addtime, ShouldEqual, 1)
			So(jsonEnt.Msg.Content, ShouldEqual, "content")
			So(jsonEnt.Msg.Uid, ShouldEqual, 1)
			So(jsonEnt.Msg.Username, ShouldEqual, "username")
		})
	})
}

func TestCommentGet(t *testing.T) {
	Convey("获取文章评论", t, func() {
		_, body, _ := request.Get(Url + "comment/" + Tools.ParseString(id)).End()
		var jsonEnt jsonComment
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		So(jsonEnt.Msg.Id, ShouldEqual, id)
	})
}

func TestCommentUpdate(t *testing.T) {
	Convey("更新文章评论", t, func() {
		_, body, _ := request.Put(Url + "comment").Send(`{"id":` + Tools.ParseString(id) + `,"articleid": 1,"addtime": 1,"content": "content","uid": 1,"username": "username","1":"1"}`).End()
		var jsonEnt jsonComment
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		Convey("更新后验证", func() {
			_, body, _ := request.Get(Url + "comment/" + Tools.ParseString(id)).End()
			var jsonEnt jsonComment
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, true)
			So(jsonEnt.Msg.Articleid, ShouldEqual, 1)
			So(jsonEnt.Msg.Addtime, ShouldEqual, 1)
			So(jsonEnt.Msg.Content, ShouldEqual, "content")
			So(jsonEnt.Msg.Uid, ShouldEqual, 1)
			So(jsonEnt.Msg.Username, ShouldEqual, "username")

		})
	})
}

func TestCommentPage(t *testing.T) {
	Convey("获取列表文章评论", t, func() {
		_, body, _ := request.Post(Url + "comment/page").Send("cp=1&mp=10&kw=").End()
		var jsonEnt jsonCommentList
		json.Unmarshal([]byte(body), &jsonEnt)
		// So(body, ShouldEqual, true)
		So(jsonEnt.State, ShouldEqual, true)
		So(jsonEnt.Count, ShouldNotEqual, 0)
	})
}

func TestCommentDele(t *testing.T) {
	Convey("删除文章评论", t, func() {
		_, body, _ := request.Delete(Url + "comment").Query("id=" + Tools.ParseString(id)).End()
		var jsonEnt jsonComment
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		Convey("获取验证", func() {
			_, body, _ := request.Get(Url + "comment/" + Tools.ParseString(id)).End()
			var jsonEnt jsonComment
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, false)
		})
	})
}
