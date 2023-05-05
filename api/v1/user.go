package user

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

// LoginDoReq 执行登录请求
type LoginDoReq struct {
	g.Meta   `path:"/user/login" method:"post" summary:"登录" tags:"用户"`
	Email    string `json:"email" v:"required#请输入邮箱号"   dc:"邮箱号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}

// LoginDoRes 执行登录响应
type LoginDoRes struct {
	Token  string    `json:"token" dc:"令牌"`
	Expire time.Time `json:"expire" dc:"过期时间"`
}

// RefreshTokenReq 刷新令牌
type RefreshTokenReq struct {
	g.Meta `path:"/user/refresh_token" method:"post" summary:"刷新令牌" tags:"用户" security:"BearerToken"`
}

// RefreshTokenRes 刷新令牌
type RefreshTokenRes struct {
	Token  string    `json:"token" dc:"令牌"`
	Expire time.Time `json:"expire" dc:"过期时间"`
}

// LogoutReq 退出登录
type LogoutReq struct {
	g.Meta `path:"/user/logout" method:"post" summary:"退出登录" tags:"用户" security:"BearerToken"`
}

// LogoutRes 退出登录
type LogoutRes struct {
	Result bool `json:"result" dc:"结果"`
}

// AddInviterCode 生产验证码
type AddInviterCode struct {
	g.Meta `path:"/user/add_inviter_code" method:"post" tags:"前台用户" summary:"发送验证码"`
}

// GetUserInfoReq 获取用户信息
type GetUserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" summary:"获取用户信息" tags:"用户" security:"BearerToken"`
}

// GetUserInfoRes 获取用户信息
type GetUserInfoRes struct {
	Id        int64       `json:"id" dc:"头像"        `    // ID
	UserName  string      `json:"userName" dc:"用户名"  `   // 用户名
	Status    int         `json:"status" dc:"状态"    `    // 状态：1=正常；-1=禁用
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间" `  // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间" `  // 更新时间
	Email     string      `json:"email" dc:"邮箱"     `    // 邮箱
	Mobile    int         `json:"mobile" dc:"手机号"    `   // 手机号
	InviterId int64       `json:"inviterId" dc:"邀请者id" ` // 邀请者id
	Code      string      `json:"code" dc:"邀请码" `        // 邀请码
}
