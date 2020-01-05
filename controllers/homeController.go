package controllers

import (
	"freefishgodoc/models"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

// HomeController 实现mvc控制器的处理Main为控制器 {Controller}的值
type HomeController struct {
	mvc.Controller
}

// 注册控制器
func init() {
	mvc.AddHandlers(&HomeController{})
}

// Index 为{Action}的值 该方法的默认路由为/Home/Index 最后的单词为请求方式  该例子为Post请求
func (home *HomeController) Index() {
	tp := home.Query["type"]
	if tp != nil && tp == "xhr" {
		home.UseTplPath()
		return
	}
	//log.Println(dd)
	//http.ServeTLS(l, handler, certFile, keyFile)
	home.LayoutPath = "layout/homeLayout.fish"
	home.Data["homeHeadLi"] = models.GetHomeHeadList("首页")
	//var cocsTrees=models.GetDocsTree()

	//home.Data["docsTree"] =template.HTML(tools.EachDocsTree(cocsTrees))
	//home.Data["centent"] = "我是内容"
	home.UseTplPath()
}
