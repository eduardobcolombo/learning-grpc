package foundation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *LoggerConfig
		wantErr bool
	}{
		{
			name:    "default config",
			cfg:     &LoggerConfig{},
			wantErr: false,
		},
		{
			name:    "debug level",
			cfg:     &LoggerConfig{Level: "debug"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := NewLogger(tt.cfg)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, logger)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, logger)
			}
		})
	}
}
func TestLogger_Errorf(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *LoggerConfig
		message string
		args    []interface{}
		wantErr bool
	}{
		{
			name:    "error message with args",
			cfg:     &LoggerConfig{Level: "error"},
			message: "error occurred: %s",
			args:    []interface{}{"file not found"},
			wantErr: false,
		},
		{
			name:    "error message without args",
			cfg:     &LoggerConfig{Level: "error"},
			message: "error occurred",
			args:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := NewLogger(tt.cfg)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, logger)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, logger)
				logger.Errorf(tt.message, tt.args...)
			}
		})
	}
}
func TestLogger_Error(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *LoggerConfig
		args    []interface{}
		wantErr bool
	}{
		{
			name:    "error message with args",
			cfg:     &LoggerConfig{Level: "error"},
			args:    []interface{}{"file not found"},
			wantErr: false,
		},
		{
			name:    "error message without args",
			cfg:     &LoggerConfig{Level: "error"},
			args:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := NewLogger(tt.cfg)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, logger)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, logger)
				logger.Error(tt.args...)
			}
		})
	}
}
func TestLogger_Infof(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *LoggerConfig
		message string
		args    []interface{}
		wantErr bool
	}{
		{
			name:    "info message with args",
			cfg:     &LoggerConfig{Level: "info"},
			message: "info message: %s",
			args:    []interface{}{"operation successful"},
			wantErr: false,
		},
		{
			name:    "info message without args",
			cfg:     &LoggerConfig{Level: "info"},
			message: "info message",
			args:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := NewLogger(tt.cfg)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, logger)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, logger)
				logger.Infof(tt.message, tt.args...)
			}
		})
	}
}
func TestLogger_Info(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *LoggerConfig
		args    []interface{}
		wantErr bool
	}{
		{
			name:    "info message with args",
			cfg:     &LoggerConfig{Level: "info"},
			args:    []interface{}{"operation successful"},
			wantErr: false,
		},
		{
			name:    "info message without args",
			cfg:     &LoggerConfig{Level: "info"},
			args:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := NewLogger(tt.cfg)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, logger)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, logger)
				logger.Info(tt.args...)
			}
		})
	}
}
