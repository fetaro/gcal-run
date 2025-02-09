package common

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	assert.True(t, FileExists("/tmp"))
	// ファイルの作成
	f, err := os.Create("/tmp/test_file_from_gcal_run_common_util_test")
	assert.NoError(t, err)
	assert.True(t, FileExists("/tmp/test_file_from_gcal_run_common_util_test"))
	f.Close()
	assert.False(t, FileExists("/does_not_exist_dir/"))
	assert.False(t, FileExists("/does_not_exist_file"))

}

func TestOpenUrl(t *testing.T) {
	err := OpenUrl(DefaultBrowserApp, "https://www.google.com")
	assert.NoError(t, err)
}
