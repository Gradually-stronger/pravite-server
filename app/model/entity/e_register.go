package entity

import (
	"context"
	"gorm.io/gorm"
	"gxt-api-frame/app/schema"
)

// GetRegisterDB 获取Register存储
func GetRegisterDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithModel(ctx, defDB, Register{})
}

// TableName 表名
func (a Register) TableName() string {
	return a.Model.userTableName("user")
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

// ToRegister 转换为register实体
func (a Register) ToSchemaRegister() *schema.Register {
	item := &schema.Register{
		RecordId: a.RecordId,
		UserName: a.UserName,
		PassWord: a.PassWord,
	}
	return item
}

// Register register实体
type Register struct {
	UserName string `gorm:"column:user_name";size:255`  // 用户名称
	PassWord string `gorm:"column:password";size:255`   // 用户密码
	RecordId string `gorm:"column:record_id";size:32;"` // 记录内码
	Model
}

// Register register列表
type Registers []*Register
