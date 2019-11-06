package controller

import (
	"time"

	"pizzaCmsApi/logic"
	"pizzaCmsApi/model"

	"github.com/kataras/iris/v12"
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
func ArticleGet(ctx iris.Context) {
	id := ctx.Params().GetIntDefault("id", 0)
	ctx.JSON(model.ArticleGet(id))
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
func ArticleUpdate(ctx iris.Context) {
	var article model.Article
	err := ctx.ReadJSON(&article)
	if err != nil {
		ctx.JSON(model.ApiJson{State: false, Msg: err.Error()})
		return
	}

	err1 := validate.Struct(article)
	if err1 != nil {
		ctx.JSON(`{"state": false, "msg": ` + err1.Error() + `}`)
		return
	}
	ctx.JSON(model.ArticleUpdate(article))
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
func ArticleCreate(ctx iris.Context) {
	var article model.Article
	err := ctx.ReadJSON(&article)
	if err != nil {
		ctx.JSON(model.ApiJson{State: false, Msg: err.Error()})
		return
	}
	err1 := validate.Struct(article)
	if err1 != nil {
		ctx.JSON(`{"state": false, "msg": ` + err1.Error() + `}`)
		return
	}
	article.Createtime = time.Now().Unix()
	ctx.JSON(model.ArticleCreate(article))
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
func ArticlePage(ctx iris.Context) {
	cp := ctx.Params().GetIntDefault("cp", 1)
	mp := ctx.Params().GetIntDefault("mp", 20)
	nodeid := ctx.Params().GetIntDefault("nodeid", 0)
	kw := ctx.Params().Get("kw")

	ctx.JSON(model.ArticlePage(kw, nodeid, cp, mp))
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
func ArticleDele(ctx iris.Context) {
	ids := ctx.Params().Get("id")
	ctx.JSON(logic.ArticleDele(ids))
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
func ArticlePass(ctx iris.Context) {
	ids := ctx.Params().Get("id")
	pass := ctx.Params().GetIntDefault("pass", 0)
	ctx.JSON(logic.ArticlePass(ids, pass))
}
