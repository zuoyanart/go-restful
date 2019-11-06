package main

import (
	"pizzaCmsApi/controller"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func main() {
	api := iris.New()
	logMiddleware := logger.New()
	api.Use(logMiddleware)

	api.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		logMiddleware(ctx)
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(map[string]interface{}{"state": false, "msg": "404 not found"})
	})

	/**
	 * 路由分组  v1版本
	 * @method Region
	 */
	// See https://github.com/kataras/iris/wiki/API-versioning
	// if you want to take advandage of the Iris built'n versioning features.
	//
	// /useradmin
	userAdminRouter := api.Party("/useradmin")
	{
		userAdminRouter.Get("/", controller.UserAdminGet)               // /useradmin?id=1
		userAdminRouter.Get("/{id:int}", controller.UserAdminGetByPath) //useradmin/1
		userAdminRouter.Put("/", controller.UserAdminUpdate)
		userAdminRouter.Post("/", controller.UserAdminCreate)
		userAdminRouter.Post("/page", controller.UserAdminPage)
		userAdminRouter.Post("/login", controller.UserAdminCheckLogin)
		userAdminRouter.Delete("/", controller.UserAdminDele)
	}

	//article
	articleRouter := api.Party("/article")
	{
		articleRouter.Get("/{id:int}", controller.ArticleGet) //user/1
		articleRouter.Put("/", controller.ArticleUpdate)
		articleRouter.Post("/", controller.ArticleCreate)
		articleRouter.Post("/page", controller.ArticlePage)
		articleRouter.Post("/pass", controller.ArticlePass)
		articleRouter.Delete("/", controller.ArticleDele)
	}

	//node
	// api.Get("/node/pageall", controller.NodePageAll)
	// api.Get("/node/{id:int}", controller.NodeGet) //user/1
	// api.Put("/node", controller.NodeUpdate)
	// api.Post("/node", controller.NodeCreate)
	// api.Post("/node/page", controller.NodePage)

	api.Run(iris.Addr(":8081"))
}
