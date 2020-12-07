package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"gxt-api-frame/library/gplus"
	"gxt-api-frame/library/utils"
)

func TraceIdMiddleware(skippers ...SkipperFunc) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		if len(skippers) > 0 && skippers[0](r) {
			r.Middleware.Next()
			return
		}
		if r.GetCtxVar(gplus.TraceIDKey).String() == "" {
			r.SetCtxVar(gplus.TraceIDKey, utils.NewTraceID())
		}
		r.Middleware.Next()
	}
}
