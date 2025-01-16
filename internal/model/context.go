package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type Context struct {
	Session *ghttp.Session // Session in context.
	User    *ContextUser   // User in context.
	Data    g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

type ContextUser struct {
	Id        int64       `json:"id"`        // 用户id
	UserName  string      `json:"userName"`  // 用户名
	Mobile    string      `json:"mobile"`    // 手机号
	App       string      `json:"app"`       // 登录应用
	LoginAt   *gtime.Time `json:"loginAt"`   // 登录时间
	CreatedAt *gtime.Time `json:"createdAt"` // 创建时间
}
