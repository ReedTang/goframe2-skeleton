package controller

import (
	"context"
	v1 "goframe2-skeleton/api/v1"
	"goframe2-skeleton/internal/service"
)

// Auth 登录管理
var Auth = cAuth{}

type cAuth struct{}

// Login 登录
func (c *cAuth) Login(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error) {
	res = &v1.LoginDoRes{}
	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
	if err != nil {
		return
	}
	return
}

// RefreshToken 刷新token
func (c *cAuth) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	res = &v1.RefreshTokenRes{}
	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
	return
}

// Logout 退出登录
func (c *cAuth) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	res = &v1.LogoutRes{}
	service.Auth().LogoutHandler(ctx)
	res.Result = true
	return
}
