package controller

import (
	// "encoding/json"
	"encoding/json"
	"github.com/ivpusic/neo"
	"log"
	"pizzaCmsApi/logic"
	"pizzaCmsApi/model"
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
	id := Tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	return 200, ctx.Res.Json(logic.NodeGet(id))
}

/**
* @api {PUT} /node 更新node信息
* @apiName 更新node信息
* @apiGroup node
* @apiVersion 1.0.0
* @apiDescription 后台管理员更新节点信息
* @apiSampleRequest /node/:id
* @apiParam {int} id 节点的id
* @apiParam {string} name 节点名称
* @apiParam {string} brief 节点描述
* @apiParam {string} link 节点连接地址
* @apiParam {int} weight 权重
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func NodeUpdate(ctx *neo.Ctx) (int, error) {
	var node model.Node
	// id := Tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	err := ctx.Req.JsonBody(&node)
	if err != nil {
		return 200, ctx.Res.Json(`{"state": false, msg:` + err.Error() + `}`)
	} else {
		err1 := validate.Struct(node)
		if err1 != nil {
			return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err1.Error()})
		} else {
			// node.ID = id
			log.Print("node=" + Tools.StructToString(node))
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
* @apiParam {string} name 节点名称
* @apiParam {string} brief 节点描述
* @apiParam {string} link 节点连接地址
* @apiParam {int} weight 权重
* @apiParam {int} pid 父节点id
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
* @api {post} /node/page 获取node列表
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
	pid := Tools.ParseInt(ctx.Req.FormValue("pid"), 0)
	return 200, ctx.Res.Json(model.NodePage(pid))
}

/**
* @api {get} /node/pageall  获取所有的节点
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

/**
* @api {get} /node/src  获取面包屑路径
* @apiName pageSrc node
* @apiGroup node
* @apiVersion 1.0.0
* @apiDescription 获取当前节点的面包屑路径
* @apiSampleRequest /node/src
* @apiParam {id} id 节点id
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func NodePageSrc(ctx *neo.Ctx) (int, error) {
	return 200, ctx.Res.Json(model.NodePageAll())
}
