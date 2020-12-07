package model

import (
	"context"
	"gorm.io/gorm"
	"gxt-api-frame/app/errors"
)

// NewTrans 创建事务管理实例
func NewTrans(db *gorm.DB) *Trans {
	return &Trans{db}
}

// Trans 事务管理
type Trans struct {
	db *gorm.DB
}

// Begin 开启事务
func (a *Trans) Begin(ctx context.Context) (interface{}, error) {
	result := a.db.Begin()
	if err := result.Error; err != nil {
		return nil, errors.ErrDBServerInternalError
	}
	return result, nil
}

// Commit 提交事务
func (a *Trans) Commit(ctx context.Context, trans interface{}) error {
	db, ok := trans.(*gorm.DB)
	if !ok {
		return errors.New("unknow trans")
	}

	result := db.Commit()
	if err := result.Error; err != nil {
		return errors.ErrDBServerInternalError
	}
	return nil
}

// Rollback 回滚事务
func (a *Trans) Rollback(ctx context.Context, trans interface{}) error {
	db, ok := trans.(*gorm.DB)
	if !ok {
		return errors.New("unknow trans")
	}

	result := db.Rollback()
	if err := result.Error; err != nil {
		return errors.ErrDBServerInternalError
	}
	return nil
}
