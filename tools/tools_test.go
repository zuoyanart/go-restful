package tools

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var Ts = New()

func TestNew(t *testing.T) {
	Convey("生成工具类的指针", t, func() {
		Convey("第二次生成", func() {
			So(New(), ShouldEqual, Ts)
		})
	})
}

func TestParseInt(t *testing.T) {
	Convey("测试parseInt", t, func() {
		Convey("为字符串型的数字", func() {
			So(Ts.ParseInt("1", 0), ShouldEqual, 1)
		})

		Convey("为字符串型的非数字", func() {
			So(Ts.ParseInt("aaa", 0), ShouldEqual, 0)
		})

		Convey("为字符串型的浮点数", func() {
			So(Ts.ParseInt("5.6", 0), ShouldEqual, 0)
		})
	})
}

func TestParseString(t *testing.T) {
	Convey("测试int转换为string，parseInt", t, func() {
		So(Ts.ParseString(1), ShouldEqual, "1")
	})
}

func TestMD5(t *testing.T) {
	Convey("测试md5", t, func() {
		So(Ts.MD5("1"), ShouldEqual, "c4ca4238a0b923820dcc509a6f75849b")
	})
}

func TestStructToString(t *testing.T) {
	Convey("struct to string", t, func() {
		type Info struct {
			Id   int
			Name string
			Des  string
		}
		// var infoE = Info{Id: 1, Name: "name", Des: "des"}
		So(Ts.StructToString(Info{Id: 1, Name: "name", Des: "des"}), ShouldContainSubstring, `{"Id":1`)
	})
}

func TestGetRandomString(t *testing.T) {
	Convey("测试随机数生成", t, func() {
		var s1 = Ts.GetRandomString(5)
		So(Ts.GetRandomString(5), ShouldNotEqual, s1)
		So(len(Ts.GetRandomString(5)), ShouldEqual, 5)
		So(Ts.GetRandomString(5), ShouldNotEqual, s1)
	})
}
