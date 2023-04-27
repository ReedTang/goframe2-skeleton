package router

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"goframe2-skeleton/internal/controller"
	"goframe2-skeleton/internal/service"
)

func InitRoutes(group *ghttp.RouterGroup) {
	group.Middleware(
		service.Middleware().Ctx,
		service.Middleware().CORS,
		service.Middleware().ResponseHandler,
	)
	group.Bind(
		controller.Auth,
	)
	// 权限控制路由
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Auth)
		group.Bind(
			controller.User, // 个人
		)
	})
}
