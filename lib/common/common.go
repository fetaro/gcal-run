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

func GetAppDir() string {
	return path.Join(os.Getenv("HOME"), ".gcal_run")
}

func GetBinPath(appDir string) string {
	return path.Join(appDir, "gcal_run")
}

func GetLogPath(appDir string) string {
	return path.Join(appDir, "gcal_run.log")
}

func GetTokenPath(appDir string) string {
	return path.Join(appDir, "oauth_token")
}

func GetEventIDStorePath(appDir string) string {
	return path.Join(appDir, "event_id_store")
}

func GetConfigPath(appDir string) string {
	return path.Join(appDir, "config.json")
}
