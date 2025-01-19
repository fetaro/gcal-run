package gcal_run

import (
	"os"
	"path/filepath"
	"runtime"
)

func OsUserCacheDir() string {
	switch runtime.GOOS {
	// TODO: Windowsの場合の処理を追加する
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Caches")
	case "linux", "freebsd":
		return filepath.Join(os.Getenv("HOME"), ".cache")
	}
	return "."
}
