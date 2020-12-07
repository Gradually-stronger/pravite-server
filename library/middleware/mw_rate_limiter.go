package middleware

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"gxt-api-frame/app/errors"
	"gxt-api-frame/library/gplus"
	"gxt-api-frame/library/logger"
	"strconv"
)

// RateLimiterMiddleware 请求频率限制中间件
func RateLimiterMiddleware(skippers ...SkipperFunc) ghttp.HandlerFunc {
	if !g.Cfg().GetBool("rate_limiter.enable") {
		return EmptyMiddleware
	}
	// check enable redis
	if !g.Cfg().GetBool("redis.enable") {
		return func(r *ghttp.Request) {
			logger.Warnf(gplus.NewContext(r), "限流中间件无法正常使用,请启用redis配置[redis.enable]")
			r.Middleware.Next()
		}
	}

	addr := g.Cfg().GetString("redis.addr")
	password := g.Cfg().GetString("redis.password")
	db := g.Cfg().GetInt("redis.db")
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": addr,
		},
		Password: password,
		DB:       db,
	})

	limiter := redis_rate.NewLimiter(ring)

	return func(r *ghttp.Request) {
		if SkipHandler(r, skippers...) {
			r.Middleware.Next()
			return
		}

		userID := gplus.GetUserID(r)
		if userID == "" {
			r.Middleware.Next()
			return
		}
		ctx := gplus.NewContext(r)
		limit := g.Cfg().GetInt("rate_limiter.count")
		result, err := limiter.Allow(ctx,
			userID, redis_rate.PerMinute(limit))
		if err != nil {
			gplus.ResError(r, errors.ErrInternalServer)
		}
		if result != nil {
			if result.Allowed == 0 {
				h := r.Response.Header()
				h.Set("X-RateLimit-Limit", strconv.FormatInt(int64(result.Limit.Burst), 10))
				h.Set("X-RateLimit-Remaining", strconv.FormatInt(int64(result.Remaining), 10))
				h.Set("X-RateLimit-Reset", strconv.FormatInt(int64(result.ResetAfter.Seconds()), 10))
				gplus.ResError(r, errors.ErrTooManyRequests)
				return
			}
		}

		r.Middleware.Next()
	}
}
