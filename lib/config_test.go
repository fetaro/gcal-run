package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_GeneratePlistStr(t *testing.T) {
	config := NewConfig("credpath", "installDir", 30, "/Applications/Google Chrome.app")
	actual := config.GeneratePlistStr()
	assert.Contains(t, actual, "credpath")
	assert.Contains(t, actual, "installDir/gcal_run")
	assert.Contains(t, actual, "installDir/gcal_run.log")
	assert.Contains(t, actual, "/Applications/Google Chrome.app")
	assert.Contains(t, actual, "30")
}
