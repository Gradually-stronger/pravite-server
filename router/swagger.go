/*
Package routers 生成swagger文档

文档规则请参考：https://github.com/swaggo/swag#declarative-comments-format

使用方式：

	go get -u github.com/swaggo/swag/cmd/swag
	swag init -g ./internal/app/routers/swagger.go -o ./docs/swagger*/
package router

// @title 高新通后台开发框架
// @version 1.0.0
// @description 高新通内部开发框架，基于gf+gorm+dig。
// @schemes http https
// @host 39.98.250.155:10076
// host 127.0.0.1:10076
// @basePath /