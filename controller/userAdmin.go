package controller

import (
	"strings"

	"pizzaCmsApi/logic"
	"pizzaCmsApi/model"

	"github.com/kataras/iris/v12"
)

/**
 * @api {get} /useradmin?id=:id get useradmin
 * @apiName 获取用户信息
 * @apiGroup useradmin
 * @apiVersion 1.0.0
 * @apiDescription 用于后台管理员获取用户信息
 * @apiSampleRequest /useradmin
 * @apiParam {int} id用户id
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 消息内容
 */
func UserAdminGet(ctx iris.Context) {
	id := ctx.Params().GetIntDefault("id", 0)
	ctx.JSON(model.UserAdminGet(id))
}

/**
 * @api {post} /useradmin/login useradmin login
 * @apiName 判断用户是否登录
 * @apiGroup useradmin
 * @apiVersion 1.0.0
 * @apiDescription 用于后台管理员获取用户信息
 * @apiSampleRequest /useradmin/login
 * @apiParam {string} username username
 * @apiParam {string} password password
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 用户信息
 */
func UserAdminCheckLogin(ctx iris.Context) {
	username := ctx.Params().Get("username")
	password := ctx.Params().Get("password")
	err1 := validate.Var(username, "required,min=4,max=20")
	err2 := validate.Var(password, "required,min=6,max=20")
	if err1 != nil || err2 != nil {
		ctx.JSON(errorValidate())
		return
	}

	ctx.JSON(logic.UserAdminCheckLogin(username, password))
}

/**
 * @api {get} /useradmin/:id get useradmin
 * @apiName 获取用户信息by path
 * @apiGroup useradmin
 * @apiVersion 1.0.0
 * @apiDescription 用于后台管理员获取用户信息
 * @apiSampleRequest /useradmin/:id
 * @apiParam {int} id用户id
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 消息
 */
func UserAdminGetByPath(ctx iris.Context) {
	id := ctx.Params().GetIntDefault("id", 0)
	ctx.JSON(model.UserAdminGet(id))
}

/**
* @api {PUT} /useradmin update useradmin info
* @apiName 更新useradmin信息
* @apiGroup useradmin
* @apiVersion 1.0.0
* @apiDescription 后台管理员更新用户信息，如果传入密码，则更新密码，如果不传，则不更新
* @apiSampleRequest /useradmin
* @apiParam {string} username 用户名
* @apiParam {string} nickname 昵称
* @apiParam {string} password 密码
* @apiParam {int} id 用户的id
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UserAdminUpdate(ctx iris.Context) {
	var user model.UserAdmin
	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.JSON(model.ApiJson{State: false, Msg: err.Error()})
		return
	}

	err1 := validate.Struct(user)
	if err1 != nil {
		ctx.JSON(model.ApiJson{State: false, Msg: err1.Error()})
		return
	}

	ctx.JSON(model.UserAdminUpdate(user))
}

/**
* @api {post} /useradmin create useradmin
* @apiName 创建useradmin信息1
* @apiGroup useradmin
* @apiVersion 1.0.0
* @apiDescription 创建用户信息
* @apiSampleRequest /useradmin
* @apiParam {string} username 用户名
* @apiParam {string} password 密码
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UserAdminCreate(ctx iris.Context) {
	var user model.UserAdmin
	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.JSON(model.ApiJson{State: false, Msg: err.Error()})
		return
	}
	err1 := validate.Struct(user)
	if err1 != nil {
		ctx.JSON(`{"state": false, "msg": ` + err1.Error() + `}`)
		return
	}
	ctx.JSON(model.UserAdminCreate(user))
}

/**
* @api {post} /useradmin/page page useradmin
* @apiName page useradmin
* @apiGroup useradmin
* @apiVersion 1.0.0
* @apiDescription page useradmin
* @apiSampleRequest /useradmin/page
* @apiParam {string} kw 关键字
* @apiParam {int} cp cp
* @apiParam {int} mp mp
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UserAdminPage(ctx iris.Context) {
	cp := ctx.Params().GetIntDefault("cp", 1)
	mp := ctx.Params().GetIntDefault("mp", 20)
	kw := ctx.Params().Get("kw")

	ctx.JSON(model.UserAdminPage(kw, cp, mp))
}

/**
* @api {delete} /useradmin delete UserAdmin
* @apiName delete useradmin
* @apiGroup useradmin
* @apiVersion 1.0.0
* @apiDescription delete useradmin by ids[]
* @apiSampleRequest /useradmin
* @apiParam {string} id 用户id
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UserAdminDele(ctx iris.Context) {
	ids := ctx.Params().Get("id")
	if strings.Index(","+ids+",", ",1,") == -1 {
		ctx.JSON(logic.UserAdminDele(ids))
	} else {
		ctx.JSON(`{"state": false, "msg": "root is not del"}`)
	}
}
