package lib

import (
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// このテストは/tmp/gcal_forcerun_secret.jsonが存在する場合のみ実行される
func TestInstaller_Install(t *testing.T) {
	dirName := time.Now().Format("20060102150405")
	installDir := path.Join(os.TempDir(), dirName)
	config := NewConfig("/tmp/gcal_run_test/credential.json", installDir, 2, "/Applications/Google Chrome.app")
	installer := NewInstaller()
	err := installer.Install(config)
	assert.NoError(t, err)

}
