package installer

import (
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
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
	os.Mkdir(installDir, 0755)
	credPath := "/tmp/gcal_run_test/credential.json"
	config := common.NewConfig(credPath, installDir, 2, "/Applications/Google Chrome.app")
	err := config.IsValid()
	if err != nil {
		t.Skip(fmt.Sprintf("Skip this test because %v", err))
	}
	installer := NewInstaller()
	err = installer.Install(config)
	assert.NoError(t, err)

}
