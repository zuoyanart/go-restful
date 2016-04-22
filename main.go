package main

import (
	"fmt"
	"github.com/ivpusic/neo"
	"pizzaCmsApi/controller"
	"time"
)

func main() {
	app := neo.App()

	/**
	 * 路由前中间件，相当于_before()函数
	 * @method Use
	 * @param  {[type]} func(ctx *neo.Ctx, next neo.Next [description]
	 */
	app.Use(func(ctx *neo.Ctx, next neo.Next) {
		start := time.Now()
		// fmt.Printf("\n\n---------------> [Req] %s to %s", ctx.Req.Method, ctx.Req.URL.Path)

		// if authorized {
		//     next()
		// } else {
		//     ctx.Res.Status = 401
		// }
		// 添加跨域响应，响应所有的跨域方法，部署的时候请去掉这行代码
		ctx.Res.Header().Add("Access-Control-Allow-Origin", "*")
		ctx.Res.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT,DELETE, OPTIONS")
		ctx.Res.Header().Add("Access-Control-Max-Age", "1728000")
		next()

		elapsed := float64(time.Now().Sub(start) / time.Millisecond)
    fmt.Printf("<----------------------[Res] (%d) %s to %s Took %vms\n\n", ctx.Res.Status, ctx.Req.Method, ctx.Req.URL.Path, elapsed)
	})
	// 添加跨域响应，响应options验证（设置请求头，代码在上方）,部署的时候请去掉这行代码
	app.Options("*", func(ctx *neo.Ctx) (int, error) {
		return 200, nil
	})
	/**
	 * 路由分组  v1版本
	 * @method Region
	 */
	region := app.Region().Prefix("/v1")
	//user
	region.Get("/user/:id", controller.UserGetByPath) //user/1
	region.Put("/user", controller.UserUpdate)
	region.Post("/user", controller.UserCreate)
	region.Post("/user/page", controller.UserPage)
	region.Post("/user/login", controller.UserCheckLogin)
	region.Delete("/user", controller.UserDele)
	//article
	region.Get("/article/:id", controller.ArticleGet)
	region.Put("/article", controller.ArticleUpdate)
	region.Post("/article", controller.ArticleCreate)
	region.Post("/article/page", controller.ArticlePage)
	region.Post("/article/pass", controller.ArticlePass)
	region.Delete("/article", controller.ArticleDele)
	//node
	region.Get("/node/pageall", controller.NodePageAll)
	region.Get("/node/:id", controller.NodeGet)
	region.Put("/node", controller.NodeUpdate)
	region.Post("/node", controller.NodeCreate)
	region.Post("/node/page", controller.NodePage)
	//block
	region.Get("/block/:id", controller.BlockGet)
	region.Put("/block", controller.BlockUpdate)
	region.Post("/block", controller.BlockCreate)
	region.Post("/block/page", controller.BlockPage)
	region.Delete("/block", controller.BlockDele)
	//comment
	region.Get("/comment/:id", controller.CommentGet)
	region.Put("/comment", controller.CommentUpdate)
	region.Post("/comment", controller.CommentCreate)
	region.Post("/comment/page", controller.CommentPage)
	region.Delete("/comment", controller.CommentDele)
	//用户组
	region.Get("/usergroup/:id", controller.UsergroupGet)
	region.Put("/usergroup", controller.UsergroupUpdate)
	region.Post("/usergroup", controller.UsergroupCreate)
	region.Post("/usergroup/page", controller.UsergroupPage)
	region.Delete("/usergroup", controller.UsergroupDele)
	//角色
	region.Get("/role/:id", controller.RoleGet)
	region.Put("/role", controller.RoleUpdate)
	region.Post("/role", controller.RoleCreate)
	region.Post("/role/page", controller.RolePage)
	region.Delete("/role", controller.RoleDele)
	//权限管理，部门和菜单
	region.Get("/power/:id", controller.ActionmenuGet)
	region.Put("/power", controller.ActionmenuUpdate)
	region.Post("/power", controller.ActionmenuCreate)
	region.Post("/power/page", controller.ActionmenuPage)
	region.Delete("/power", controller.ActionmenuDele)

	app.Start()
}
