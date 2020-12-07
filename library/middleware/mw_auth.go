package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"gxt-api-frame/app/errors"
	"gxt-api-frame/library/auth"
	"gxt-api-frame/library/gplus"
)

func UserAuthMiddleware(skippers ...SkipperFunc) ghttp.HandlerFunc {
	jwt := auth.New()
	return func(r *ghttp.Request) {
		if len(skippers) >0 && skippers[0](r) {
			r.Middleware.Next()
			return
		}
		var userId string
		if t := gplus.GetToken(r); t != "" {
			id, err := jwt.ParseUserID(t)
			if err != nil {
				gplus.ResError(r, err)
			}
			userId = id
		}
		if userId != "" {
			gplus.SetUserId(r, userId)
		}
		if userId == "" {
			if g.Cfg().GetString("common.RunMode") == "debug" {
				gplus.SetUserId(r, g.Cfg().GetString("root.user_name"))
				r.Middleware.Next()
				return
			}
			gplus.ResError(r, errors.ErrNoPerm)
		}
		r.Middleware.Next()
	}
}