package controllers

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type froalaController struct {
	mvc.Controller
}

func init() {
	mvc.AddHandlers(&froalaController{})
}

func (froala *froalaController) Index() {
	froala.UseTplPath()
}

// 重写 指定动作的路由 该方法会在路由注册时调用
func (froala *froalaController) OverwriteRouter() []*mvc.ControllerActionRouter {
	return nil
}

// 控制器执行前调用
func (froala *froalaController) Prepare() {
	//log.Println("子类的Prepare")
}

// 控制器结束时调用
func (froala *froalaController) Finish() {
	//log.Println("子类的Finish")
}
