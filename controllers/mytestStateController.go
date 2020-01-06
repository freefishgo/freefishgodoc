package controllers

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

func init() {
	mvc.SetStateCodeHandlers(&mytestStateCodeController{})
}

type mytestStateCodeController struct {
	mvc.StateCodeController
}

// 500 错误处理函数
func (mytest *mytestStateCodeController) Error500() {
	mytest.StateCodeController.Error500()
}

// 403 处理函数
func (mytest *mytestStateCodeController) Forbidden403() {
	mytest.StateCodeController.Forbidden403()
}

// 404 处理函数
func (mytest *mytestStateCodeController) NotFind404() {
	mytest.StateCodeController.NotFind404()
}
