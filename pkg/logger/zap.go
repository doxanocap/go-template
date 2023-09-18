package logger

import (
	"app/internal/cns"
	"app/internal/config"
	"os"
)

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ExtractContextFunc func(ctx context.Context) (key string, value interface{})

var Log *zap.Logger

func InitLogger(cfg *config.Cfg) *zap.Logger {
	Log = initLogger(cfg.AppEnvironment == cns.AppProductionEnv, cfg.LogJSON)
	return Log
}

func initLogger(isProduction bool, isJSON bool) *zap.Logger {
	writer := zapcore.Lock(os.Stdout)
	encoder := getEncoder(isProduction, isJSON)
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	return zap.New(core)
}

func getEncoder(isProduction bool, isJSON bool) zapcore.Encoder {
	encoderConfig := getEncoderConfig(isProduction)
	encoderConfig.EncodeLevel = getLevelEncoding(isJSON)
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "message"

	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getEncoderConfig(isProduction bool) zapcore.EncoderConfig {
	if isProduction {
		return zap.NewProductionEncoderConfig()
	}
	return zap.NewDevelopmentEncoderConfig()
}

func getLevelEncoding(isJSON bool) zapcore.LevelEncoder {
	if isJSON {
		return zapcore.LowercaseLevelEncoder
	}
	return zapcore.LowercaseColorLevelEncoder
}
