package controllers

import (
	"freefishgodoc/models"
	"html/template"
	"io/ioutil"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

// HomeController 实现mvc控制器的处理Main为控制器 {Controller}的值
type HomeController struct {
	mvc.Controller
}

var homeContent = ""

const (
	homeIndexPath = "confStaic/home/index.fish"
)

// 注册控制器
func init() {
	b, _ := ioutil.ReadFile(homeIndexPath)
	homeContent = string(b)
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
	home.Data["content"] = template.HTML(homeContent)
	if tp != nil && tp == "xhr" {
		home.UseTplPath()
		return
	}
	home.Data["homeHeadLi"] = models.GetHomeHeadList("首页")
	home.LayoutPath = "layout/homeLayout.fish"
	home.UseTplPath()
}

func (home *HomeController) GetEditContent() {
	home.Response.Write([]byte(homeContent))
}

func (home *HomeController) SavePost() {
	if v, ok := home.Query["content"]; ok {
		v := v.(string)
		homeContent = v
		ioutil.WriteFile(homeIndexPath, []byte(homeContent), 0644)
		home.Response.WriteJson(true)
		return
	}
	home.Response.WriteJson(false)
}
