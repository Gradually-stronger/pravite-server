package entity

import (
	"context"
	"gorm.io/gorm"
	"gxt-api-frame/app/schema"
)

// GetDemoDB 获取demo存储
func GetRegisterDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithModel(ctx, defDB, Register{})
}

// SchemaRegister register对象
type SchemaRegister schema.Register

// ToRegister 转换为register实体
func (a SchemaRegister) ToRegister() *Register {
	item := &Register{
		RecordId: a.RecordId,
		UserName: a.UserName,
		PassWord: a.PassWord,
	}
	return item
}

// Register register实体
type Register struct {
	Model
	RecordId string `gorm:"column:record_id";size:32;"` // 记录内码
	UserName string `gorm:"column:user_name";size:255`  // 用户名称
	PassWord string `gorm:"column:password";size:255`   // 用户密码
}

// Register register列表
type Registers []*Register
