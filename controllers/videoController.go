package controllers

import (
	"freefishgodoc/models"
	"html/template"
	"io/ioutil"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type videoController struct {
	mvc.Controller
}

var videoContent = ""

const (
	videoIndexPath = "confStaic/video/index.fish"
)

func init() {
	b, _ := ioutil.ReadFile(videoIndexPath)
	videoContent = string(b)
	video := &videoController{}
	video.ActionRouterList = []*mvc.ActionRouter{
		{RouterPattern: "{Controller}/",
			ControllerActionFuncName: "Index"},
		{RouterPattern: "{Controller}", ControllerActionFuncName: "Index"}}
	mvc.AddHandlers(video)
}

func (video *videoController) Index() {
	tp := video.Query["type"]
	video.Data["content"] = template.HTML(videoContent)
	if tp == "xhr" {
		video.UseTplPath()
		return
	}
	video.LayoutPath = "Layout/homeLayout.fish"
	video.Data["homeHeadLi"] = models.GetHomeHeadList("视频教程")
	video.UseTplPath()
}
func (video *videoController) GetEditContent() {
	video.Response.Write([]byte(videoContent))
}

func (video *videoController) SavePost() {
	if v, ok := video.Query["content"]; ok {
		v := v.(string)
		videoContent = v
		ioutil.WriteFile(videoIndexPath, []byte(videoContent), 0644)
		video.Response.WriteJson(true)
		return
	}
	video.Response.WriteJson(false)
}
