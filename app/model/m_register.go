package model

import (
	"context"
	"gxt-api-frame/app/schema"
)

// IDemo demo存储接口
type Register interface {
	// 创建数据
	Create(ctx context.Context, item schema.Register) error
}
