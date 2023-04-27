package controller

import (
	"context"

	"goframe2-skeleton/api/v1"
	"goframe2-skeleton/internal/service"
)

// User 用户管理
var User = cUser{}

type cUser struct{}

// Info 获取用户信息
func (c *cUser) Info(ctx context.Context, req *user.GetUserInfoReq) (res *user.GetUserInfoRes, err error) {
	res, err = service.User().GetInfo(ctx, req)
	if err != nil {
		return
	}
	return
}
