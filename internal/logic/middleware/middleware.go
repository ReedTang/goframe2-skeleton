package middleware

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"goframe2-skeleton/internal/consts"
	"goframe2-skeleton/internal/library/token"
	"goframe2-skeleton/internal/model"
	"goframe2-skeleton/internal/service"
	"goframe2-skeleton/utility/response"
	"goframe2-skeleton/utility/simple"
	"goframe2-skeleton/utility/validate"
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
	ctx := r.GetCtx()
	path := r.URL.Path
	// 不需要验证登录的路由地址
	if s.IsExceptLogin(ctx, path) {
		r.Middleware.Next()
		return
	}
	// 根据token解析用户信息
	user, err := token.ParseLoginUser(r)
	if err != nil {
		if err.Error() == "登录身份已失效，请重新登录！" {
			response.JsonExit(r, consts.ErrorCodeUnAuthorized, err.Error())
		}
		if err.Error() == "账号已在其他地方登录，如非本人操作请及时修改登录密码！" {
			response.JsonExit(r, consts.ErrorCodeMultiLogin, err.Error())
		}
		response.JsonExit(r, consts.ErrorCode, err.Error())
		return
	}
	if id := gconv.Int64(user.Id); id != 0 {
		appUser, err := service.User().GetUserInfo(r.GetCtx(), id)
		if err != nil {
			response.JsonExit(r, consts.ErrorCode, err.Error())
			return
		}
		service.BizCtx().SetUser(ctx, &model.ContextUser{
			Id:       int64(appUser.Id),
			UserName: appUser.UserName,
			Mobile:   appUser.Mobile,
			LoginAt:  gtime.NewFromStr(appUser.LastActiveAt),
			App:      user.App,
		})
	}
	r.Middleware.Next()
}

// IsExceptLogin 是否是不需要登录的路由地址
func (s *sMiddleware) IsExceptLogin(ctx context.Context, path string) bool {
	pathList := g.Cfg().MustGet(ctx, "router.api.exceptLogin").Strings()
	for i := 0; i < len(pathList); i++ {
		if validate.InSliceExistStr(pathList[i], path) {
			return true
		}
	}
	return false
}

// ResponseHandler 返回处理中间件
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	// 错误状态码接管
	switch r.Response.Status {
	case 403:
		r.Response.Writeln("403 - 网站拒绝显示此网页")
		return
	case 404:
		r.Response.Writeln("404 - 你似乎来到了没有知识存在的荒原…")
		return
	}
	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}
	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		message            = "操作成功"
		err                = r.GetError()
		res                = r.GetHandlerResponse()
		code    gcode.Code = gcode.CodeOK
	)
	if err != nil {
		// 处理全局错误
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		// 判断是否开启 DeBug
		if simple.Debug(r.GetCtx()) {
			// 开启了直接报错对应的内容
			message = gerror.Current(err).Error()
		} else {
			// 没开启
			// 非正常报错 报服务器居然开小差了，请稍后再试吧！
			message = consts.ErrorMessage(gerror.Current(err), code.Code())
		}
		codeExit := consts.ErrorCode
		codeUnwrap := gerror.Unwrap(err)
		if gconv.Int(codeUnwrap) != 0 {
			codeExit = gconv.Int(codeUnwrap)
		}
		r.Response.ClearBuffer()
		response.JsonExit(r, codeExit, message)
	} else {
		response.JsonExit(r, code.Code(), message, res)
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
