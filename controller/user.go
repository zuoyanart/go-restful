package controller

import (
	// "encoding/json"
	"github.com/ivpusic/neo"
	"pizzaCmsApi/logic"
	"pizzaCmsApi/model"
	"pizzaCmsApi/tools"
)

/**
 * @api {get} /user?id=:id 获取用户信息
 * @apiName 获取用户信息
 * @apiGroup user
 * @apiVersion 1.0.0
 * @apiDescription 用于后台管理员获取用户信息
 * @apiSampleRequest /user
 * @apiParam {int} id用户id
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 消息内容
 */
func UserGet(ctx *neo.Ctx) (int, error) {
	id := tools.ParseInt(ctx.Req.FormValue("id"), 0)
	return 200, ctx.Res.Json(model.UserGet(id))
}

/**
 * @api {post} /user/login 判断用户是否登录
 * @apiName 判断用户是否登录
 * @apiGroup user
 * @apiVersion 1.0.0
 * @apiDescription 用于后台管理员获取用户信息
 * @apiSampleRequest /user/login
 * @apiParam {string} username username
 * @apiParam {string} password password
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 用户信息
 */
func UserCheckLogin(ctx *neo.Ctx) (int, error) {
	username := ctx.Req.FormValue("username")
	password := ctx.Req.FormValue("password")
	err1 := validate.Field(username, "required,min=4,max=20")
	err2 := validate.Field(password, "required,min=6,max=20")
	if err1 != nil || err2 != nil {
		return 200, ctx.Res.Json(errorValidate())
	}
	return 200, ctx.Res.Json(logic.UserCheckLogin(username, password))
}

/**
 * @api {get} /user/:id 获取用户信息
 * @apiName 获取用户信息by path
 * @apiGroup user
 * @apiVersion 1.0.0
 * @apiDescription 用于后台管理员获取用户信息
 * @apiSampleRequest /user/:id
 * @apiParam {int} id用户id
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 消息
 */
func UserGetByPath(ctx *neo.Ctx) (int, error) {
	id := tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	return 200, ctx.Res.Json(model.UserGet(id))
}

/**
* @api {PUT} /user/:id 更新user信息
* @apiName 更新user信息
* @apiGroup user
* @apiVersion 1.0.0
* @apiDescription 后台管理员更新用户信息
* @apiSampleRequest /user/:id
* @apiParam {string} username 用户名
* @apiParam {int} id 用户的id
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UserUpdateName(ctx *neo.Ctx) (int, error) {
	var user model.User
	// err := ctx.Req.JsonBody(&user)
	// if err != nil {
	//   	return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error() })
	// }

	// b, err := json.Marshal(user)
	// if err != nil {
	// 	log.Printf("json err:", err)
	// } else {
	// 	log.Printf(string(b))
	// }

	id := tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	user.ID = id
	user.Username = ctx.Req.FormValue("username")
	return 200, ctx.Res.Json(model.UserUpdateName(user))
}

/**
* @api {post} /user/ 创建user信息1
* @apiName 创建user信息1
* @apiGroup user
* @apiVersion 1.0.0
* @apiDescription 创建用户信息
* @apiSampleRequest /user/
* @apiParam {string} username 用户名
* @apiParam {string} password 密码
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UserCreate(ctx *neo.Ctx) (int, error) {
	var user model.User

	user.Username = ctx.Req.FormValue("username")
	user.Password = ctx.Req.FormValue("password")
	user.State = 3
	err := validate.Struct(user)
	if err != nil {
		return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error()})
	}
	return 200, ctx.Res.Json(model.UserCreate(user))
}

/**
* @api {post} /user/page page user
* @apiName page user
* @apiGroup user
* @apiVersion 1.0.0
* @apiDescription page user
* @apiSampleRequest /user/page
* @apiParam {string} kw 关键字
* @apiParam {int} cp cp
* @apiParam {int} mp mp
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UserPage(ctx *neo.Ctx) (int, error) {
	cp := tools.ParseInt(ctx.Req.FormValue("cp"), 1)
	mp := tools.ParseInt(ctx.Req.FormValue("mp"), 20)
	kw := ctx.Req.FormValue("kw")

	return 200, ctx.Res.Json(model.UserPage(kw, cp, mp))
}
