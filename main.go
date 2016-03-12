package main

import (
	"github.com/ivpusic/neo"
	"pizzaCmsApi/controller"
)

func main() {
	app := neo.App()

	/**
	 * 路由前中间件，相当于_before()函数
	 * @method Use
	 * @param  {[type]} func(ctx *neo.Ctx, next neo.Next [description]
	 */
	app.Use(func(ctx *neo.Ctx, next neo.Next) {
		// if authorized {
		//     next()
		// } else {
		//     ctx.Res.Status = 401
		// }
		// 添加跨域响应，响应所有的跨域方法，部署的时候请去掉这行代码
		ctx.Res.Header().Add("Access-Control-Allow-Origin", "*")
		ctx.Res.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT,DELETE, OPTIONS");
		ctx.Res.Header().Add("Access-Control-Max-Age", "1728000");
		next()
	})
		// 添加跨域响应，响应options验证（设置请求头，代码在上方）,部署的时候请去掉这行代码
	app.Options("*", func(ctx *neo.Ctx) (int, error) {
	    return 200,nil
	})
	/**
	 * 路由分组  v1版本
	 * @method Region
	 */
	region := app.Region().Prefix("/v1")
	//user
	region.Get("/user", controller.UserGet)// /user?id=1
	region.Get("/user/:id", controller.UserGetByPath) //user/1
	region.Put("/user/:id", controller.UserUpdateName)
	region.Post("/user", controller.UserCreate)
	region.Post("/user/page", controller.UserPage)
	region.Post("/user/login", controller.UserCheckLogin)
	//article
	region.Get("/article/:id", controller.ArticleGet) //user/1
	region.Put("/article/:id", controller.ArticleUpdate)
	region.Post("/article", controller.ArticleCreate)
	region.Post("/article/page", controller.ArticlePage)
	//node
	region.Get("/node/:id", controller.NodeGet) //user/1
	region.Get("/node/pageall", controller.NodePageAll)
	region.Put("/node", controller.NodeUpdate)
	region.Post("/node", controller.NodeCreate)
	region.Post("/node/page", controller.NodePage)


	app.Start()
}
