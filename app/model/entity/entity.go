package entity

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	iContext "gxt-api-frame/app/context"
	"time"
)

// Model base model
type Model struct {
	ID        uint           `gorm:"column:id;primary_key;auto_increment;"`
	CreatedAt time.Time      `gorm:"column:created_at;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;"`
}

// TableName table name
func (Model) userTableName(name string) string {
	return fmt.Sprintf("%s%s", "p_", name)
}

func getDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := iContext.FromTrans(ctx)
	if ok {
		db, ok := trans.(*gorm.DB)
		if ok {
			if iContext.FromTransLock(ctx) {
				db = db.Set("gorm:query_option", "FOR UPDATE")
			}
			return db
		}
	}
	return defDB
}

func getDBWithModel(ctx context.Context, defDB *gorm.DB, m interface{}) *gorm.DB {
	return getDB(ctx, defDB).Model(m).WithContext(ctx)
}
