package internal

import (
	"context"
	"github.com/gogf/gf/util/guid"
	"gxt-api-frame/app/model"
	"gxt-api-frame/app/schema"
)

// register实例
type Register struct {
	RegisterModal model.Register
}

// 创建register
func NewRegister(mRegister model.Register) *Register {
	return &Register{
		RegisterModal: mRegister,
	}
}

// POST创建账号
func (a *Register) Create(ctx context.Context, item schema.Register) (*schema.Register, error) {
	item.RecordId = guid.S()
	err := a.RegisterModal.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	return nil, err
}
