package controller

import (
	"github.com/ivpusic/neo"
	"pizzaCmsApi/model"
)

/**
 * @api {get} /power/:id 获取部门栏目
 * @apiName get actionmenu
 * @apiGroup actionmenu
 * @apiVersion 1.0.0
 * @apiDescription 获取部门栏目信息
 * @apiSampleRequest /power/:id
 * @apiParam {int} id部门栏目的id
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 消息
 * @apiSuccess {int} --id 主键id
 * @apiSuccess {string} --name 栏目模块名称
 * @apiSuccess {string} --link 栏目模块编码
 * @apiSuccess {int} --state 栏目模块状态
 * @apiSuccess {int} --pid 父节点
 * @apiSuccess {string} --icon 图标
 */
func ActionmenuGet(ctx *neo.Ctx) (int, error) {
	id := Tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	err1 := validate.Field(id, "omitempty,min=1")
	if err1 != nil {
		return 200, ctx.Res.Json(errorValidate())
	}
	return 200, ctx.Res.Json(model.ActionmenuGet(id))
}

/**
* @api {PUT} /power 更新部门栏目
* @apiName update actionmenu
* @apiGroup actionmenu
* @apiVersion 1.0.0
* @apiDescription 更新部门栏目信息
* @apiSampleRequest /power
* @apiParam {Actionmenu} actionmenu
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func ActionmenuUpdate(ctx *neo.Ctx) (int, error) {
	var actionmenu model.Actionmenu
	err := ctx.Req.JsonBody(&actionmenu)
	if err != nil {
		return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(actionmenu)
	if err1 != nil {
		return 200, ctx.Res.Json(`{"state": false, "msg": ` + err1.Error() + `}`)
	}
	return 200, ctx.Res.Json(model.ActionmenuUpdate(actionmenu))
}

/**
* @api {post} /power 创建部门栏目
* @apiName create actionmenu
* @apiGroup actionmenu
* @apiVersion 1.0.0
* @apiDescription 创建部门栏目信息
* @apiSampleRequest /power
* @apiParam {Actionmenu} actionmenu 部门栏目信息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func ActionmenuCreate(ctx *neo.Ctx) (int, error) {
	var actionmenu model.Actionmenu
	err := ctx.Req.JsonBody(&actionmenu)
	if err != nil {
		return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(actionmenu)
	if err1 != nil {
		return 200, ctx.Res.Json(`{"state": false, "msg": ` + err1.Error() + `}`)
	}
	return 200, ctx.Res.Json(model.ActionmenuCreate(actionmenu))
}

/**
* @api {post} /power/page 获取部门栏目列表
* @apiName page actionmenu
* @apiGroup actionmenu
* @apiVersion 1.0.0
* @apiDescription page actionmenu
* @apiSampleRequest /power/page
* @apiParam {int} pid 父节点id
* @apiParam {int} cp cp
* @apiParam {int} mp mp
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func ActionmenuPage(ctx *neo.Ctx) (int, error) {
	cp := Tools.ParseInt(ctx.Req.FormValue("cp"), 1)
	mp := Tools.ParseInt(ctx.Req.FormValue("mp"), 20)
	pid := Tools.ParseInt(ctx.Req.FormValue("pid"), 0)
	err1 := validate.Field(pid, "omitempty,min=1")
	err2 := validate.Field(cp, "required,min=1")
	err3 := validate.Field(mp, "required,min=1,max=50")
	if err1 != nil || err2 != nil || err3 != nil {
		return 200, ctx.Res.Json(errorValidate())
	}
	return 200, ctx.Res.Json(model.ActionmenuPage(pid, cp, mp))
}

/**
* @api {delete} /power 删除部门栏目
* @apiName delete actionmenu
* @apiGroup actionmenu
* @apiVersion 1.0.0
* @apiDescription delete actionmenu by id
* @apiSampleRequest /power
* @apiParam {string} id 部门栏目id
* @apiParam {int} uid 用户的uid
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func ActionmenuDele(ctx *neo.Ctx) (int, error) {
	ids := Tools.ParseInt(ctx.Req.FormValue("id"), 0)
	uid := Tools.ParseInt(ctx.Req.FormValue("uid"), 0)
	err1 := validate.Field(ids, "required,min=1")
	err2 := validate.Field(uid, "required,min=1")
	if err1 != nil || err2 != nil {
		return 200, ctx.Res.Json(errorValidate())
	}
	return 200, ctx.Res.Json(model.ActionmenuDel(ids, uid))

}
