package controller

import (
	"github.com/ivpusic/neo"
	"pizzaCmsApi/logic"
	"pizzaCmsApi/model"
)

/**
 * @api {get} /role/:id 获取用户组角色
 * @apiName get role
 * @apiGroup role
 * @apiVersion 1.0.0
 * @apiDescription 获取用户组角色信息
 * @apiSampleRequest /role/:id
 * @apiParam {int} id用户组角色的id
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 消息
 * @apiSuccess {int} --id 主键id
 * @apiSuccess {int} --groupid 用户组id
 * @apiSuccess {string} --name 角色名称
 * @apiSuccess {string} --des 角色描述
 * @apiSuccess {int} --state 角色状态
 */
func RoleGet(ctx *neo.Ctx) (int, error) {
	id := Tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	err1 := validate.Field(id, "omitempty,min=1")
	if err1 != nil {
		return 200, ctx.Res.Json(errorValidate())
	}
	return 200, ctx.Res.Json(model.RoleGet(id))
}

/**
* @api {PUT} /role 更新用户组角色
* @apiName update role
* @apiGroup role
* @apiVersion 1.0.0
* @apiDescription 更新用户组角色信息
* @apiSampleRequest /role
* @apiParam {Role} role
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func RoleUpdate(ctx *neo.Ctx) (int, error) {
	var role model.Role
	err := ctx.Req.JsonBody(&role)
	if err != nil {
		return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(role)
	if err1 != nil {
		return 200, ctx.Res.Json(`{"state": false, "msg": ` + err1.Error() + `}`)
	}
	return 200, ctx.Res.Json(logic.RoleUpdate(role))
}

/**
* @api {post} /role 创建用户组角色
* @apiName create role
* @apiGroup role
* @apiVersion 1.0.0
* @apiDescription 创建用户组角色信息
* @apiSampleRequest /role
* @apiParam {Role} role 用户组角色信息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func RoleCreate(ctx *neo.Ctx) (int, error) {
	var role model.Role
	err := ctx.Req.JsonBody(&role)
	if err != nil {
		return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(role)
	if err1 != nil {
		return 200, ctx.Res.Json(`{"state": false, "msg": ` + err1.Error() + `}`)
	}
	return 200, ctx.Res.Json(logic.RoleCreate(role))
}

/**
* @api {post} /role/page 获取用户组角色列表
* @apiName page role
* @apiGroup role
* @apiVersion 1.0.0
* @apiDescription page role
* @apiSampleRequest /role/page
* @apiParam {string} kw 关键字
* @apiParam {int} groupid groupid
* @apiParam {int} cp cp
* @apiParam {int} mp mp
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func RolePage(ctx *neo.Ctx) (int, error) {
	groupid := Tools.ParseInt(ctx.Req.FormValue("groupid"), 0)
	cp := Tools.ParseInt(ctx.Req.FormValue("cp"), 1)
	mp := Tools.ParseInt(ctx.Req.FormValue("mp"), 20)
	kw := ctx.Req.FormValue("kw")
	err1 := validate.Field(kw, "omitempty,min=1,max=20")
	err2 := validate.Field(cp, "required,min=1")
	err3 := validate.Field(mp, "required,min=1,max=50")
	err4 := validate.Field(groupid, "required,min=1")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return 200, ctx.Res.Json(errorValidate())
	}
	return 200, ctx.Res.Json(model.RolePage(groupid, kw, cp, mp))
}

/**
* @api {delete} /role 删除用户组角色
* @apiName delete role
* @apiGroup role
* @apiVersion 1.0.0
* @apiDescription delete role by id
* @apiSampleRequest /role
* @apiParam {string} id 用户组角色id
* @apiParam {int} uid 用户的uid
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func RoleDele(ctx *neo.Ctx) (int, error) {
	ids := Tools.ParseInt(ctx.Req.FormValue("id"), 0)
	err1 := validate.Field(ids, "required,min=1")
	if err1 != nil {
		return 200, ctx.Res.Json(errorValidate())
	}
	return 200, ctx.Res.Json(model.RoleDel(ids))

}
