package cmd

import (
	"context"

	"goframe2-skeleton/internal/crontab"
	"goframe2-skeleton/internal/router"
	"goframe2-skeleton/pkg/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
)

var Http = &gcmd.Command{
	Name:  "http",
	Usage: "http",
	Brief: "HTTP服务，也可以称为主服务",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		s := g.Server()
		// 错误状态码接管
		s.BindStatusHandler(404, func(r *ghttp.Request) {
			r.Response.Writeln("404 - 你似乎来到了没有知识存在的荒原…")
		})
		s.BindStatusHandler(403, func(r *ghttp.Request) {
			r.Response.Writeln("403 - 网站拒绝显示此网页")
		})
		// 初始化路由
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.GET("/swagger-ui", func(r *ghttp.Request) {
				r.Response.Write(renderSwaggerUIPageContent())
			})
			// 注册路由
			router.Register(group)
		})
		enhanceOpenAPIDoc(s)
		// s.EnableAdmin()
		// 初始化定时任务
		crontab.Run(ctx)
		s.AddStaticPath("/resource/public", "./resource/public")
		serverWg.Add(1)
		// 信号监听
		signalListen(ctx, signalHandlerForOverall)
		go func() {
			<-serverCloseSignal
			// websocket.StopWebSocket() // 关闭websocket
			_ = s.Shutdown() // 关闭http服务，主服务建议放在最后一个关闭
			g.Log().Debug(ctx, "http successfully closed ..")
			serverWg.Done()
		}()
		// Just run the server.
		s.Run()
		return
	},
}

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = response.JsonRes{}
	openapi.Config.CommonResponseDataField = `Data`
	openapi.Components = goai.Components{
		SecuritySchemes: goai.SecuritySchemes{
			"BearerToken": goai.SecuritySchemeRef{
				Ref: "",
				Value: &goai.SecurityScheme{
					// 此处type是openApi的规定，详见 https://swagger.io/docs/specification/authentication/api-keys/
					Type:         "http",
					In:           "header",
					Scheme:       "bearer",
					BearerFormat: "JWT",
				},
			},
		},
	}

	// API description.
	openapi.Info.Title = `API For NineStars`
	openapi.Info.Description = `This documentation describes the API.`
}

func renderSwaggerUIPageContent() (swaggerUIPageContent string) {
	swaggerUIPageContent = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<meta name="description" content="SwaggerUI"/>
	<title>SwaggerUI</title>
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.11.0/swagger-ui.min.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.11.0/swagger-ui-bundle.js" crossorigin></script>
<script>
	window.onload = () => {
		window.ui = SwaggerUIBundle({
			url:    '/api.json',
			dom_id: '#swagger-ui',
			docExpansion: 'none',
			deepLinking: true,
		});
	};
</script>
</body>
</html>
`

	return
}
