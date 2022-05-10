package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func New(name string, enviroment string) *Logger {
	logger, _ := newDevelopmentLogger()
	if enviroment == "production" {

		logger, _ = zap.NewProduction()

	}
	logger = logger.Named(name)
	defer logger.Sync()
	return &Logger{logger}

}

func newDevelopmentLogger() (*zap.Logger, error) {

	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return cfg.Build(zap.AddCaller())
}
