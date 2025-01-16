package response

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// JsonRes 数据返回通用JSON数据结构
type JsonRes struct {
	Code    int         `json:"code" dc:"错误码(0:成功, -1:失败)"`    // 错误码(0:成功, -1:失败)
	Message string      `json:"msg"  example:"操作成功" dc:"提示信息"` // 提示信息
	Data    interface{} `json:"data" dc:"返回数据(业务接口定义具体数据结构)"`  // 返回数据(业务接口定义具体数据结构)
}

// Json 返回标准JSON数据。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}
	r.Response.WriteJson(JsonRes{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	Json(r, code, message, data...)
	r.Exit()
}

type PageRes struct {
	Current int `json:"page" dc:"当前页码"`
	Total   int `json:"total" dc:"总数"`
}
