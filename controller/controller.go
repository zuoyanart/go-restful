package controller

import (
	"gopkg.in/go-playground/validator.v8"
	"pizzaCmsApi/tools"

)

var (
	validate *validator.Validate
	Tools    *tools.Tools
)

func init() {
	valiconf := &validator.Config{TagName: "validate"}
	validate = validator.New(valiconf)
}

//////////私有方法
/**
 * 返回数据格式不合法的字符串
 * @method ErrorValidate
 */
func errorValidate() string {
	return `{"state": false, "msg": "数据格式不合法"}`
}
