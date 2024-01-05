package lib

import (
	"io"
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

func CopyFile(src, dst string) error {
	w, err := os.Create(dst)
	if err != nil {
		return err
	}

	r, err := os.Open(src)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}
	return nil
}
