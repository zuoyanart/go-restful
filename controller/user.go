package controller

import (
	// "encoding/json"
	"github.com/ivpusic/neo"
	"pizzaCmsApi/logic"
	"pizzaCmsApi/model"
	"strings"
)

/**
 * @api {post} /user/login 用户登录
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
 * @apiSuccess {String} msg 消息， msg:{id:id,username:username,nickname:nickname,state:state},密码和盐不返回
 */
func UserGetByPath(ctx *neo.Ctx) (int, error) {
	id := Tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	return 200, ctx.Res.Json(model.UserGet(id))
}

/**
* @api {PUT} /user 更新用户信息
* @apiName 更新user信息
* @apiGroup user
* @apiVersion 1.0.0
* @apiDescription 后台管理员更新用户信息，如果传入密码，则更新密码，如果不传，则不更新
* @apiSampleRequest /user
* @apiParam {string} username 用户名
* @apiParam {string} nickname 昵称
* @apiParam {string} password 密码
* @apiParam {int} id 用户的id
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UserUpdate(ctx *neo.Ctx) (int, error) {
	var user model.User
	err := ctx.Req.JsonBody(&user)
	if err != nil {
			return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error() })
	} else {
		err1 := validate.Struct(user)
		if err1!= nil {
			return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err1.Error() })
		}
		return 200, ctx.Res.Json(model.UserUpdate(user))
	}

}

/**
* @api {post} /user  创建用户
* @apiGroup user
* @apiVersion 1.0.0
* @apiDescription 创建用户信息
* @apiSampleRequest /user
* @apiParam {string} username 用户名
* @apiParam {string} nickname 昵称
* @apiParam {string} password 密码
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UserCreate(ctx *neo.Ctx) (int, error) {
var user model.User
	err := ctx.Req.JsonBody(&user)
	if err != nil {
			return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error() })
	}
	err1 := validate.Struct(user)
	if err1 != nil {
		return 200, ctx.Res.Json(`{"state": false, "msg": ` + err1.Error() + `}`)
	}
	return 200, ctx.Res.Json(model.UserCreate(user))
}

/**
* @api {post} /user/page 获取列表
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
	cp := Tools.ParseInt(ctx.Req.FormValue("cp"), 1)
	mp := Tools.ParseInt(ctx.Req.FormValue("mp"), 20)
	kw := ctx.Req.FormValue("kw")

	return 200, ctx.Res.Json(model.UserPage(kw, cp, mp))
}

/**
* @api {delete} /user 删除用户
* @apiName delete user
* @apiGroup user
* @apiVersion 1.0.0
* @apiDescription delete user by ids[]
* @apiSampleRequest /user
* @apiParam {string} id 用户id，可批量
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UserDele(ctx *neo.Ctx) (int, error) {
	ids := ctx.Req.FormValue("id")
	if strings.Index("," + ids + ",", ",1,") == -1 {
			return 200, ctx.Res.Json(logic.UserDele(ids))
	} else {
			return 200, ctx.Res.Json(`{"state": false, "msg": "root is not del"}`)
	}
}
