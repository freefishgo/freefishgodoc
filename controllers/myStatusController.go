package controllers

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

func init() {
	mvc.SetStatusCodeHandlers(&myStatusCodeController{})
}

type myStatusCodeController struct {
	mvc.StatusCodeController
}

// 500 错误处理函数
func (my *myStatusCodeController) Error500() {
	my.StatusCodeController.Error500()
}

// 403 处理函数
func (my *myStatusCodeController) Forbidden403() {
	my.StatusCodeController.Forbidden403()
}

// 404 处理函数
func (my *myStatusCodeController) NotFind404() {
	my.StatusCodeController.NotFind404()
}
