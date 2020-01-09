package controllers

import (
	"freefishgodoc/models"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type communityController struct {
	mvc.Controller
}

func init() {
	community := &communityController{}
	community.ActionRouterList = []*mvc.ActionRouter{
		{RouterPattern: "{Controller}/",
			ControllerActionFuncName: "Index"},
		{RouterPattern: "{Controller}", ControllerActionFuncName: "Index"}}
	mvc.AddHandlers(community)
}

func (community *communityController) Index() {
	tp := community.Query["type"]
	if tp == "xhr" {
		community.UseTplPath()
		return
	}
	community.LayoutPath = "layout/homeLayout.fish"
	community.Data["homeHeadLi"] = models.GetHomeHeadList("开发者社区")
	community.UseTplPath()
}
