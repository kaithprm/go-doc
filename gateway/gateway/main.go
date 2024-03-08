package main

import (
	"gateway/internal"
	"gateway/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// api gateway需要一个http server来接收 client请求，并根据请求的路径分发到不同的服务
func main() {
	internal.Init()
	service.InitNacosClient()
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	r.POST("/test", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World!")
		name := internal.Query("/test")
		println("name is ---//==>>", name)
		// 服务发现 api gateway查找根据nacos_name寻找ip
		ip := service.Query(name)
		context.Redirect(http.StatusMovedPermanently, ip)
		println("ip is ---", ip)
	})
	//3.监听端口，默认8080
	r.Run(":8080")
}
