package model

// LoginModel 统一登录响应
type LoginModel struct {
	Id       int64  `json:"id"              dc:"用户ID"`
	Username string `json:"username"        dc:"用户名"`
	Token    string `json:"token"           dc:"登录token"`
	Expires  int64  `json:"expires"         dc:"登录有效期"`
}

// TokenConfig 登录令牌配置
type TokenConfig struct {
	SecretKey       string `json:"secretKey"`
	Expires         int64  `json:"expires"`
	AutoRefresh     bool   `json:"autoRefresh"`
	RefreshInterval int64  `json:"refreshInterval"`
	MaxRefreshTimes int64  `json:"maxRefreshTimes"`
	MultiLogin      bool   `json:"multiLogin"`
}
