package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// このテストは/tmp/gcal_forcerun_secret.jsonが存在する場合のみ実行される
func TestRunner(t *testing.T) {
	config := NewConfig("/tmp/gcal_run_test/credential.json", "/tmp/gcal_run_test", 30, "/Applications/Google Chrome.app")
	err := config.IsValid()
	assert.NoError(t, err)
	runner := NewRunner(config)
	err = runner.Run()
	assert.NoError(t, err)
}
