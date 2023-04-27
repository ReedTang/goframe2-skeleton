// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AiUser is the golang structure for table ai_user.
type AiUser struct {
	Id         int64       `json:"id"         ` // ID
	UserName   string      `json:"userName"   ` // 用户名
	Password   string      `json:"password"   ` // 密码
	Status     int         `json:"status"     ` // 状态：1=正常；2=禁用
	CreatedAt  *gtime.Time `json:"createdAt"  ` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` // 更新时间
	Email      string      `json:"email"      ` // 邮箱
	Mobile     int         `json:"mobile"     ` // 手机号
	InviterId  int64       `json:"inviterId"  ` // 邀请者id
	Salt       string      `json:"salt"       ` // 盐
	Code       string      `json:"code"       ` // 邀请码
	AiDeadline *gtime.Time `json:"aiDeadline" ` // AI有效截止日期
	AiNum      int         `json:"aiNum"      ` // ai次数
}
