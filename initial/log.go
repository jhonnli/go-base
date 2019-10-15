package initial

import (
	"github.com/jhonnli/go-base/initial/config"
	//"github.com/jhonnli/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLog() {

	var level zapcore.Level
	switch config.Config.LogConfig.LogLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	logger, _ := zap.Config{
		Encoding:         config.Config.LogConfig.Encoding,                                           // 输出格式 console 或 json
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),                                          // 编码器配置
		Level:            zap.NewAtomicLevelAt(level),                                                // 日志级别
		InitialFields:    map[string]interface{}{"serviceName": config.Config.LogConfig.ServiceName}, // 初始化字段，如：服务器名称
		OutputPaths:      []string{"stdout", config.Config.LogConfig.StdoutPath},                     // 输出到指定文件 stdout（标准输出，正常颜色）
		ErrorOutputPaths: []string{"stderr", config.Config.LogConfig.StderrPath},                     // 输出到指定文件 stderr（错误输出，红色）
	}.Build()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	Log = zap.L()
}
