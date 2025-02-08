package gcal_run

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSloggerWrapper(t *testing.T) {
	logger, err := GetLogger("")
	assert.NoError(t, err)
	logger.Debug("debug %s", "test")
	logger.Info("info %s", "test")
	logger.Warn("warn %s", "test")
	logger.Error("error %s", "test")
	assert.NotNil(t, logger)
}
func TestGetSloggerWrapperDebug(t *testing.T) {
	os.Setenv("DEBUG", "1")
	logger, err := GetLogger("")
	assert.NoError(t, err)
	logger.Debug("debug %s", "test")
	logger.Info("info %s", "test")
	logger.Warn("warn %s", "test")
	logger.Error("error %s", "test")
	os.Unsetenv("DEBUG")
	assert.NotNil(t, logger)
}
