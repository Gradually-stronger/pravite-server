package model

import (
	"context"
	_ "github.com/gogf/gf/frame/g"
	"gorm.io/gorm"
	"gxt-api-frame/app/errors"
	"gxt-api-frame/app/model/entity"
	"gxt-api-frame/app/schema"
)

// NewRegister 创建Regiter储存实例
func NewRegister(db *gorm.DB) *Register {
	return &Register{db}
}

// Register register存储
type Register struct {
	db *gorm.DB
}

// 查询用户名是否重复
func (a *Register) QueryName(ctx context.Context, params schema.Register) (*schema.Register, error) {
	var item entity.Register
	db := entity.GetRegisterDB(ctx, a.db)
	if v := params.UserName; v != "" {
		db = db.Where("user_name", v)
	}
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.ErrDBServerInternalError
	} else if ok {
		return item.ToSchemaRegister(), nil
	}
	return nil, nil
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

// 注销账号
func (a *Register) Delete(ctx context.Context, recordID string) error {
	result := entity.GetRegisterDB(ctx, a.db).Where("record_id=?", recordID).Delete(entity.Register{})
	if err := result.Error; err != nil {
		return errors.ErrDBServerInternalError
	}
	return nil

}

// Get 查询指定数据
func (a *Register) Get(ctx context.Context, recordID string, opts ...schema.RegisterQueryOptions) (*schema.Register, error) {
	db := entity.GetRegisterDB(ctx, a.db).Where("record_id=?", recordID)
	var item entity.Register
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.ErrDBServerInternalError
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaRegister(), nil
}

//// 登陆验证
//func (a *Register) Login(ctx context.Context, params schema.Register) (*schema.Register, error) {
//	var item entity.Register
//	db := entity.GetRegisterDB(ctx, a.db)
//	if v := params.UserName; v != "" {
//		db = db.Where("user_name", v)
//	}
//	ok, err := FindOne(ctx, db, &item)
//
//}
