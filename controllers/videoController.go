package controllers

import (
	"freefishgodoc/models"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type videoController struct {
	mvc.Controller
}

func init() {
	video := &videoController{}
	video.ActionRouterList = []*mvc.ActionRouter{
		{RouterPattern: "{Controller}/",
			ControllerActionFuncName: "Index"},
		{RouterPattern: "{Controller}", ControllerActionFuncName: "Index"}}
	mvc.AddHandlers(video)
}

func (video *videoController) Index() {
	tp := video.Query["type"]
	if tp == "xhr" {
		video.UseTplPath()
		return
	}
	video.LayoutPath = "layout/homeLayout.fish"
	video.Data["homeHeadLi"] = models.GetHomeHeadList("视频教程")
	video.UseTplPath()
}
