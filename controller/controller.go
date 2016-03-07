package controller

import (
	"gopkg.in/go-playground/validator.v8"
)

var (
	validate *validator.Validate
)


func init() {
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)
}



//////////私有方法
/**
 * 返回数据格式不合法的字符串
 * @method ErrorValidate
 */
func errorValidate() string {
		return `{"state": false, "msg": "数据格式不合法"}`
}
