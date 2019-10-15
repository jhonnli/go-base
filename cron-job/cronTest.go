package cronjob

import (
	"fmt"
	"github.com/jhonnli/go-base/initial"
	"time"
)

func CronTest() {
	initial.Log.Info(fmt.Sprintf("定时任务开始执行，现在时间是：%s", time.Now().Format(time.RFC3339)))
}
