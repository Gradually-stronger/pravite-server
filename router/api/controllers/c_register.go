package controllers

import (
	"github.com/gogf/gf/net/ghttp"
	"gxt-api-frame/app/bll"
	"gxt-api-frame/app/errors"
	"gxt-api-frame/app/schema"
	"gxt-api-frame/library/gplus"
)

type Register struct {
	cBll bll.IRegister
}

func NewRegister(cb bll.IRegister) *Register {
	return &Register{
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
func (a *Register) Create(r *ghttp.Request) {
	var data schema.Register
	if err := gplus.ParseJson(r, &data); err != nil {
		gplus.ResError(r, err)
	}
	ctx := gplus.NewContext(r)
	_, err := a.cBll.Create(ctx, data)
	if err != nil {
		gplus.ResError(r, err)
	}

	gplus.ResOK(r)
}

// Delete 注销账号

func (a *Register) Delete(r *ghttp.Request) {
	ctx := gplus.NewContext(r)
	err := a.cBll.Delete(ctx, r.GetQueryString("id"))
	if err != nil {
		gplus.ResError(r, err)
	}
	gplus.ResOK(r)
}

// Login 验证登陆账号
func (a *Register) Login(r *ghttp.Request) {
	var data schema.Register
	if err := gplus.ParseJson(r, &data); err != nil {
		gplus.ResError(r, err)
	}
	ctx := gplus.NewContext(r)
	result, err := a.cBll.Login(ctx, data)
	if err != nil {
		gplus.ResError(r, err)
	}
	if result != nil {
		gplus.ResOK(r)
	} else {
		gplus.ResError(r, errors.New400Response("用户名或密码错误"))
	}

}
