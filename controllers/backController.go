package controllers

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type backController struct {
	mvc.Controller
}

func init() {
	mvc.AddHandlers(&backController{})
}

func (back *backController) Index() {
	back.UseTplPath()
}

// 控制器执行前调用
func (back *backController) Prepare() {
	sv := back.Response.GetSession("userinfo")
	if sv == nil {
		back.Response.Redirect("/user/login")
		back.SkipController()
	}
}

// 控制器结束时调用
func (back *backController) Finish() {
	//log.Println("子类的Finish")
}
