package restest

import (
	"encoding/json"
	// "github.com/parnurzeal/gorequest"
	. "github.com/smartystreets/goconvey/convey"
	"pizzaCmsApi/model"
	"testing"
)

// var Url = "http://127.0.0.1:3000/v1/"
// var request = gorequest.New()
// var id = 0

type jsonUser struct {
	State bool
	Msg   model.User
}

type jsonUserGet struct {
	State bool
	Msg   int
}

type jsonUserList struct {
	State bool
	Msg   []model.User
	Count int
}

func TestUserCreate(t *testing.T) {
	Convey("创建用户", t, func() {
		_, body, _ := request.Post(Url + "user").Send(`{ "username": "username","nickname":"nickname", "password": "password"}`).End()
		var jb jsonUserGet
		json.Unmarshal([]byte(body), &jb) //转换成struct
		id = jb.Msg                       //赋值id
		So(jb.State, ShouldEqual, true)
		Convey("创建后验证", func() {
			_, body, _ := request.Get(Url + "user/" + Tools.ParseString(id)).End()
			var jb jsonUser
			json.Unmarshal([]byte(body), &jb)
			So(jb.State, ShouldEqual, true)
			So(jb.Msg.Username, ShouldEqual, "username")

		})
	})
}

func TestUserGet(t *testing.T) {
	Convey("获取用户", t, func() {
		_, body, _ := request.Get(Url + "user/" + Tools.ParseString(id)).End()
		var jb jsonUser
		json.Unmarshal([]byte(body), &jb)
		So(jb.State, ShouldEqual, true)
		So(jb.Msg.ID, ShouldEqual, id)
	})
}

func TestUserUpdate(t *testing.T) {
	Convey("更新用户", t, func() {
		_, body, _ := request.Put(Url + "user").Send(`{"id":` + Tools.ParseString(id) + `, "username": "username1","nickname":"nickname1", "password":"password1"}`).End()
		var jb jsonUser
		json.Unmarshal([]byte(body), &jb)
		So(jb.State, ShouldEqual, true)
		Convey("更新后验证", func() {
			_, body, _ := request.Get(Url + "user/" + Tools.ParseString(id)).End()
			var jb jsonUser
			json.Unmarshal([]byte(body), &jb)
			So(jb.State, ShouldEqual, true)
			So(jb.Msg.ID, ShouldEqual, id)
			So(jb.Msg.Username, ShouldEqual, "username1")
		})
	})
}

func TestUserDele(t *testing.T) {
	Convey("删除用户", t, func() {
		_, body, _ := request.Delete(Url + "user").Query("id=" + Tools.ParseString(id)).End()
		var jb jsonUser
		json.Unmarshal([]byte(body), &jb)
		So(jb.State, ShouldEqual, true)
		Convey("获取验证", func() {
			_, body, _ := request.Get(Url + "user/" + Tools.ParseString(id)).End()
			var jb jsonUser
			json.Unmarshal([]byte(body), &jb)
			So(jb.State, ShouldEqual, false)
		})
	})
}
