package model

// UserInfo 用户信息
type UserInfo struct {
	Id           int    `json:"id"                dc:"用户ID"`
	UserName     string `json:"userName"          dc:"用户名"`
	Avatar       string `json:"avatar"            dc:"头像"`
	Email        string `json:"email"             dc:"邮箱"`
	Mobile       string `json:"mobile"            dc:"手机号码"`
	Status       int    `json:"status"            dc:"状态 默认1 正常"`
	CreatedAt    string `json:"createdAt"         dc:"创建时间"`
	RegisterIp   string `json:"registerIp"        dc:"注册IP"`
	LoginIp      string `json:"loginIp"           dc:"最后登录IP"`
	LastActiveAt string `json:"lastActiveAt"      dc:"最后活跃时间"`
}
