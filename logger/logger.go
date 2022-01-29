package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	return log
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	if err := log.Sync(); err != nil {
		fmt.Println(err)
	}
}

func Debug(msg string, tags ...zap.Field) {
	log.Debug(msg, tags...)
	if err := log.Sync(); err != nil {
		fmt.Println(err)
	}
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	if err := log.Sync(); err != nil {
		fmt.Println(err)
	}
}

func Critical(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("critical", err))
	log.Error(msg, tags...)
	if err := log.Sync(); err != nil {
		fmt.Println(err)
	}
}
