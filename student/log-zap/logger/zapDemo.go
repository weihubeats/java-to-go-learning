package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"
)

var logger *zap.Logger

func main() {
	InitLogger()
	defer logger.Sync()
	for true {
		logger.Error("test error")
	}
	/*simpleHttpGet("www.google.com")
	simpleHttpGet("https://www.baidu.com")*/
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller())
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("error fetching url..", zap.String("url", url), zap.Error(err))

	} else {
		logger.Info("success..", zap.String("statuCode", resp.Status), zap.String("url", url))
	}
}

/*func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)

}*/

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	//config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(config)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
