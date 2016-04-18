package controller

import (
	"github.com/ivpusic/neo"
	"pizzaCmsApi/model"
	"pizzaCmsApi/logic"
)

/**
 * @api {get} /usergroup/:id 获取用户组
 * @apiName get usergroup
 * @apiGroup usergroup
 * @apiVersion 1.0.0
 * @apiDescription 获取用户组信息
 * @apiSampleRequest /usergroup/:id
 * @apiParam {int} id用户组的id
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 消息
 * @apiSuccess {int} --id 主键id
 * @apiSuccess {string} --name 用户组名称
 * @apiSuccess {string} --des 用户组描述
 * @apiSuccess {int} --state 用户组状态
 */
func UsergroupGet(ctx *neo.Ctx) (int, error) {
	id := Tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	err1 := validate.Field(id, "omitempty,min=1")
	if err1 != nil {
		return 200, ctx.Res.Json(errorValidate())
	}
	return 200, ctx.Res.Json(model.UsergroupGet(id))
}

/**
* @api {PUT} /usergroup 更新用户组
* @apiName update usergroup
* @apiGroup usergroup
* @apiVersion 1.0.0
* @apiDescription 更新用户组信息
* @apiSampleRequest /usergroup
* @apiParam {Usergroup} usergroup
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UsergroupUpdate(ctx *neo.Ctx) (int, error) {
	var usergroup model.Usergroup
	err := ctx.Req.JsonBody(&usergroup)
	if err != nil {
		return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(usergroup)
	if err1 != nil {
		return 200, ctx.Res.Json(`{"state": false, "msg": ` + err1.Error() + `}`)
	}
	return 200, ctx.Res.Json(logic.UsergroupUpdate(usergroup))
}

/**
* @api {post} /usergroup 创建用户组
* @apiName create usergroup
* @apiGroup usergroup
* @apiVersion 1.0.0
* @apiDescription 创建用户组信息
* @apiSampleRequest /usergroup
* @apiParam {Usergroup} usergroup 用户组信息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UsergroupCreate(ctx *neo.Ctx) (int, error) {
	var usergroup model.Usergroup
	err := ctx.Req.JsonBody(&usergroup)
	if err != nil {
		return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(usergroup)
	if err1 != nil {
		return 200, ctx.Res.Json(`{"state": false, "msg": ` + err1.Error() + `}`)
	}
	return 200, ctx.Res.Json(logic.UsergroupCreate(usergroup))
}

/**
* @api {post} /usergroup/page 获取用户组列表
* @apiName page usergroup
* @apiGroup usergroup
* @apiVersion 1.0.0
* @apiDescription page usergroup
* @apiSampleRequest /usergroup/page
* @apiParam {string} kw 关键字
* @apiParam {int} cp cp
* @apiParam {int} mp mp
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UsergroupPage(ctx *neo.Ctx) (int, error) {
	cp := Tools.ParseInt(ctx.Req.FormValue("cp"), 1)
	mp := Tools.ParseInt(ctx.Req.FormValue("mp"), 20)
	kw := ctx.Req.FormValue("kw")
	err1 := validate.Field(kw, "omitempty,min=1,max=20")
	err2 := validate.Field(cp, "required,min=1")
	err3 := validate.Field(mp, "required,min=1,max=50")
	if err1 != nil || err2 != nil || err3 != nil {
		return 200, ctx.Res.Json(errorValidate())
	}
	return 200, ctx.Res.Json(model.UsergroupPage(kw, cp, mp))
}

/**
* @api {delete} /usergroup 删除用户组
* @apiName delete usergroup
* @apiGroup usergroup
* @apiVersion 1.0.0
* @apiDescription delete usergroup by id
* @apiSampleRequest /usergroup
* @apiParam {string} id 用户组id
* @apiParam {int} uid 用户的uid
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func UsergroupDele(ctx *neo.Ctx) (int, error) {
	ids := Tools.ParseInt(ctx.Req.FormValue("id"), 0)
	err1 := validate.Field(ids, "required,min=1")
	if err1 != nil {
		return 200, ctx.Res.Json(errorValidate())
	}
	return 200, ctx.Res.Json(model.UsergroupDel(ids))

}
