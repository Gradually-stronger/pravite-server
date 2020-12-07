package logger

import (
	"context"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

func (e *Entry) LogMode(lv logger.LogLevel) logger.Interface {
	switch lv {
	case logger.Error:
		_ = e.entry.SetLevelStr("ERROR")
		break
	case logger.Info:
		_ = e.entry.SetLevelStr("INFO")
		break
	}
	return e
}

func (e *Entry) Info(ctx context.Context, format string, args ...interface{}) {
	e.entry.Ctx(ctx).Infof(format, args...)
}

func (e *Entry) Warn(ctx context.Context, format string, args ...interface{}) {
	e.entry.Ctx(ctx).Warningf(format, args...)
}

func (e *Entry) Error(ctx context.Context, format string, args ...interface{}) {
	e.entry.Ctx(ctx).Errorf(format, args...)
}

func (e *Entry) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	limit := g.Cfg().GetInt("gorm.slow_sql_limit")
	if e.entry.GetLevel() > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && e.entry.GetLevel() >= glog.LEVEL_ERRO:
			{
				sql, _ := fc()
				e.Info(ctx, "%s\nSQL查询出错:%s\n执行SQL:%s", utils.FileWithLineNum(), err, sql)
			}
		case elapsed > time.Duration(limit)*time.Millisecond && e.entry.GetLevel() >= glog.LEVEL_WARN:
			sql, rows := fc()
			slowLog := fmt.Sprintf("执行时间 %v", elapsed)
			if rows == -1 {
				e.Warnf("%s\n慢查询SQL:%s\n%s \n影响行数:[%s]", utils.FileWithLineNum(), sql, slowLog, "-")
			} else {
				e.Warnf("%s\n慢查询SQL:%s\n%s\n影响行数:[%d]", utils.FileWithLineNum(), sql, slowLog, rows)
			}
		case e.entry.GetLevel() >= glog.LEVEL_INFO:
			sql, rows := fc()
			if rows == -1 {
				e.Infof("执行SQL:[%s],影响行数:[%s]", sql, "-")
			} else {
				e.Infof("执行SQL:[%s], 影响行数：[%d]", sql, rows)
			}
		}
	}
}
