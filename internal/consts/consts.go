package consts

import "github.com/gogf/gf/v2/errors/gcode"

const (
	SuccessCode int = 0
	// ErrorCode 失败code
	ErrorCode             int = -1
	ErrorCodeUnAuthorized int = 401 // 登录失效
	ErrorCodeMultiLogin   int = 402 // 账号已在其他地方登录
)

// cache
const (
	CacheToken           = "token"            // 登录token
	CacheTokenBind       = "token_bind"       // 登录用户身份绑定
	CacheMultipartUpload = "multipart_upload" // 分片上传
)

// 错误解释
const (
	ErrorORM         = "sql执行异常"
	ErrorNotData     = "数据不存在"
	ErrorRotaPointer = "指针转换异常"
)

// ErrorMessage 错误显示信息，非debug模式有效
func ErrorMessage(err error, code int) (message string) {
	if err == nil {
		return "服务器暂时无法响应，请稍后再试。"
	}

	message = err.Error()
	if code != gcode.CodeNil.Code() {
		message = "服务器居然开小差了，请稍后再试吧！"
	}

	return
}

const (
	UserSessionKey = "UserSessionKey"
	// 阿里云短信配置
	AccessKey       = ""
	AccessKeySecret = ""
	SignName        = ""
	TemplateCode    = ""
)

// RequestEncryptKey
// 请求加密密钥用于敏感数据加密，16位字符，前后端需保持一致
// 安全起见，生产环境运行时请注意修改
var RequestEncryptKey = []byte("f080a463654b2279")
