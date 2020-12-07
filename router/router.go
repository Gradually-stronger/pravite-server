package router

import (
	"github.com/gogf/gf/net/ghttp"
	"go.uber.org/dig"
	"gxt-api-frame/router/api"
)
// InitRouters 初始化路由注册
func InitRouters(s *ghttp.Server, container *dig.Container) {
	// 注册api路由组
	err := api.RegisterRouters(s, container)
	if err != nil {
		panic(err)
	}
}
