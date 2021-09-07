package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func New() *zap.Logger {
	return logger
}

func init() {
	path := createDirectory()
	writer := getWriter(path)
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(logger)
}

func createDirectory() string {
	var err error

	path, err := os.Getwd()
	if err != nil {
		log.Fatal("logger | createDirectory | Getwd | err: ", err)
	}

	_, err = os.Stat(fmt.Sprintf("%s/logs", path))
	if os.IsNotExist(err) {
		err = os.Mkdir("logs", os.ModePerm)
	}

	if err != nil {
		log.Fatal("logger | createDirectory | Mkdir | err: ", err)
	}

	return path
}

func getWriter(path string) zapcore.WriteSyncer {
	path = path + "/logs/all.log"

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("logger | getWriter | OpenFile | err: ", err)
	}

	return zapcore.AddSync(file)
}

func getEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:0405Z0700"))
	})

	cfg.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(cfg)
}
