package user

import (
	"context"
	"goframe2-skeleton/internal/model"
	"goframe2-skeleton/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// GetUserInfo 获取用户信息
func (s *sUser) GetUserInfo(ctx context.Context, userId int64) (userInfo *model.UserInfo, err error) {
	err = g.Model("user").Ctx(ctx).Where("id", userId).Scan(&userInfo)
	if err != nil {
		return nil, err
	}
	if userInfo == nil {
		err = gerror.New("用户不存在")
		return
	}
	return
}
