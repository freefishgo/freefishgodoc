package controllers

import (
	"freefishgodoc/models"
	"freefishgodoc/tools"
	"html/template"
	"strings"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type docsController struct {
	mvc.Controller
}

func init() {
	docs := &docsController{}
	docs.ActionRouterList = []*mvc.ActionRouter{
		{RouterPattern: "{Controller}/{path:allString}",
			ControllerActionFuncName: "Index"},
		{RouterPattern: "{Controller}", ControllerActionFuncName: "Index"}}
	mvc.AddHandlers(docs)
}

func (docs *docsController) Index() {
	tp := docs.Query["type"]
	path := docs.Query["path"]
	if path == nil {
		path = "0"
	}
	path = strings.Trim(path.(string), "/")
	if path == "index" {
		path = "0"
	}
	var cocsTrees = models.GetDocsTree()
	b, ok := tools.GetArticle(cocsTrees, path.(string))
	if ok {
		docs.Data["centent"] = b
	} else {
		docs.Response.WriteHeader(404)
	}
	if tp == "xhr" {
		docs.Data["docsTree"] = template.HTML(tools.EachDocsTree(cocsTrees, "0", path.(string)))
		docs.UseTplPath()
		return
	} else if tp == "xhrSon" {
		docs.Response.Write([]byte(docs.Data["centent"].(string)))
		return
	}
	docs.Data["docsTree"] = template.HTML(tools.EachDocsTree(cocsTrees, "0", path.(string)))
	docs.LayoutPath = "layout/homeLayout.fish"
	docs.Data["homeHeadLi"] = models.GetHomeHeadList("开发文档")
	docs.UseTplPath()
}

// 删除开发文档的文档树
func (docs *docsController) DeleteDoc() {
	index := docs.Query["index"]
	if index != nil {

	}
	docs.Response.WriteJson(true)
}

// 控制器执行前调用
func (docs *docsController) Prepare() {
	//log.Println("子类的Prepare")
}

// 控制器结束时调用
func (docs *docsController) Finish() {
	//log.Println("子类的Finish")
}
