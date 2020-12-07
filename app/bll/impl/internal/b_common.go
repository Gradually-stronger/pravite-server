package internal

import (
	"context"
	"gxt-api-frame/app/model"
	iContext "gxt-api-frame/app/context"
)

// TransFunc 定义事务执行函数
type TransFunc func(context.Context) error

// ExecTrans 执行事务
func ExecTrans(ctx context.Context, transModel model.ITrans, fn TransFunc) error {
	if _, ok := iContext.FromTrans(ctx); ok {
		return fn(ctx)
	}
	trans, err := transModel.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = transModel.Rollback(ctx, trans)
			panic(r)
		}
	}()

	err = fn(iContext.NewTrans(ctx, trans))
	if err != nil {
		_ = transModel.Rollback(ctx, trans)
		return err
	}
	return transModel.Commit(ctx, trans)
}

// ExecTransWithLock 执行事务（加锁）
func ExecTransWithLock(ctx context.Context, transModel model.ITrans, fn TransFunc) error {
	if !iContext.FromTransLock(ctx) {
		ctx = iContext.NewTransLock(ctx)
	}
	return ExecTrans(ctx, transModel, fn)
}