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
}
