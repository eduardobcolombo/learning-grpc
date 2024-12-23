package api

import (
	"testing"

	"github.com/eduardobcolombo/learning-grpc/foundation"
	mock_portpb "github.com/eduardobcolombo/learning-grpc/portpb/mock_port"
	"github.com/stretchr/testify/assert"
)

func TestNewCoreConfig(t *testing.T) {
	log, _ := foundation.NewLogger(&foundation.LoggerConfig{})
	psc := &mock_portpb.MockPortServiceClient{}
	coreConfig := NewCoreConfig(log, psc)

	assert.NotNil(t, coreConfig)
	assert.NotNil(t, coreConfig.Port)
}
