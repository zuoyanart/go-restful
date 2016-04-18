package restest

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"pizzaCmsApi/model"
	"testing"
)

type jsonUsergroup struct {
	State bool
	Msg   model.Usergroup
}

type jsonUsergroupGet struct {
	State bool
	Msg   int
}

type jsonUsergroupList struct {
	State bool
	Msg   []model.Usergroup
	Count int
}

func TestUsergroupCreate(t *testing.T) {
	Convey("创建用户组", t, func() {
		_, body, _ := request.Post(Url + "usergroup").Send(`{"name": "name","des": "des","state": 1,"1":"1"}`).End()
		var jsonEnt jsonUsergroupGet
		json.Unmarshal([]byte(body), &jsonEnt) //转换成struct
		id = jsonEnt.Msg                       //赋值id
		So(jsonEnt.State, ShouldEqual, true)
		Convey("创建后验证", func() {
			_, body, _ := request.Get(Url + "usergroup/" + Tools.ParseString(id)).End()
			var jsonEnt jsonUsergroup
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, true)
			So(jsonEnt.Msg.Name, ShouldEqual, "name")
			So(jsonEnt.Msg.Des, ShouldEqual, "des")
			So(jsonEnt.Msg.State, ShouldEqual, 0)

		})
	})
}

func TestUsergroupGet(t *testing.T) {
	Convey("获取用户组", t, func() {
		_, body, _ := request.Get(Url + "usergroup/" + Tools.ParseString(id)).End()
		var jsonEnt jsonUsergroup
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		So(jsonEnt.Msg.Id, ShouldEqual, id)
	})
}

func TestUsergroupUpdate(t *testing.T) {
	Convey("更新用户组", t, func() {
		_, body, _ := request.Put(Url + "usergroup").Send(`{"id":` + Tools.ParseString(id) + `,"name": "name","des": "des","state": 1,"1":"1"}`).End()
		var jsonEnt jsonUsergroup
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		Convey("更新后验证", func() {
			_, body, _ := request.Get(Url + "usergroup/" + Tools.ParseString(id)).End()
			var jsonEnt jsonUsergroup
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, true)
			So(jsonEnt.Msg.Name, ShouldEqual, "name")
			So(jsonEnt.Msg.Des, ShouldEqual, "des")
			So(jsonEnt.Msg.State, ShouldEqual, 0)

		})
	})
}

func TestUsergroupPage(t *testing.T) {
	Convey("获取列表用户组", t, func() {
		_, body, _ := request.Post(Url + "usergroup/page").Send("cp=1&mp=10&kw=11").End()
		var jsonEnt jsonUsergroupList
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		So(jsonEnt.Count, ShouldNotEqual, 0)
	})
}

func TestUsergroupDele(t *testing.T) {
	Convey("删除用户组", t, func() {
		_, body, _ := request.Delete(Url + "usergroup").Query("id=" + Tools.ParseString(id)).End()
		var jsonEnt jsonUsergroup
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		Convey("获取验证", func() {
			_, body, _ := request.Get(Url + "usergroup/" + Tools.ParseString(id)).End()
			var jsonEnt jsonUsergroup
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, false)
		})
	})
}
