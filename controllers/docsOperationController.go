package controllers

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type docsOperationController struct {
	mvc.Controller
}

func init() {
	mvc.AddHandlers(&docsOperationController{})
}

func (docsOperation *docsOperationController) Index() {

}

// 控制器执行前调用
func (docsOperation *docsOperationController) Prepare() {
	//log.Println("子类的Prepare")
}

// 控制器结束时调用
func (docsOperation *docsOperationController) Finish() {
	//log.Println("子类的Finish")
}