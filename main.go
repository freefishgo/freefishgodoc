package main

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"

	"github.com/freefishgo/freefishgo"
	"github.com/freefishgo/freefishgo/middlewares/printTimeMiddleware"
	_ "freefishgodoc/conf"
	_ "freefishgodoc/controllers"
	_ "freefishgodoc/routers"

)

func main() {
	// 通过注册中间件来打印任务处理时间服务
	freefishgo.UseMiddleware(&printTimeMiddleware.PrintTimeMiddleware{})
	// 利用中间件来实现http到https的转换
	//conf.Build.UseMiddleware(&httpToHttps.HttpToHttps{})
	// 把mvc实例注册到管道中
	freefishgo.UseMiddleware(mvc.DefaultMvcApp)
	freefishgo.DefaultConfig.Listen.HTTPPort = 8080
	err := freefishgo.Run()
	for {
		select {
		case e := <-err:
			panic(e)

		}
	}
}
