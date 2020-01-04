package controllers

import (
	"freefishgodoc/models"
	"freefishgodoc/tools"
	"github.com/freefishgo/freefishgo/middlewares/mvc"
	"html/template"
	"strings"
)

type docsController struct {
	mvc.Controller
}

func init() {
	docs:=&docsController{}
	docs.ActionRouterList=[]*mvc.ActionRouter{
		{RouterPattern:"{Controller}/{path:allString}",
		ControllerActionFuncName:"Index"},
		{RouterPattern:"{Controller}",ControllerActionFuncName:"Index"}}
	mvc.AddHandlers(docs)
}

func (docs *docsController) Index() {
	var cocsTrees = models.GetDocsTree()
	docs.Data["docsTree"] = template.HTML(tools.EachDocsTree(cocsTrees,"0"))
	docs.Data["centent"] = "我是内容"
	tp := docs.Query["type"]
	path:= docs.Query["path"]
	if path!=nil{
		pathStr := strings.Trim(path.(string),"/")
		strings.Split(pathStr,"/")
	}
	if tp == "xhr" {
		docs.UseTplPath()
		return
	} else if tp== "xhrSon" {
		docs.Response.Write([]byte(docs.Data["centent"].(string)))
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
