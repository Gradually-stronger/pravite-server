package errors

import "github.com/gogf/gf/errors/gerror"

// 定义别名
var (
	New       = gerror.New
	Wrap      = gerror.Wrap
	Wrapf     = gerror.Wrapf
	WithStack = gerror.Stack
)

var (
	ErrBadRequest              = New400Response("请求发生错误")

	ErrNoPerm                = NewResponse(401, "无访问权限", 401)
	ErrInvalidToken          = NewResponse(9999, "令牌失效", 401)
	ErrNotFound              = NewResponse(404, "资源不存在", 404)
	ErrTooManyRequests       = NewResponse(429, "请求过于频繁", 429)
	ErrInternalServer        = NewResponse(500, "服务器发生错误", 500)
	ErrDBServerInternalError = NewResponse(50001, "数据库发生错误", 500)
)
