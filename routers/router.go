package routers

import (
	// 控制器注入到默认的处理器中
	_ "freefishgodoc/controllers"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

func init() {
	// 注册主路由  可多主路由格式      但 主页面 设置只有第一个有效
	mvc.AddMainRouter(&mvc.MainRouter{
		RouterPattern:  "/{ Controller}/{Action}",
		HomeController: "Home",
		IndexAction:    "Index"})
}
