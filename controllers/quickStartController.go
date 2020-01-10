package controllers

import (
	"freefishgodoc/models"
	"html/template"
	"io/ioutil"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type quickStartController struct {
	mvc.Controller
}

var quickStartContent = ""

const (
	quickStartIndexPath = "confStaic/quickStart/index.fish"
)

func init() {
	b, _ := ioutil.ReadFile(quickStartIndexPath)
	quickStartContent = string(b)
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
	quickStart.Data["content"] = template.HTML(quickStartContent)
	if tp == "xhr" {
		quickStart.UseTplPath()
		return
	}
	quickStart.LayoutPath = "layout/homeLayout.fish"
	quickStart.Data["homeHeadLi"] = models.GetHomeHeadList("快速入门")
	quickStart.UseTplPath()
}
func (quickStart *quickStartController) GetEditContent() {
	quickStart.Response.Write([]byte(quickStartContent))
}

func (quickStart *quickStartController) SavePost() {
	if v, ok := quickStart.Query["content"]; ok {
		v := v.(string)
		quickStartContent = v
		ioutil.WriteFile(quickStartIndexPath, []byte(quickStartContent), 0644)
		quickStart.Response.WriteJson(true)
		return
	}
	quickStart.Response.WriteJson(false)
}
