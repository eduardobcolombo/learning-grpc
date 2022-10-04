package foundation

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger will start and return a zap logger
func NewLogger(service string) (*zap.SugaredLogger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"}
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.DisableStacktrace = true
	cfg.InitialFields = map[string]interface{}{
		"service": service,
	}

	log, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return log.Sugar(), nil
}
