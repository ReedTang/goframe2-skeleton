package cmd

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"

	"goframe2-skeleton/internal/router"
	"goframe2-skeleton/utility/response"
)

func ServerBoot() {
	s := g.Server()
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		return
	}
	g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/swagger-ui", func(r *ghttp.Request) {
			r.Response.Write(renderSwaggerUIPageContent())
		})
		router.InitRoutes(group)
	})
	// 自定义丰富文档
	enhanceOpenAPIDoc(s)
	s.EnableAdmin()
	s.Run()
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
	var servers goai.Servers
	servers = append(servers, goai.Server{
		URL:         "http://127.0.0.1:8000/",
		Description: "本地开发环境",
	})
	servers = append(servers, goai.Server{
		URL:         "https://api.****.com/",
		Description: "正式环境",
	})
	openapi.Servers = &servers

	// API description.
	openapi.Info.Title = `goFrame2-skeleton`
	openapi.Info.Description = `goFrame2-skeleton接口`

}

func renderSwaggerUIPageContent() (swaggerUIPageContent string) {
	swaggerUIPageContent = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta
      name="description"
      content="SwaggerUI"
    />
    <title>SwaggerUI</title>
    <link rel="stylesheet" href="https://cdn.staticfile.org/swagger-ui/4.15.5/swagger-ui.min.css" />
  </head>
  <body>
  <div id="swagger-ui"></div>
  <script src="https://cdn.staticfile.org/swagger-ui/4.15.5/swagger-ui-bundle.min.js" crossorigin></script>
  <script src="https://cdn.staticfile.org/swagger-ui/4.15.5/swagger-ui-standalone-preset.min.js" crossorigin></script>
  <script>
    window.onload = () => {
      window.ui = SwaggerUIBundle({
        url: '/api.json',
        dom_id: '#swagger-ui',
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        layout: "StandaloneLayout",
      });
    };
  </script>
  </body>
</html>
`

	return
}
