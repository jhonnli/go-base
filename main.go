package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jhonnli/go-base/controller"
	"github.com/jhonnli/go-base/cron-job"
	"github.com/jhonnli/go-base/dao"
	"github.com/jhonnli/go-base/initial"
	"github.com/jhonnli/go-base/initial/config"

	//"github.com/jhonnli/container-orchestration-service/service/proxy"
	"github.com/jhonnli/logs"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	config.InitConfig()
	initial.InitLog()
	dao.Init()
	//proxy.Init()
	controller.Init(r)
	logs.Info("容器编排服务启动成功")
	r.Run(config.Config.Listen.Address)
	//go cron-job.DoWork()
}
