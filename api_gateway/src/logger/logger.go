package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func New() *zap.Logger {
	return logger
}

func Init() error {
	path, err := createDirectory()
	if err != nil {
		return fmt.Errorf("createDirectory %w", err)
	}

	writer, err := getWriter(path)
	if err != nil {
		return fmt.Errorf("getWriter %w", err)
	}

	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(logger)

	return nil
}

func createDirectory() (string, error) {
	var err error

	path, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("getwd %w", err)
	}

	_, err = os.Stat(fmt.Sprintf("%s/logs", path))
	if os.IsNotExist(err) {
		err = os.Mkdir("logs", os.ModePerm)
	}

	if err != nil {
		return "", fmt.Errorf("mkdir %w", err)
	}

	return path, nil
}

func getWriter(path string) (zapcore.WriteSyncer, error) {
	path = path + "/logs/all.log"

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("openFile %w", err)
	}

	return zapcore.AddSync(file), nil
}

func getEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:0405Z0700"))
	})

	cfg.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(cfg)
}
