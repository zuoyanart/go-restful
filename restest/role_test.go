package restest

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"pizzaCmsApi/model"
	"testing"
)

type jsonRole struct {
	State bool
	Msg   model.Role
}

type jsonRoleGet struct {
	State bool
	Msg   int
}

type jsonRoleList struct {
	State bool
	Msg   []model.Role
	Count int
}

func TestRoleCreate(t *testing.T) {
	Convey("创建用户组角色", t, func() {
		_, body, _ := request.Post(Url + "role").Send(`{"groupid": 1,"name": "name","des": "des","state": 1,"1":"1"}`).End()
		var jsonEnt jsonRoleGet
		json.Unmarshal([]byte(body), &jsonEnt) //转换成struct
		id = jsonEnt.Msg                       //赋值id
		So(jsonEnt.State, ShouldEqual, true)
		Convey("创建后验证", func() {
			_, body, _ := request.Get(Url + "role/" + Tools.ParseString(id)).End()
			var jsonEnt jsonRole
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, true)
      So(jsonEnt.Msg.Groupid, ShouldEqual, 0)
      So(jsonEnt.Msg.Name, ShouldEqual, "name")
      So(jsonEnt.Msg.Des, ShouldEqual, "des")
      So(jsonEnt.Msg.State, ShouldEqual, 0)
      
		})
	})
}

func TestRoleGet(t *testing.T) {
	Convey("获取用户组角色", t, func() {
		_, body, _ := request.Get(Url + "role/" + Tools.ParseString(id)).End()
		var jsonEnt jsonRole
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		So(jsonEnt.Msg.Id, ShouldEqual, id)
	})
}

func TestRoleUpdate(t *testing.T) {
	Convey("更新用户组角色", t, func() {
		_, body, _ := request.Put(Url + "role").Send(`{"id":`+Tools.ParseString(id)+`,"groupid": 1,"name": "name","des": "des","state": 1,"1":"1"}`).End()
		var jsonEnt jsonRole
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		Convey("更新后验证", func() {
			_, body, _ := request.Get(Url + "role/" + Tools.ParseString(id)).End()
			var jsonEnt jsonRole
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, true)
      So(jsonEnt.Msg.Groupid, ShouldEqual, 0)
      So(jsonEnt.Msg.Name, ShouldEqual, "name")
      So(jsonEnt.Msg.Des, ShouldEqual, "des")
      So(jsonEnt.Msg.State, ShouldEqual, 0)
      
		})
	})
}

func TestRolePage(t *testing.T) {
	Convey("获取列表用户组角色", t, func() {
		_, body, _ := request.Post(Url + "role/page").Send("cp=1&mp=10&kw=11").End()
		var jsonEnt jsonRoleList
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		So(jsonEnt.Count, ShouldNotEqual, 0)
	})
}

func TestRoleDele(t *testing.T) {
	Convey("删除用户组角色", t, func() {
		_, body, _ := request.Delete(Url + "role").Query("id=" + Tools.ParseString(id)).End()
		var jsonEnt jsonRole
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		Convey("获取验证", func() {
			_, body, _ := request.Get(Url + "role/" + Tools.ParseString(id)).End()
			var jsonEnt jsonRole
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, false)
		})
	})
}
