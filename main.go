package main

import (
	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
	"rmApi/controller"
)

func main() {
	api := iris.New()
	api.Use(logger.New())
	errorLogger := logger.New()

	api.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.JSON(404, map[string]interface{}{"state": false, "msg": "404 not found"})
	})

	/**
	 * 路由分组  v1版本
	 * @method Region
	 */
	// region := app.Region().Prefix("/v1")
	//useradmin
	api.Get("/useradmin", controller.UserAdminGet)           // /useradmin?id=1
	api.Get("/useradmin/:id", controller.UserAdminGetByPath) //useradmin/1
	api.Put("/useradmin", controller.UserAdminUpdate)
	api.Post("/useradmin", controller.UserAdminCreate)
	api.Post("/useradmin/page", controller.UserAdminPage)
	api.Post("/useradmin/login", controller.UserAdminCheckLogin)
	api.Delete("/useradmin", controller.UserAdminDele)
	//article
	api.Get("/article/:id", controller.ArticleGet) //user/1
	api.Put("/article", controller.ArticleUpdate)
	api.Post("/article", controller.ArticleCreate)
	api.Post("/article/page", controller.ArticlePage)
	api.Post("/article/pass", controller.ArticlePass)
	api.Delete("/article", controller.ArticleDele)
	//node
	// api.Get("/node/pageall", controller.NodePageAll)
	// api.Get("/node/:id", controller.NodeGet) //user/1
	// api.Put("/node", controller.NodeUpdate)
	// api.Post("/node", controller.NodeCreate)
	// api.Post("/node/page", controller.NodePage)

	api.Listen("0.0.0.0:8081")
}
