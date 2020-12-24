package bll

import (
	"context"
	"gxt-api-frame/app/schema"
)

type IRegister interface {
	Create(ctx context.Context, item schema.Register) (*schema.Register, error)
}
