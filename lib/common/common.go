package common

import (
	"os"
	"path"
)

const (
	//DefaultBrowserApp = "/Applications/Google Chrome.app"
	DefaultBrowserApp  = "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	DefaultMinutesAgo  = 2
	PoolingIntervalSec = 30
	ToolName           = "GoogleカレンダーTV会議強制起動ツール"
)

func GetAppDir() string {
	if IsWindows() {
		return path.Join(os.Getenv("HOMEPATH"), ".gcal_run")
		// TODO
		//return "c:\\Users\\fetaro\\.gcal_run"
	} else {
		return path.Join(os.Getenv("HOME"), ".gcal_run")
	}
}

func GetBinPath(appDir string) string {
	return path.Join(appDir, "gcal_run")
}
func GetLogPath(appDir string) string {
	//return "c:\\tmp\\gcal_run.log"
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
