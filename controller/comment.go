package controller

import (
	"github.com/kataras/iris"
	"pizzaCmsApi/model"
)

/**
 * @api {get} /comment/:id 获取文章评论
 * @apiName get comment
 * @apiGroup comment
 * @apiVersion 1.0.0
 * @apiDescription 获取文章评论信息
 * @apiSampleRequest /comment/:id
 * @apiParam {int} id文章评论的id
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 消息
 * @apiSuccess {int} --id 主键id
 * @apiSuccess {int} --articleid 文章id
 * @apiSuccess {int} --addtime 添加时间
 * @apiSuccess {string} --content 评论内容
 * @apiSuccess {int} --uid 用户id
 * @apiSuccess {string} --username 用户昵称
 */
func CommentGet(ctx *iris.Context) {
	id := Tools.ParseInt(ctx.Param("id"), 0)
	err1 := validate.Var(id, "omitempty,min=1")
	if err1 != nil {
		ctx.JSON(iris.StatusOK, errorValidate())
	}
	ctx.JSON(iris.StatusOK, model.CommentGet(id))
}

/**
* @api {PUT} /comment 更新文章评论
* @apiName update comment
* @apiGroup comment
* @apiVersion 1.0.0
* @apiDescription 更新文章评论信息
* @apiSampleRequest /comment
* @apiParam {Comment} comment
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func CommentUpdate(ctx *iris.Context) {
	var comment model.Comment
	err := ctx.ReadJSON(&comment)
	if err != nil {
		ctx.JSON(iris.StatusOK, model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(comment)
	if err1 != nil {
		ctx.JSON(iris.StatusOK, `{"state": false, "msg": `+err1.Error()+`}`)
	}
	ctx.JSON(iris.StatusOK, model.CommentUpdate(comment))
}

/**
* @api {post} /comment 创建文章评论
* @apiName create comment
* @apiGroup comment
* @apiVersion 1.0.0
* @apiDescription 创建文章评论信息
* @apiSampleRequest /comment
* @apiParam {Comment} comment 文章评论信息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func CommentCreate(ctx *iris.Context) {
	var comment model.Comment
	err := ctx.ReadJSON(&comment)
	if err != nil {
		ctx.JSON(iris.StatusOK, model.ApiJson{State: false, Msg: err.Error()})
	}
	err1 := validate.Struct(comment)
	if err1 != nil {
		ctx.JSON(iris.StatusOK, `{"state": false, "msg": `+err1.Error()+`}`)
	}
	ctx.JSON(iris.StatusOK, model.CommentCreate(comment))
}

/**
* @api {post} /comment/page 获取文章评论列表
* @apiName page comment
* @apiGroup comment
* @apiVersion 1.0.0
* @apiDescription page comment
* @apiSampleRequest /comment/page
* @apiParam {string} kw 关键字
* @apiParam {int} cp cp
* @apiParam {int} mp mp
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func CommentPage(ctx *iris.Context) {
	cp := Tools.ParseInt(ctx.Param("cp"), 1)
	mp := Tools.ParseInt(ctx.Param("mp"), 20)
	kw := ctx.Param("kw")
	err1 := validate.Var(kw, "required, min=1,max=20")
	err2 := validate.Var(cp, "required,min=1")
	err3 := validate.Var(mp, "required,min=1,max=50")
	if err1 != nil || err2 != nil || err3 != nil {
		ctx.JSON(iris.StatusOK, errorValidate())
	}
	ctx.JSON(iris.StatusOK, model.CommentPage(kw, cp, mp))
}

/**
* @api {delete} /comment 删除文章评论
* @apiName delete comment
* @apiGroup comment
* @apiVersion 1.0.0
* @apiDescription delete comment by id
* @apiSampleRequest /comment
* @apiParam {string} id 文章评论id
* @apiParam {int} uid 用户的uid
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func CommentDele(ctx *iris.Context) {
	ids := Tools.ParseInt(ctx.Param("id"), 0)
	uid := Tools.ParseInt(ctx.Param("uid"), 0)
	err1 := validate.Var(ids, "required, min=1")
	err2 := validate.Var(uid, "required, min=1")
	if err1 != nil || err2 != nil {
		ctx.JSON(iris.StatusOK, errorValidate())
	}
	ctx.JSON(iris.StatusOK, model.CommentDel(ids, uid))

}
