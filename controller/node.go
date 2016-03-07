package controller

import (
	// "encoding/json"
	"github.com/ivpusic/neo"
	"pizzaCmsApi/logic"
	"pizzaCmsApi/model"
	"pizzaCmsApi/tools"
	 "encoding/json"
	 "log"
)

/**
 * @api {get} /node/:id 获取节点信息
 * @apiName 获取节点信息by path
 * @apiGroup node
 * @apiVersion 1.0.0
 * @apiDescription 用于后台管理员获取节点信息
 * @apiSampleRequest /node/:id
 * @apiParam {int} id节点id
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 消息
 */
func NodeGet(ctx *neo.Ctx) (int, error) {
	id := tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	return 200, ctx.Res.Json(logic.NodeGet(id))
}

/**
* @api {PUT} /node/:id 更新node信息
* @apiName 更新node信息
* @apiGroup node
* @apiVersion 1.0.0
* @apiDescription 后台管理员更新节点信息
* @apiSampleRequest /node/:id
* @apiParam {string} title 节点名
* @apiParam {int} id 节点的id
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func NodeUpdate(ctx *neo.Ctx) (int, error) {
	var node model.Node
	id := tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	node.ID = id
	node.Name = ctx.Req.FormValue("title")
	return 200, ctx.Res.Json(model.NodeUpdate(node))
}

/**
* @api {post} /node/ 创建node
* @apiName 创建node
* @apiGroup node
* @apiVersion 1.0.0
* @apiDescription 创建节点信息
* @apiParam {string} name name
* @apiParam {string} brief brief
* @apiParam {int} pid pid
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func NodeCreate(ctx *neo.Ctx) (int, error) {
	var node model.Node
	err := ctx.Req.JsonBody(&node)
	if err != nil {
		return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(node)
	if err != nil {
		return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err1.Error()})
	}
	b, err := json.Marshal(node)
	if err != nil {
		log.Printf("json err:", err)
	} else {
		log.Printf(string(b))
	}
	//node.Name = ctx.Req.FormValue("name")
	//node.Brief = ctx.Req.FormValue("brief")
	//node.Pid = tools.ParseInt(ctx.Req.FormValue("pid"), 1)
	return 200, ctx.Res.Json(logic.NodeCreate(node))
}

/**
* @api {post} /node/page page node
* @apiName page node
* @apiGroup node
* @apiVersion 1.0.0
* @apiDescription page node
* @apiSampleRequest /node/page
* @apiParam {pid} pid 节点id
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func NodePage(ctx *neo.Ctx) (int, error) {
	pid := tools.ParseInt(ctx.Req.FormValue("pid"), 0)
	return 200, ctx.Res.Json(model.NodePage(pid))
}
