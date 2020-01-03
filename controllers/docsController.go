package controllers

import (
	"freefishgodoc/models"
	"freefishgodoc/tools"
	"github.com/freefishgo/freefishgo/middlewares/mvc"
	"html/template"
)

type docsController struct {
	mvc.Controller
}

func init() {
	mvc.AddHandlers(&docsController{})
}

func (docs *docsController) Index() {
	var cocsTrees = models.GetDocsTree()
	docs.Data["docsTree"] = template.HTML(tools.EachDocsTree(cocsTrees))
	docs.Data["centent"] = "我是内容"
	tp := docs.Query["type"]
	if tp != nil && tp == "xhr" {
		docs.UseTplPath()
		return
	}
	docs.LayoutPath = "layout/homeLayout.fish"
	docs.Data["homeHeadLi"] = models.GetHomeHeadList("开发文档")
	docs.UseTplPath()
}

// 控制器执行前调用
func (docs *docsController) Prepare() {
	//log.Println("子类的Prepare")
}

// 控制器结束时调用
func (docs *docsController) Finish() {
	//log.Println("子类的Finish")
}
