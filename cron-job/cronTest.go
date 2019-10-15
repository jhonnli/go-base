package crontab

import (
	"github.com/jhonnli/logs"
	"time"
)

func CronTest() {
	logs.Info("定时任务开始执行，现在时间是：%s", time.Now().Format(time.RFC3339))
}
