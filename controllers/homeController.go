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
	home := &HomeController{}
	home.ActionRouterList = []*mvc.ActionRouter{
		{RouterPattern: "{Controller}/",
			ControllerActionFuncName: "Index"},
		{RouterPattern: "{Controller}", ControllerActionFuncName: "Index"}}
	mvc.AddHandlers(home)
}

// Index 为{Action}的值 该方法的默认路由为/Home/Index 最后的单词为请求方式  该例子为Post请求
func (home *HomeController) Index() {
	tp := home.Query["type"]
	if tp != nil && tp == "xhr" {
		home.UseTplPath()
		return
	}
	home.Data["homeHeadLi"] = models.GetHomeHeadList("首页")
	home.LayoutPath = "layout/homeLayout.fish"
	home.UseTplPath()
}
