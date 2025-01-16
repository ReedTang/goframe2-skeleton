// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// Auth 前台系统权限控制，用户必须登录才能访问
		Auth(r *ghttp.Request)
		// IsExceptLogin 是否是不需要登录的路由地址
		IsExceptLogin(ctx context.Context, path string) bool
		// ResponseHandler 返回处理中间件
		ResponseHandler(r *ghttp.Request)
		// Ctx 自定义上下文对象
		Ctx(r *ghttp.Request)
		// CORS 跨域处理中间件
		CORS(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
