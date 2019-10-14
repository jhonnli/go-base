package utils

import (
	"errors"
	"github.com/jhonnli/golibs"
	"github.com/jhonnli/logs"
	"time"
)

func TimeFormatISO8601(timeStr string) time.Time {
	if timeStr != "" || len(timeStr) > 0 {
		newTime, err := time.Parse(golibs.Time_TIMEISO8601, timeStr)
		if err != nil {
			//common.HttpResponse(c, "publish_CreatePublish", "保存构建信息失败", nil)
			logs.Error("【%s】时间格式错误: %v", "Publish", err)
		}
		return newTime
	} else {
		timeStr = "0001-01-01T00:00:00.000-0700"
		newTime, err := time.Parse(golibs.Time_TIMEISO8601, timeStr)
		if err != nil {
			//common.HttpResponse(c, "publish_CreatePublish", "保存构建信息失败", nil)
			logs.Error("【%s】时间格式错误: %v", "Publish", err)
		}
		return newTime
	}

}

func TimeFormatStandard(timeStr string) (time.Time, error) {
	if timeStr != "" || len(timeStr) > 0 {
		newTime, err := time.Parse(golibs.Time_TIMEStandard, timeStr)
		if err != nil {
			logs.Error("【%s】时间格式错误: %v", "Publish", err)
		}
		return newTime, nil
	} else {
		timeStr = "0001-01-01 00:00:00"
		newTime, err := time.Parse(golibs.Time_TIMEStandard, timeStr)
		if err != nil {
			logs.Error("【%s】时间格式错误: %v", "Publish", err)
		}
		return newTime, errors.New("时间格式不正确")
	}

}
