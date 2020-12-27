package api

import (
	"github.com/gogf/gf/net/ghttp"
	"go.uber.org/dig"
	"gxt-api-frame/router/api/controllers"
)

// 注册路由
func RegisterRouters(s *ghttp.Server, container *dig.Container) error {
	controllers.Inject(container)
	gr := s.Group("/api")
	// 注册请求限制中间件
	return container.Invoke(func(
		cDemo *controllers.Demo,
		cRegister *controllers.Register,
	) {
		v1 := gr.Group("/v1")
		{
			gDemo := v1.Group("/demos")
			{
				gDemo.POST("/", cDemo.Create)
			}
		}
		gRegister := v1.Group("/register")
		{
			gRegister.POST("/", cRegister.Create)
			gRegister.DELETE("/", cRegister.Delete)
			gRegister.POST("/login", cRegister.Login)
		}
	})
}
