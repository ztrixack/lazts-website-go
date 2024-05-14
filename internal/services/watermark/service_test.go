package watermark

import (
	"lazts/internal/modules/imaging"
	"lazts/internal/modules/log"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	mockLogger := new(log.Mock)
	mockImager := new(imaging.Mock)
	config := &config{}

	service := New(config, mockLogger, mockImager)

	assert.NotNil(t, service, "Service should not be nil")
	assert.Equal(t, config, service.config, "Config should match the one provided")
	assert.IsType(t, &sync.Mutex{}, service.mutex, "Mutex should be initialized")
	assert.NotNil(t, service.cache, "Cache map should be initialized and not nil")
	assert.Equal(t, mockLogger, service.logger, "Logger should match the mock logger")
	assert.Equal(t, mockImager, service.imager, "Imager should match the mock imager")
}
