package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jhonnli/go-base/controller"
	cronjob "github.com/jhonnli/go-base/cron-job"
	"github.com/jhonnli/go-base/dao"
	"github.com/jhonnli/go-base/initial"
	"github.com/jhonnli/go-base/initial/config"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	config.InitConfig()
	initial.InitLog()
	dao.Init()
	controller.Init(r)
	go cronjob.DoWork()
	initial.Log.Info("go-base示例启动成功")
	r.Run(config.Config.Listen.Address)
}
