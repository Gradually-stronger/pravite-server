package model

import (
	"context"
	"gxt-api-frame/app/schema"
)

// IDemo demo存储接口
type Register interface {
	// 创建数据
	Create(ctx context.Context, item schema.Register) error
	// 查询单条
	QueryName(ctx context.Context, params schema.Register) (*schema.Register, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.RegisterQueryOptions) (*schema.Register, error)
	// 删除数据
	Delete(ctx context.Context, recordID string) error
}
