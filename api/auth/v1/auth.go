package v1

import (
	"goframe2-skeleton/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// LoginDoReq 执行登录请求
type LoginDoReq struct {
	g.Meta   `path:"/auth/login" tags:"登录注册" method:"post" summary:"登录"`
	Account  string `p:"account" v:"required#账号不能为空" dc:"账号"`
	Password string `p:"password" dc:"密码(明文)[账号密码登录使用]"`
}

// LoginDoRes 执行登录响应
type LoginDoRes struct {
	*model.LoginModel
}

// LoginLogoutReq 退出登录请求
type LoginLogoutReq struct {
	g.Meta `path:"/auth/logout" tags:"登录注册" method:"post" summary:"退出登录" security:"BearerToken"`
}

// LoginLogoutRes 退出登录响应
type LoginLogoutRes struct{}
