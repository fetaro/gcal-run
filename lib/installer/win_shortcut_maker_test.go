package installer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWinShortcutMaker_MakeDesktopShortCut(t *testing.T) {
	appDir := "c:\\tmp\\gcal_run_test"
	shortcutPath := "c:\\tmp\\gcal_run_test\\gcal_run.lnk"
	err := NewWinShortcutMaker(appDir).MakeShortCut(shortcutPath)
	assert.NoError(t, err)
	assert.FileExists(t, shortcutPath)
}
