package controllers

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type docsController struct {
	mvc.Controller
}

func init() {
	mvc.AddHandlers(&docsController{})
}

func (docs *docsController) Index() {
	docs.LayoutPath = "layout/homeLayout.fish"
	docs.Data["homeHeadLi"] = "homeHeadLi"
	docs.Data["docsTree"] = "docsTree"
	docs.Data["centent"] = "我是内容"
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
