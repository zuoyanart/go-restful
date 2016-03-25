package controller

import (
	"github.com/ivpusic/neo"
	"pizzaCmsApi/model"
	"pizzaCmsApi/logic"
)

/**
 * @api {get} /block/:id 获取模块
 * @apiName 获取模块信息by path
 * @apiGroup block
 * @apiVersion 1.0.0
 * @apiDescription 用于后台管理员获取模块信息
 * @apiSampleRequest /block/:id
 * @apiParam {int} id模块id
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 消息
 */
func BlockGet(ctx *neo.Ctx) (int, error) {
	id := Tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	return 200, ctx.Res.Json(model.BlockGet(id))
}

/**
* @api {PUT} /block 更新模块
* @apiName 更新block信息
* @apiGroup block
* @apiVersion 1.0.0
* @apiDescription 后台管理员更新模块信息
* @apiSampleRequest /block
* @apiParam {int} id 模块id
* @apiParam {string} title 模块名称
* @apiParam {string} content 模块内容
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func BlockUpdate(ctx *neo.Ctx) (int, error) {
	var block model.Block
	err := ctx.Req.JsonBody(&block)
	if err != nil {
		return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(block)
	if err1 != nil {
		return 200, ctx.Res.Json(`{"state": false, "msg": ` + err1.Error() + `}`)
	}
	return 200, ctx.Res.Json(model.BlockUpdate(block))
}

/**
* @api {post} /block 创建模块
* @apiName 创建block信息1
* @apiGroup block
* @apiVersion 1.0.0
* @apiDescription 创建模块信息
* @apiSampleRequest /block
* @apiParam {string} title 模块名称
* @apiParam {string} content 模块内容
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func BlockCreate(ctx *neo.Ctx) (int, error) {
	var block model.Block
	err := ctx.Req.JsonBody(&block)
	if err != nil {
		return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(block)
	if err1 != nil {
		return 200, ctx.Res.Json(`{"state": false, "msg": ` + err1.Error() + `}`)
	}
	return 200, ctx.Res.Json(model.BlockCreate(block))
}

/**
* @api {post} /block/page 获取模块列表
* @apiName page block
* @apiGroup block
* @apiVersion 1.0.0
* @apiDescription page block
* @apiSampleRequest /block/page
* @apiParam {string} kw 关键字
* @apiParam {int} cp cp
* @apiParam {int} mp mp
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func BlockPage(ctx *neo.Ctx) (int, error) {
	cp := Tools.ParseInt(ctx.Req.FormValue("cp"), 1)
	mp := Tools.ParseInt(ctx.Req.FormValue("mp"), 20)
	kw := ctx.Req.FormValue("kw")
	return 200, ctx.Res.Json(model.BlockPage(kw, cp, mp))
}

/**
* @api {delete} /block 删除模块
* @apiName delete block
* @apiGroup block
* @apiVersion 1.0.0
* @apiDescription delete block by ids[]
* @apiSampleRequest /block
* @apiParam {string} id 模块id，可传多个用逗号隔开
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func BlockDele(ctx *neo.Ctx) (int, error) {
	ids := ctx.Req.FormValue("id")
	return 200, ctx.Res.Json(logic.BlockDele(ids))

}
