package controller

import (
	// "encoding/json"
	"encoding/json"
	"github.com/ivpusic/neo"
	"log"
	"pizzaCmsApi/logic"
	"pizzaCmsApi/model"
	"pizzaCmsApi/tools"
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
* @api {PUT} /node 更新node信息
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
	// id := tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	err := ctx.Req.JsonBody(&node)
	if err != nil {
		return 200, ctx.Res.Json(`{"state": false, msg:` + err.Error() + `}`)
	} else {
		err1 := validate.Struct(node)
		if err1 != nil {
			return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err1.Error()})
		} else {
			// node.ID = id
			log.Print("node=" + tools.StructToString(node))
			return 200, ctx.Res.Json(model.NodeUpdate(node))
		}
	}
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

/**
* @api {get} /node/pageall  page all node
* @apiName pageAll node
* @apiGroup node
* @apiVersion 1.0.0
* @apiDescription page node
* @apiSampleRequest /node/pageall
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func NodePageAll(ctx *neo.Ctx) (int, error) {
	return 200, ctx.Res.Json(model.NodePageAll())
}
