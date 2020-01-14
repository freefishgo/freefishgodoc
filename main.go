package main

import (
	"encoding/json"
	_ "freefishgodoc/routers"
	"os"

	"github.com/freefishgo/freefishgo"
	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type config struct {
	*freefishgo.Config
	WebConfig *mvc.MvcWebConfig
}

func init() {
	conf := new(config)
	f, err := os.Open("conf/app.conf")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	json.NewDecoder(f).Decode(conf)
	freefishgo.SetDefaultApplicationBuilderConfig(conf.Config)
	mvc.SetDefaultMvcWebConfig(conf.WebConfig)

}
func main() {
	// 通过注册中间件来打印任务处理时间服务
	//freefishgo.UseMiddleware(&middlewares.PrintTimeMiddleware{})
	// 利用中间件来实现http到https的转换
	//freefishgo.UseMiddleware(&httpToHttps.HttpToHttps{})
	// 把mvc实例注册到管道中
	freefishgo.UseMiddleware(mvc.GetDefaultMvcApp())
	//freefishgo.DefaultConfig.Listen.HTTPPort = 8080
	freefishgo.Run()
}
