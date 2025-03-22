package common

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	DefaultMinutesAgo  = 2
	PoolingIntervalSec = 30
	ToolName           = "GoogleカレンダーTV会議強制起動ツール"
)

func GetAppDir() string {
	if IsWindows() {
		return filepath.Join(os.Getenv("APPDATA"), "gcal_run")
	} else {
		return filepath.Join(os.Getenv("HOME"), ".gcal_run")
	}
}

func GetBinPath(appDir string) string {
	if IsWindows() {
		return filepath.Join(appDir, "gcal_run.exe")
	} else {
		return filepath.Join(appDir, "gcal_run")
	}
}
func GetWinIconPath(appDir string) string {
	return filepath.Join(appDir, "gcal_run.ico")
}

func GetLogPath(appDir string) string {
	return filepath.Join(appDir, "gcal_run.log")
}

func GetServiceLogPath(appDir string) string {
	return filepath.Join(appDir, "service.log")
}

func GetTokenPath(appDir string) string {
	return filepath.Join(appDir, "oauth_token")
}

func GetEventIDStorePath(appDir string) string {
	return filepath.Join(appDir, "event_id_store")
}

func GetConfigPath(appDir string) string {
	return filepath.Join(appDir, "config.json")
}

func GetWinDesktopShortcutPath() string {
	return filepath.Join(os.Getenv("HOMEPATH"), "Desktop", "gcal_run.lnk")
}
func GetWinStartupShortcutPath() string {
	return filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Startup", "gcal_run.lnk")
}

func SJisToUtf8(str string) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}
