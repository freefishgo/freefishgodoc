package controllers

import (
	"freefishgodoc/models"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type quickStartController struct {
	mvc.Controller
}

func init() {
	quickStart := &quickStartController{}
	quickStart.ActionRouterList = []*mvc.ActionRouter{
		{RouterPattern: "{Controller}/",
			ControllerActionFuncName: "Index"},
		{RouterPattern: "{Controller}", ControllerActionFuncName: "Index"}}
	mvc.AddHandlers(quickStart)
}

// 快速入门页面
func (quickStart *quickStartController) Index() {
	tp := quickStart.Query["type"]
	if tp == "xhr" {
		quickStart.UseTplPath()
		return
	}
	quickStart.LayoutPath = "layout/homeLayout.fish"
	quickStart.Data["homeHeadLi"] = models.GetHomeHeadList("快速入门")
	quickStart.UseTplPath()
}
