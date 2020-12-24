package controllers

import "go.uber.org/dig"

// 注入 controllers
func Inject(container *dig.Container) {
	_ = container.Provide(NewDemo)
	container.Provide(NewRegister)
}
