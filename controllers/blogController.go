package controllers

import (
	"freefishgodoc/models"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type blogController struct {
	mvc.Controller
}

func init() {
	blog := &blogController{}
	blog.ActionRouterList = []*mvc.ActionRouter{
		{RouterPattern: "{Controller}/",
			ControllerActionFuncName: "Index"},
		{RouterPattern: "{Controller}", ControllerActionFuncName: "Index"}}
	mvc.AddHandlers(blog)
}

func (blog *blogController) Index() {
	tp := blog.Query["type"]
	if tp == "xhr" {
		blog.UseTplPath()
		return
	}
	blog.LayoutPath = "layout/homeLayout.fish"
	blog.Data["homeHeadLi"] = models.GetHomeHeadList("官方博客")
	blog.UseTplPath()
}
