package common

import (
	"os"
	"path"
)

const (
	DefaultMinutesAgo  = 2
	PoolingIntervalSec = 30
	ToolName           = "GoogleカレンダーTV会議強制起動ツール"
)

func GetDefaultBrowserApp() string {
	if IsWindows() {
		return "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	} else {
		return "/Applications/Google Chrome.app"
	}
}

func GetAppDir() string {
	if IsWindows() {
		return path.Join(os.Getenv("APPDATA"), "gcal_run")
	} else {
		return path.Join(os.Getenv("HOME"), ".gcal_run")
	}
}

func GetBinPath(appDir string) string {
	return path.Join(appDir, "gcal_run")
}
func GetLogPath(appDir string) string {
	return path.Join(appDir, "gcal_run.log")
}

func GetServiceLogPath(appDir string) string {
	return path.Join(appDir, "service.log")
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
