package middleware

import "github.com/gogf/gf/net/ghttp"

// EmptyMiddleware 不执行业务处理的中间件
func EmptyMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
}

type SkipperFunc func(request *ghttp.Request) bool
// AllowPathPrefixSkipper 检查请求路径是否包含指定的前缀，如果包含则跳过
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(request *ghttp.Request) bool {
		path := request.URL.Path
		pathLen := len(path)
		for _, p := range prefixes {
			if pl := len(p);pathLen >=pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// AllowPathPrefixNoSkipper 检查请求路径是否包含指定的前缀，如果包含则不跳过
func AllowPathPrefixNoSkipper(prefixes ...string) SkipperFunc {
	return func(request *ghttp.Request) bool {
		path := request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}

// SkipHandler 统一处理跳过函数
func SkipHandler(r *ghttp.Request, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(r) {
			return true
		}
	}
	return false
}