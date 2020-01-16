package controllers

import (
	"freefishgodoc/models"
	"freefishgodoc/tools"
	"html/template"
	"io/ioutil"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type communityController struct {
	mvc.Controller
}

var communityContent = ""

const (
	communityIndexPath = "confStaic/community/index.fish"
)

func init() {
	b, _ := ioutil.ReadFile(communityIndexPath)
	communityContent = string(b)
	community := &communityController{}
	community.ActionRouterList = []*mvc.ActionRouter{
		{RouterPattern: "{Controller}/",
			ControllerActionFuncName: "Index"},
		{RouterPattern: "{Controller}", ControllerActionFuncName: "Index"}}
	mvc.AddHandlers(community)
}

func (community *communityController) Index() {
	tp := community.Query["type"]
	community.Data["content"] = template.HTML(communityContent)
	if tp == "xhr" {
		community.UseTplPath()
		return
	}
	community.LayoutPath = "Layout/homeLayout.fish"
	community.Data["homeHeadLi"] = models.GetHomeHeadList("开发者社区")
	community.UseTplPath()
}
func (community *communityController) GetEditContent() {
	community.Response.Write([]byte(communityContent))
}

func (community *communityController) SavePost() {
	if tools.IsLogin(community.Response, community.Request) {
		if v, ok := community.Query["content"]; ok {
			v := v.(string)
			communityContent = v
			ioutil.WriteFile(communityIndexPath, []byte(communityContent), 0644)
			community.Response.WriteJson(true)
			return
		}
		community.Response.WriteJson(false)
	} else {
		community.Response.Write([]byte("登录过期"))
	}
}
