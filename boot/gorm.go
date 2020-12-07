package boot

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"go.uber.org/dig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gxt-api-frame/app/model"
	"gxt-api-frame/app/model/entity"
	iModel "gxt-api-frame/app/model/impl/model"
	"gxt-api-frame/library/logger"
	"time"
)

// 初始化gorm
func initGorm() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		g.Cfg().GetString("mysql.user"),
		g.Cfg().GetString("mysql.password"),
		g.Cfg().GetString("mysql.host"),
		g.Cfg().GetInt("mysql.port"),
		g.Cfg().GetString("mysql.db_name"),
		g.Cfg().GetString("mysql.parameters"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   g.Cfg().GetString("gorm.table_prefix"),
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		return nil, err
	}
	db.Logger = logger.NewEntry(logger.GetLogger())
	if g.Cfg().GetString("common.run_mode") == "debug" {
		db.Debug()
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(g.Cfg().GetInt("gorm.max_idle_conns"))
	sqlDb.SetMaxOpenConns(g.Cfg().GetInt("gorm.max_open_conns"))
	sqlDb.SetConnMaxLifetime(time.Duration(g.Cfg().GetInt("gorm.max_open_conns")) * time.Second)
	return db, nil
}

// 自动创建数据表映射
func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		new(entity.Demo),
	)
}

// InjectModel
func InjectModel(container *dig.Container) error {
	_ = container.Provide(iModel.NewDemo)
	_ = container.Provide(func(m *iModel.Demo) model.IDemo { return m })
	return nil
}
