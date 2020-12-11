package model

import (
	"context"
	"gorm.io/gorm"
	"gxt-api-frame/app/errors"
	"gxt-api-frame/app/model/entity"
	"gxt-api-frame/app/schema"
)

// NewRegister 创建Register储存实例
func NewRegister(db *gorm.DB) *Register {
	return &Register{db}
}

// Register register存储
type Register struct {
	db *gorm.DB
}

//Create 创建数据
func (a *Register) Create(ctx context.Context, item schema.Register) error {
	register := entity.SchemaRegister(item).ToRegister()
	result := entity.GetRegisterDB(ctx, a.db).Create(register)
	if err := result.Error; err != nil {
		return errors.ErrDBServerInternalError
	}
	return nil
}
