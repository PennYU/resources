package resources

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

var Logger = &zap.SugaredLogger{}

func Conf(fileName string) {
	hook := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    500, // megabytes
		MaxBackups: 10000,
		MaxAge:     100000, //days
		Compress:   false,  // disabled by default
	}
	defer hook.Close()
	enConfig := zap.NewProductionEncoderConfig()
	enConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	level := zap.InfoLevel
	w := zapcore.AddSync(hook)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(enConfig),
		w,
		level,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	_log := log.New(hook, "", log.LstdFlags)
	Logger = logger.Sugar()
	_log.Println("Start...")
}
