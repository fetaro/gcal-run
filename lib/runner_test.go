package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// このテストは/tmp/gcal_forcerun_secret.jsonが存在する場合のみ実行される
func TestRunner(t *testing.T) {
	config := NewConfig("/tmp/gcal_run_test/credential.json", "/tmp/gcal_run_test", 30, "/Applications/Google Chrome.app")
	err := config.IsValid()
	if err != nil {
		t.Skip("Skip this test because credential.json and oauth_token is not found")
	}
	runner := NewRunner(config)
	err = runner.Run()
	assert.NoError(t, err)
}
