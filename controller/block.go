package controller

import (
	"github.com/ivpusic/neo"
	"pizzaCmsApi/logic"
	"pizzaCmsApi/model"
)

/**
 * @api {get} /block/:id 获取模块
 * @apiName get block
 * @apiGroup block
 * @apiVersion 1.0.0
 * @apiDescription 获取模块信息
 * @apiSampleRequest /block/:id
 * @apiParam {int} id模块的id
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 消息
 * @apiSuccess {int} --id 主键id
 * @apiSuccess {string} --title 标题
 * @apiSuccess {string} --content 描述
 */
func BlockGet(ctx *neo.Ctx) (int, error) {
	id := Tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	err1 := validate.Field(id, "omitempty,min=1")
	if err1 != nil {
		return 200, ctx.Res.Json(errorValidate())
	}
	return 200, ctx.Res.Json(model.BlockGet(id))
}

/**
* @api {PUT} /block 更新模块
* @apiName update block
* @apiGroup block
* @apiVersion 1.0.0
* @apiDescription 更新模块信息
* @apiSampleRequest /block
* @apiParam {Block} block
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
* @apiName create block
* @apiGroup block
* @apiVersion 1.0.0
* @apiDescription 创建模块信息
* @apiSampleRequest /block
* @apiParam {Block} block 模块信息
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
	err1 := validate.Field(kw, "omitempty,max=20")
	err2 := validate.Field(cp, "required,min=1")
	err3 := validate.Field(mp, "required,min=1,max=50")
	if err1 != nil || err2 != nil || err3 != nil {
		return 200, ctx.Res.Json(errorValidate())
	}
	return 200, ctx.Res.Json(model.BlockPage(kw, cp, mp))
}

/**
* @api {delete} /block 删除模块
* @apiName delete block
* @apiGroup block
* @apiVersion 1.0.0
* @apiDescription delete block by id
* @apiSampleRequest /block
* @apiParam {string} id 模块id
* @apiParam {int} uid 用户的uid
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func BlockDele(ctx *neo.Ctx) (int, error) {
	ids := ctx.Req.FormValue("id")
	err1 := validate.Field(ids, "required,min=1")
	if err1 != nil {
		return 200, ctx.Res.Json(errorValidate())
	}
	return 200, ctx.Res.Json(logic.BlockDel(ids))

}
