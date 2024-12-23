package foundation

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Z *zap.Logger
}

type LoggerConfig struct {
	Level       string `default:"error"`
	ServiceName string `default:"service"`
}

// NewLogger will start and return a zap logger
func NewLogger(cfg *LoggerConfig) (*Logger, error) {
	var zapConfig zap.Config

	if cfg.Level == "debug" {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		zapConfig = zap.NewProductionConfig()
		// zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	zapConfig.OutputPaths = []string{"stdout"}
	zapConfig.DisableStacktrace = true
	zapConfig.InitialFields = map[string]interface{}{
		"service": cfg.ServiceName,
	}
	logger, err := zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return &Logger{
		Z: logger,
	}, nil
}

func (l Logger) Info(args ...interface{}) {
	l.Z.Sugar().Info(args...)
}

func (l Logger) Infof(msg string, args ...interface{}) {
	l.Z.Sugar().Infof(msg, args...)
}

func (l Logger) Error(args ...interface{}) {
	l.Z.Sugar().Error(args...)
}

func (l Logger) Errorf(msg string, args ...interface{}) {
	l.Z.Sugar().Errorf(msg, args...)
}
