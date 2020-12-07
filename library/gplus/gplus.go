package gplus

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	iContext "gxt-api-frame/app/context"
	"gxt-api-frame/app/errors"
	"gxt-api-frame/app/schema"
	"gxt-api-frame/library/logger"
	"gxt-api-frame/library/utils"
	"net/http"
	"strings"
)

// 定义上下文中的键
const (
	prefix = "gao-xin-tong"
	// UserIDKey 存储上下文中的键(用户ID)
	UserIDKey = prefix + "/user-id"
	// UserTypeKey 存储上下文中的键(用户类型)
	UserTypeKey = prefix + "/user-type"
	// TraceIDKey 存储上下文中的键(跟踪ID)
	TraceIDKey = prefix + "/trace-id"
	// ResBodyKey 存储上下文中的键(响应Body数据)
	ResBodyKey = prefix + "/res-body"
)

// ParseJson 解析请求参数Json
func ParseJson(r *ghttp.Request, out interface{}) error {
	if err := r.Parse(out); err != nil {
		m := "解析请求参数发生错误"
		if g.Cfg().GetString("common.run_mode") == "debug" {
			m += "[" + err.Error() + "]"
		}
		return errors.Wrap400Response(err, m)
	}
	return nil
}

// GetPageIndex 获取当前页
func GetPageIndex(r *ghttp.Request) int {
	defaultVal := 1
	if v := r.GetQueryInt("current"); v > 0 {
		return v
	}
	return defaultVal
}

// GetPageSize 获取分页的页大小(最大50)
func GetPageSize(r *ghttp.Request) int {
	defaultVal := 10
	if v := r.GetQueryInt("pageSize"); v > 0 {
		if v > 50 {
			v = 50
		}
		return v
	}
	return defaultVal
}

// ResPage 分页响应
func ResPage(r *ghttp.Request, v interface{}, pr *schema.PaginationResult) {
	result := schema.HTTPList{
		List: v,
		Pagination: &schema.HTTPPagination{
			Current:  GetPageIndex(r),
			PageSize: GetPageSize(r),
		},
	}
	if pr != nil {
		result.Pagination.Total = pr.Total
	}
	ResSuccess(r, result)
}

// ResSuccess 响应成功
func ResSuccess(c *ghttp.Request, v interface{}) {
	ResJSON(c, http.StatusOK, v)
}

// ResOK 响应OK
func ResOK(c *ghttp.Request) {
	ResSuccess(c, schema.HTTPStatus{Status: "OK"})
}

// ResList 响应列表数据
func ResList(c *ghttp.Request, v interface{}) {
	ResSuccess(c, schema.HTTPList{List: v})
}

// ResJSON 响应JSON结果
func ResJSON(r *ghttp.Request, code int, v interface{}) {
	errors.JsonExit(r, code, v)
}

// ResError 响应错误
func ResError(r *ghttp.Request, err error, status ...int) {
	var res *errors.ResponseError
	if err != nil {
		if e, ok := err.(*errors.ResponseError); ok {
			res = e
		} else {
			res = errors.UnWrapResponse(errors.Wrap500Response(err))
		}
	} else {

		res = errors.UnWrapResponse(errors.ErrInternalServer)
	}

	if len(status) > 0 {
		res.StatusCode = status[0]
	}

	if err := res.ERR; err != nil {
		if res.StatusCode >= 500 {
			logger.Errorf(NewContext(r), "%s", err)
		}
	}

	eitem := schema.HTTPErrorItem{
		Code:    res.Code,
		Message: res.Message,
		TraceId: GetTraceID(r),
	}
	ResJSON(r, res.StatusCode, schema.HTTPError{Error: eitem})
}

// SetUserId 上下文中设置用户Id
func SetUserId(r *ghttp.Request, userId string) {
	r.SetCtxVar(UserIDKey, userId)
}

// GetUserId 获取上下文中的用户Id
func GetUserId(r *ghttp.Request) string {
	return r.GetCtxVar(UserIDKey).String()
}

// GetToken 获取token
func GetToken(r *ghttp.Request) string {
	var token string
	auth := r.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// NewContext 封装上下文入口
func NewContext(r *ghttp.Request) context.Context {
	parent := context.Background()

	traceId := GetTraceID(r)
	if traceId == "" {
		traceId = utils.NewTraceID()
	}
	parent = iContext.NewTraceID(parent, traceId)
	parent = logger.NewTraceIDContext(parent, traceId)

	if v := GetUserID(r); v != "" {
		parent = iContext.NewUserID(parent, v)
		parent = logger.NewUserIDContext(parent, v)
	}

	return parent
}

// GetTraceID 获取追踪ID
func GetTraceID(c *ghttp.Request) string {
	return c.GetCtxVar(TraceIDKey).String()
}

// GetUserID 获取用户ID
func GetUserID(c *ghttp.Request) string {
	return c.GetCtxVar(UserIDKey).String()
}

// GetUserType 获取用户类型
func GetUserType(c *ghttp.Request) string {
	return c.GetCtxVar(UserTypeKey).String()
}
