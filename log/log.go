package log

import (
	"gin/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugarLogger *zap.SugaredLogger
var logConfig config.LogConfig

/*Init : 初始化日志*/
func Init(c config.LogConfig) {
	logConfig = c
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	var syncer zapcore.WriteSyncer
	logInConsole := logConfig.Console
	if logInConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writeSyncer)
	} else {
		syncer = writeSyncer
	}
	core := zapcore.NewCore(encoder, syncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./runtime/logs/" + logConfig.Filename + ".log",
		MaxSize:    2,
		MaxBackups: 5,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

/*Info : 打印信息日志*/
func Info(template string, args ...interface{}) {
	sugarLogger.Infof(template, args)
}

/*Error : 打印错误日志*/
func Error(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args)
}
