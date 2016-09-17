package controller

import (
	// "encoding/json"
	"github.com/kataras/iris"
	"pizzaCmsApi/logic"
	"pizzaCmsApi/model"
	"strings"
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
func UserAdminGet(ctx *iris.Context) {
	id := Tools.ParseInt(ctx.Param("id"), 0)
	ctx.JSON(iris.StatusOK, model.UserAdminGet(id))
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
func UserAdminCheckLogin(ctx *iris.Context) {
	username := ctx.Param("username")
	password := ctx.Param("password")
	err1 := validate.Var(username, "required,min=4,max=20")
	err2 := validate.Var(password, "required,min=6,max=20")
	if err1 != nil || err2 != nil {
		ctx.JSON(iris.StatusOK, errorValidate())
	}
	ctx.JSON(iris.StatusOK, logic.UserAdminCheckLogin(username, password))
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
func UserAdminGetByPath(ctx *iris.Context) {
	id := Tools.ParseInt(ctx.Param("id"), 0)
	ctx.JSON(iris.StatusOK, model.UserAdminGet(id))
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
func UserAdminUpdate(ctx *iris.Context) {
	var user model.UserAdmin
	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.JSON(iris.StatusOK, model.ApiJson{State: false, Msg: err.Error()})
	} else {
		err1 := validate.Struct(user)
		if err1 != nil {
			ctx.JSON(iris.StatusOK, model.ApiJson{State: false, Msg: err1.Error()})
		}
		ctx.JSON(iris.StatusOK, model.UserAdminUpdate(user))
	}

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
func UserAdminCreate(ctx *iris.Context) {
	var user model.UserAdmin
	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.JSON(iris.StatusOK, model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(user)
	if err1 != nil {
		ctx.JSON(iris.StatusOK, `{"state": false, "msg": `+err1.Error()+`}`)
	}
	ctx.JSON(iris.StatusOK, model.UserAdminCreate(user))
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
func UserAdminPage(ctx *iris.Context) {
	cp := Tools.ParseInt(ctx.Param("cp"), 1)
	mp := Tools.ParseInt(ctx.Param("mp"), 20)
	kw := ctx.Param("kw")

	ctx.JSON(iris.StatusOK, model.UserAdminPage(kw, cp, mp))
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
func UserAdminDele(ctx *iris.Context) {
	ids := ctx.Param("id")
	if strings.Index(","+ids+",", ",1,") == -1 {
		ctx.JSON(iris.StatusOK, logic.UserAdminDele(ids))
	} else {
		ctx.JSON(iris.StatusOK, `{"state": false, "msg": "root is not del"}`)
	}
}
