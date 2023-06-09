// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AiUser is the golang structure of table ai_user for DAO operations like Where/Data.
type AiUser struct {
	g.Meta     `orm:"table:ai_user, do:true"`
	Id         interface{} // ID
	UserName   interface{} // 用户名
	Password   interface{} // 密码
	Status     interface{} // 状态：1=正常；2=禁用
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	Email      interface{} // 邮箱
	Mobile     interface{} // 手机号
	InviterId  interface{} // 邀请者id
	Salt       interface{} // 盐
	Code       interface{} // 邀请码
	AiDeadline *gtime.Time // AI有效截止日期
	AiNum      interface{} // ai次数
}
