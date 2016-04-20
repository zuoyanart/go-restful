package restest

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"pizzaCmsApi/model"
	"testing"
)

type jsonActionmenu struct {
	State bool
	Msg   model.Actionmenu
}

type jsonActionmenuGet struct {
	State bool
	Msg   int
}

type jsonActionmenuList struct {
	State bool
	Msg   []model.Actionmenu
	Count int
}

func TestActionmenuCreate(t *testing.T) {
	Convey("创建部门栏目", t, func() {
		_, body, _ := request.Post(Url + "power").Send(`{"name": "name","link": "link","state": 1,"pid": 1,"icon": "icon","1":"1"}`).End()
		var jsonEnt jsonActionmenuGet
		json.Unmarshal([]byte(body), &jsonEnt) //转换成struct
		id = jsonEnt.Msg                       //赋值id
		So(jsonEnt.State, ShouldEqual, true)
		Convey("创建后验证", func() {
			_, body, _ := request.Get(Url + "power/" + Tools.ParseString(id)).End()
			var jsonEnt jsonActionmenu
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, true)
      So(jsonEnt.Msg.Name, ShouldEqual, "name")
      So(jsonEnt.Msg.Link, ShouldEqual, "link")
      So(jsonEnt.Msg.State, ShouldEqual, 0)
      So(jsonEnt.Msg.Pid, ShouldEqual, 0)
      So(jsonEnt.Msg.Icon, ShouldEqual, "icon")
      
		})
	})
}

func TestActionmenuGet(t *testing.T) {
	Convey("获取部门栏目", t, func() {
		_, body, _ := request.Get(Url + "power/" + Tools.ParseString(id)).End()
		var jsonEnt jsonActionmenu
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		So(jsonEnt.Msg.Id, ShouldEqual, id)
	})
}

func TestActionmenuUpdate(t *testing.T) {
	Convey("更新部门栏目", t, func() {
		_, body, _ := request.Put(Url + "power").Send(`{"id":`+Tools.ParseString(id)+`,"name": "name","link": "link","state": 1,"pid": 1,"icon": "icon","1":"1"}`).End()
		var jsonEnt jsonActionmenu
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		Convey("更新后验证", func() {
			_, body, _ := request.Get(Url + "power/" + Tools.ParseString(id)).End()
			var jsonEnt jsonActionmenu
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, true)
      So(jsonEnt.Msg.Name, ShouldEqual, "name")
      So(jsonEnt.Msg.Link, ShouldEqual, "link")
      So(jsonEnt.Msg.State, ShouldEqual, 0)
      So(jsonEnt.Msg.Pid, ShouldEqual, 0)
      So(jsonEnt.Msg.Icon, ShouldEqual, "icon")
      
		})
	})
}

func TestActionmenuPage(t *testing.T) {
	Convey("获取列表部门栏目", t, func() {
		_, body, _ := request.Post(Url + "power/page").Send("cp=1&mp=10&kw=11").End()
		var jsonEnt jsonActionmenuList
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		So(jsonEnt.Count, ShouldNotEqual, 0)
	})
}

func TestActionmenuDele(t *testing.T) {
	Convey("删除部门栏目", t, func() {
		_, body, _ := request.Delete(Url + "power").Query("id=" + Tools.ParseString(id)).End()
		var jsonEnt jsonActionmenu
		json.Unmarshal([]byte(body), &jsonEnt)
		So(jsonEnt.State, ShouldEqual, true)
		Convey("获取验证", func() {
			_, body, _ := request.Get(Url + "power/" + Tools.ParseString(id)).End()
			var jsonEnt jsonActionmenu
			json.Unmarshal([]byte(body), &jsonEnt)
			So(jsonEnt.State, ShouldEqual, false)
		})
	})
}
