package controllers

import (
	"freefishgodoc/models"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type productsController struct {
	mvc.Controller
}

func init() {
	products := &productsController{}
	products.ActionRouterList = []*mvc.ActionRouter{
		{RouterPattern: "{Controller}/",
			ControllerActionFuncName: "Index"},
		{RouterPattern: "{Controller}", ControllerActionFuncName: "Index"}}
	mvc.AddHandlers(products)
}

func (products *productsController) Index() {
	tp := products.Query["type"]
	if tp == "xhr" {
		products.UseTplPath()
		return
	}
	products.LayoutPath = "layout/homeLayout.fish"
	products.Data["homeHeadLi"] = models.GetHomeHeadList("产品案例")
	products.UseTplPath()
}
