package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jhonnli/go-base/controller"
	"github.com/jhonnli/go-base/initial"
	//"github.com/jhonnli/container-orchestration-service/service/proxy"
	"github.com/jhonnli/logs"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	initial.InitConfig()
	initial.InitLog()
	//proxy.Init()
	controller.Init(r)
	logs.Info("容器编排服务启动成功")
	r.Run(initial.Config.Listen.Address)
}
