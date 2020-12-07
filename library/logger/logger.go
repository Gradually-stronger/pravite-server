package logger

import (
	"context"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

// 定义键名
const (
	TraceIDKey      = "trace_id"
	UserIDKey       = "user_id"
	SpanTitleKey    = "span_title"
	SpanFunctionKey = "span_function"
	VersionKey      = "version"
)

var log *glog.Logger

// TraceIDFunc 定义获取跟踪ID的函数
type TraceIDFunc func() string

var (
	version     string
	traceIDFunc TraceIDFunc
)

// SetTraceIdFunc 设置获取追踪Id的生成函数
func SetTraceIdFunc(fn TraceIDFunc) {
	traceIDFunc = fn
}

// SetVersion 设置版本号
func SetVersion(ver string) {
	version = ver
}

// FromTraceIDContext 从上下文中获取跟踪ID
func FromTraceIDContext(ctx context.Context) string {
	v := ctx.Value(TraceIDKey)
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return getTraceID()
}

// NewUserIDContext 创建用户ID上下文
func NewUserIDContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// FromUserIDContext 从上下文中获取用户ID
func FromUserIDContext(ctx context.Context) string {
	v := ctx.Value(UserIDKey)
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func getTraceID() string {
	if traceIDFunc != nil {
		return traceIDFunc()
	}
	return ""
}

// GetLogger 获取logger
func GetLogger() *glog.Logger {
	return log
}

// NewTraceIDContext 创建跟踪ID上下文
func NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

// Init 初始化日志配置
func Init(runMode string) {
	if log != nil {
		panic(errors.New("重复初始化logger"))
	}
	log = g.Log()
	var lv string
	switch runMode {
	case "debug":
		lv = "DEV"
		break
	case "test":
		lv = "DEV"
		break
	case "release":
		lv = "PRODUCT"
		break
	}
	err := log.SetLevelStr(lv)
	if err != nil {
		panic(err)
	}
}

type spanOptions struct {
	Title    string
	FuncName string
}

// SpanOption 定义跟踪单元的数据项
type SpanOption func(*spanOptions)

// SetSpanTitle 设置跟踪单元的标题
func SetSpanTitle(title string) SpanOption {
	return func(o *spanOptions) {
		o.Title = title
	}
}

// SetSpanFuncName 设置跟踪单元的函数名
func SetSpanFuncName(funcName string) SpanOption {
	return func(o *spanOptions) {
		o.FuncName = funcName
	}
}

// StartSpan 开始一个追踪单元
func StartSpan(ctx context.Context, opts ...SpanOption) *Entry {
	if ctx == nil {
		ctx = context.Background()
	}

	var o spanOptions
	for _, opt := range opts {
		opt(&o)
	}
	return NewEntry(log)
}

// Entry 定义统一的日志写入方式
type Entry struct {
	entry *glog.Logger
}

func NewEntry(entry *glog.Logger) *Entry {
	return &Entry{entry: entry}
}

// Debugf 写入调试日志
func Debugf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).entry.Ctx(ctx).Debugf(format, args)
}

// Printf 写入消息日志
func Printf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).entry.Ctx(ctx).Printf(format, args...)
}

// Warnf 写入警告日志
func Warnf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).entry.Ctx(ctx).Warningf(format, args...)
}

// Fatalf 写入重大错误日志
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).entry.Ctx(ctx).Fatalf(format, args...)
}

// Errorf 错误日志
func Errorf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).entry.Ctx(ctx).Errorf(format, args...)
}

// Errorf 错误日志
func (e *Entry) Errorf(format string, args ...interface{}) {
	e.entry.Errorf(format, args...)
}

// Warnf 警告日志
func (e *Entry) Warnf(format string, args ...interface{}) {
	e.entry.Warningf(format, args...)
}

// Infof 消息日志
func (e *Entry) Infof(format string, args ...interface{}) {
	e.entry.Infof(format, args...)
}

// Printf 消息日志
func (e *Entry) Printf(format string, args ...interface{}) {
	e.entry.Printf(format, args...)
}

// Debugf 写入调试日志
func (e *Entry) Debugf(format string, args ...interface{}) {
	e.entry.Debugf(format, args...)
}
