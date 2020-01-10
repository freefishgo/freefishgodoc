package controllers

import (
	"freefishgodoc/models"
	"html/template"
	"io/ioutil"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type blogController struct {
	mvc.Controller
}

var blogContent = ""

const (
	blogIndexPath = "confStaic/blog/index.fish"
)

func init() {
	b, _ := ioutil.ReadFile(blogIndexPath)
	blogContent = string(b)
	blog := &blogController{}
	blog.ActionRouterList = []*mvc.ActionRouter{
		{RouterPattern: "{Controller}/",
			ControllerActionFuncName: "Index"},
		{RouterPattern: "{Controller}", ControllerActionFuncName: "Index"}}
	mvc.AddHandlers(blog)
}

func (blog *blogController) Index() {
	tp := blog.Query["type"]
	blog.Data["content"] = template.HTML(blogContent)
	if tp == "xhr" {
		blog.UseTplPath()
		return
	}
	blog.LayoutPath = "layout/homeLayout.fish"
	blog.Data["homeHeadLi"] = models.GetHomeHeadList("官方博客")
	blog.UseTplPath()
}
func (blog *blogController) GetEditContent() {
	blog.Response.Write([]byte(blogContent))
}

func (blog *blogController) SavePost() {
	if v, ok := blog.Query["content"]; ok {
		v := v.(string)
		blogContent = v
		ioutil.WriteFile(blogIndexPath, []byte(blogContent), 0644)
		blog.Response.WriteJson(true)
		return
	}
	blog.Response.WriteJson(false)
}
