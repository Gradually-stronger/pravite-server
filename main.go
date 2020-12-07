package main

import (
	"github.com/gogf/gf/frame/g"
	_ "gxt-api-frame/boot"
	_ "gxt-api-frame/router"
)

func main() {
	g.Server().Run()
}
