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
	if path == "" || path == nil {
		path = "0"
	}
	path = strings.Trim(path.(string), "/")
	var cocsTrees = models.GetDocsTree()
	b, ok := tools.GetArticle(cocsTrees, path.(string))
	if ok {
		docs.Data["content"] = template.HTML(b)
	} else {
		docs.Response.WriteHeader(404)
	}
	if tp == "xhr" {
		docs.Data["docsTree"] = template.HTML(tools.EachDocsTree(cocsTrees, "0", path.(string)))
		docs.UseTplPath()
		return
	} else if tp == "xhrSon" {
		docs.Response.WriteJson(docs.Data["content"])
		return
	}
	docs.Data["docsTree"] = template.HTML(tools.EachDocsTree(cocsTrees, "0", path.(string)))
	docs.LayoutPath = "layout/homeLayout.fish"
	docs.Data["homeHeadLi"] = models.GetHomeHeadList("开发文档")
	docs.UseTplPath()
}
