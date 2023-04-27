package middleware

import (
	"fmt"
	"goframe2-skeleton/internal/model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"goframe2-skeleton/internal/consts"
	"goframe2-skeleton/internal/service"
	"goframe2-skeleton/utility/response"
)

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

// Auth 前台系统权限控制，用户必须登录才能访问
func (s *sMiddleware) Auth(r *ghttp.Request) {
	service.Auth().MiddlewareFunc()(r)
	if id := gconv.Int(service.Auth().GetIdentity(r.GetCtx())); id != 0 {
		user, err := service.User().UserInfoPublic(r.GetCtx(), id)
		if err != nil {
			response.JsonExit(r, consts.ErrorNoAuthCode, err.Error())
		}
		service.BizCtx().SetUser(r.Context(), &model.ContextUser{
			Id:   gconv.Int(user.Id),
			Code: user.Code,
		})
	}

	r.Middleware.Next()
}

// ResponseHandler 返回处理中间件
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		err             = r.GetError()
		res             = r.GetHandlerResponse()
		code gcode.Code = gcode.CodeOK
	)
	if err != nil {
		// 处理全局错误
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		if gstr.Contains(err.Error(), "pq:") {

			request := g.Map{
				"method": r.Request.Method,
				"ip":     r.GetClientIp(),
				"uri":    r.Request.RequestURI,
			}
			env, _ := g.Cfg().Get(r.GetCtx(), "server.env")
			webhookRobotUrl, _ := g.Cfg().Get(r.GetCtx(), "wework.webhook_robot_url")
			bodyMap := g.Map{
				"msgtype": "markdown",
				"markdown": g.Map{
					"content": fmt.Sprintf("❌**%s %s**\n>**请求环境**: `%s`\n>**请求参数**:```%s```  \n>**报错信息**:`%s` \n", request["method"], request["uri"], env.String(), gconv.String(request), err.Error()),
				},
			}
			err = gerror.New("服务器居然开小差了，请稍后再试吧！")
			client := g.Client()
			client.SetHeader("Content-Type", "application/json; charset=UTF-8")
			if _, e := client.Post(r.GetCtx(), webhookRobotUrl.String(), gconv.String(bodyMap)); e != nil {
				g.Log().Error(r.GetCtx(), e)
			}

		}
		r.Response.ClearBuffer()
		response.JsonExit(r, consts.ErrorCode, err.Error())
	} else {
		response.JsonExit(r, code.Code(), "", res)
	}

}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	r.Middleware.Next()
}

// CORS 跨域处理中间件
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
