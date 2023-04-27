package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Context struct {
	Session *ghttp.Session // Session in context.
	User    *ContextUser   // User in context.
	Data    g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

type ContextUser struct {
	Id   int    `json:"id"`   // 用户id
	Code string `json:"code"` // 邀请码
}
