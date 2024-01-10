package lib

import (
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOAuthTokenGetter_GetAndSaveToken(t *testing.T) {
	tokenGetter := NewOAuthTokenGetter()
	tokenPath := path.Join("/tmp", time.Now().Format("20060102150405"))
	token, err := tokenGetter.GetAndSaveToken("/tmp/gcal_run_test/credential.json", tokenPath, "/Applications/Google Chrome.app")
	if err != nil {
		t.Skip("Skip this test because credential.json and oauth_token is not found")
	}
	assert.NotEmpty(t, token)
	defer os.Remove(tokenPath)
}
