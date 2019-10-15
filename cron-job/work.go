package crontab

import (
	"github.com/jhonnli/logs"
	"github.com/robfig/cron"
	"time"
)

func DoWork() {
	for {
		start()
		time.Sleep(3 * time.Second)
	}
}

func start() {
	defer recoverPanic()
	worker()
}

func recoverPanic() {
	if err := recover(); err != nil {
		logs.Error("【recoverPanic】global-ex：%v ", err)
	}
}

func worker() {
	c := cron.New()
	c.AddFunc("*/5 * * * * *", CronTest)
	c.Start()
	select {}
}
