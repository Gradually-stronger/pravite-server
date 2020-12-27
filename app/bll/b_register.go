package bll

import (
	"context"
	"gxt-api-frame/app/schema"
)

type IRegister interface {
	// 创建
	Create(ctx context.Context, item schema.Register) (*schema.Register, error)
	// 注销账号
	Delete(ctx context.Context, recordID string) error
	// 登陆验证账号
	Login(ctx context.Context, item schema.Register) (*schema.Register, error)
}
