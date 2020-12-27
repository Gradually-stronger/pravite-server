package internal

import (
	"context"
	"github.com/gogf/gf/util/guid"
	"gxt-api-frame/app/errors"
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
	ok, err := a.RegisterModal.QueryName(ctx, item)
	if err != nil {
		return nil, err
	} else if ok != nil {
		if ok.UserName != item.UserName {
			item.RecordId = guid.S()
			err := a.RegisterModal.Create(ctx, item)
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, errors.New400Response("已有账号名称。")
}

// 注销账号
func (a *Register) Delete(ctx context.Context, recordID string) error {
	oldItem, err := a.RegisterModal.Get(ctx, recordID)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	return a.RegisterModal.Delete(ctx, recordID)
}

// 验证登陆
func (a *Register) Login(ctx context.Context, item schema.Register) (*schema.Register, error) {
	result, err := a.RegisterModal.QueryName(ctx, item)
	if err != nil {
		return nil, err
	}
	if result != nil {
		if result.UserName == item.UserName {
			if result.PassWord == item.PassWord {
				return result, err
			}
		} else {
			return nil, err
		}
	}
	return nil, err
}
