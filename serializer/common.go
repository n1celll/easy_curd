package serializer

import "github.com/gin-gonic/gin"

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error,omitempty"`
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// ForbiddenErr 未授权访问
	ForbiddenErr = 403
	// ServerErr 服务器错误
	ServerErr = 500
)

// CheckLogin 检查登录
func AuthErr() Response {
	return Response{
		Status: CodeCheckLogin,
		Msg:    "登录认证失败",
	}
}

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Status: errCode,
		Msg:    msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

func ParamsErr(err error) Response {
	return Err(400, "参数错误", err)
}
