package controller

import (
	// "encoding/json"
	"github.com/ivpusic/neo"
	"pizzaCmsApi/logic"
	"pizzaCmsApi/model"
	"time"
)

/**
 * @api {get} /article/:id get article
 * @apiName 获取文章信息by path
 * @apiGroup article
 * @apiVersion 1.0.0
 * @apiDescription 用于后台管理员获取文章信息
 * @apiSampleRequest /article/:id
 * @apiParam {int} id文章id
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 消息
 */
func ArticleGet(ctx *neo.Ctx) (int, error) {
	id := Tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	return 200, ctx.Res.Json(model.ArticleGet(id))
}

/**
* @api {PUT} /article update article
* @apiName 更新article信息
* @apiGroup article
* @apiVersion 1.0.0
* @apiDescription 后台管理员更新文章信息
* @apiSampleRequest /article
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func ArticleUpdate(ctx *neo.Ctx) (int, error) {
	var article model.Article
	err := ctx.Req.JsonBody(&article)
	if err != nil {
		return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(article)
	if err1 != nil {
		return 200, ctx.Res.Json(`{"state": false, "msg": ` + err1.Error() + `}`)
	}
	return 200, ctx.Res.Json(model.ArticleUpdate(article))
}

/**
* @api {post} /article create article
* @apiName 创建article信息1
* @apiGroup article
* @apiVersion 1.0.0
* @apiDescription 创建文章信息
* @apiSampleRequest /article
* @apiParam {string} title title
* @apiParam {string} brief brief
* @apiParam {string} content content
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func ArticleCreate(ctx *neo.Ctx) (int, error) {
	var article model.Article
	err := ctx.Req.JsonBody(&article)
	if err != nil {
		return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(article)
	if err1 != nil {
		return 200, ctx.Res.Json(`{"state": false, "msg": ` + err1.Error() + `}`)
	}
	article.Createtime = time.Now().Unix()
	return 200, ctx.Res.Json(model.ArticleCreate(article))
}

/**
* @api {post} /article/page page article
* @apiName page article
* @apiGroup article
* @apiVersion 1.0.0
* @apiDescription page article
* @apiSampleRequest /article/page
* @apiParam {string} kw 关键字
* @apiParam {int} cp cp
* @apiParam {int} mp mp
* @apiParam {nodeid} nodeid 节点id
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func ArticlePage(ctx *neo.Ctx) (int, error) {
	cp := Tools.ParseInt(ctx.Req.FormValue("cp"), 1)
	mp := Tools.ParseInt(ctx.Req.FormValue("mp"), 20)
	nodeid := Tools.ParseInt(ctx.Req.FormValue("nodeid"), 0)
	kw := ctx.Req.FormValue("kw")

	return 200, ctx.Res.Json(model.ArticlePage(kw, nodeid, cp, mp))
}

/**
* @api {delete} /article delete article
* @apiName delete article
* @apiGroup article
* @apiVersion 1.0.0
* @apiDescription delete article by ids[]
* @apiSampleRequest /article
* @apiParam {string} id 文章id，可传多个用逗号隔开
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func ArticleDele(ctx *neo.Ctx) (int, error) {
	ids := ctx.Req.FormValue("id")
	return 200, ctx.Res.Json(logic.ArticleDele(ids))

}

/**
* @api {post} /article/pass pass article
* @apiName pass article
* @apiGroup article
* @apiVersion 1.0.0
* @apiDescription delete article by ids[]
* @apiSampleRequest /article/pass
* @apiParam {string} id 用户id
* @apiParam {string} pass  pass状态
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func ArticlePass(ctx *neo.Ctx) (int, error) {
	ids := ctx.Req.FormValue("id")
	pass := Tools.ParseInt(ctx.Req.FormValue("pass"), 0)
	return 200, ctx.Res.Json(logic.ArticlePass(ids, pass))

}
