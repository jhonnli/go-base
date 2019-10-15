package initial

import (
	//"github.com/jhonnli/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLog() {
	logger, _ := zap.Config{
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	Log = zap.L()
}
