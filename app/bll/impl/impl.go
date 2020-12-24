package impl

import (
	"go.uber.org/dig"
	"gxt-api-frame/app/bll"
	"gxt-api-frame/app/bll/impl/internal"
)

func Inject(container *dig.Container) {
	_ = container.Provide(internal.NewTrans)
	_ = container.Provide(func(b *internal.Trans) bll.ITrans { return b })
	_ = container.Provide(internal.NewDemo)
	_ = container.Provide(func(b *internal.Demo) bll.IDemo { return b })
	container.Provide(internal.NewRegister)
	container.Provide(func(b *internal.Register) bll.IRegister { return b })

}
