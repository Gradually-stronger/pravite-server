package controllers

import (
	"github.com/gogf/gf/net/ghttp"
	"gxt-api-frame/app/bll"
	"gxt-api-frame/app/schema"
	"gxt-api-frame/library/gplus"
)

type Demo struct {
	cBll bll.IDemo
}

func NewDemo(cb bll.IDemo) *Demo {
	return &Demo{
		cBll: cb,
	}
}

// Create 创建数据
// @Tags API-Demo
// @Summary 创建数据
// @Param body body schema.Demo true "提交的数据"
// @Success 200 {object} schema.Demo
// @Failure 400 {object} schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/demos/ [post]
func (a *Demo) Create(r *ghttp.Request) {
	var data schema.Demo
	if err := gplus.ParseJson(r, &data); err != nil {
		gplus.ResError(r, err)
	}

	ctx := gplus.NewContext(r)
	result, err := a.cBll.Create(ctx, data)
	if err != nil {
		gplus.ResError(r, err)
	}
	gplus.ResSuccess(r, result)
}
