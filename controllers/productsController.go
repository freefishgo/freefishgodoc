package controllers

import (
	"freefishgodoc/models"
	"html/template"
	"io/ioutil"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type productsController struct {
	mvc.Controller
}

var productsContent = ""

const (
	productsIndexPath = "confStaic/products/index.fish"
)

func init() {
	b, _ := ioutil.ReadFile(productsIndexPath)
	productsContent = string(b)
	products := &productsController{}
	products.ActionRouterList = []*mvc.ActionRouter{
		{RouterPattern: "{Controller}/",
			ControllerActionFuncName: "Index"},
		{RouterPattern: "{Controller}", ControllerActionFuncName: "Index"}}
	mvc.AddHandlers(products)
}

func (products *productsController) Index() {
	tp := products.Query["type"]
	products.Data["content"] = template.HTML(productsContent)
	if tp == "xhr" {
		products.UseTplPath()
		return
	}
	products.LayoutPath = "layout/homeLayout.fish"
	products.Data["homeHeadLi"] = models.GetHomeHeadList("产品案例")
	products.UseTplPath()
}
func (products *productsController) GetEditContent() {
	products.Response.Write([]byte(productsContent))
}

func (products *productsController) SavePost() {
	if v, ok := products.Query["content"]; ok {
		v := v.(string)
		productsContent = v
		ioutil.WriteFile(productsIndexPath, []byte(productsContent), 0644)
		products.Response.WriteJson(true)
		return
	}
	products.Response.WriteJson(false)
}
