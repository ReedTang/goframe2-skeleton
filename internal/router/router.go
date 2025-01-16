package router

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"goframe2-skeleton/internal/controller/auth"
	"goframe2-skeleton/internal/service"
)

func Register(group *ghttp.RouterGroup) {
	group.Middleware(
		service.Middleware().Ctx,
		service.Middleware().CORS,
		service.Middleware().ResponseHandler,
	)
	group.Bind(
		auth.NewV1(),
	)
	// 权限控制路由
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Auth)
		group.Bind()
	})
}
