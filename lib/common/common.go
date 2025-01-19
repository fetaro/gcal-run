package common

import (
	"os"
	"path"
)

const (
	DefaultBrowserApp = "/Applications/Google Chrome.app"
	DefaultMinutesAgo = 2
)

func GetPListPath() string {
	// ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist を返す
	return path.Join(os.Getenv("HOME"), "Library/LaunchAgents/com.github.fetaro.gcal_run.plist")
}

func DefaultInstallDir() string {
	return path.Join(os.Getenv("HOME"), ".gcal_run")
}
