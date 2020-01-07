package controllers

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

func init() {
	mvc.SetStatusCodeHandlers(&ourStatusCodeController{})
}

type ourStatusCodeController struct {
	mvc.StatusCodeController
}

// 500 错误处理函数
func (our *ourStatusCodeController) Error500() {
	our.StatusCodeController.Error500()
}

// 403 处理函数
func (our *ourStatusCodeController) Forbidden403() {
	our.StatusCodeController.Forbidden403()
}

// 404 处理函数
func (our *ourStatusCodeController) NotFind404() {
	our.StatusCodeController.NotFind404()
}
