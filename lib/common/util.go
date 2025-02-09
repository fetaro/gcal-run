package common

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func OpenUrl(browserApp, url string) error {
	if IsWindows() {
		err := exec.Command(browserApp, url).Run()
		if err != nil {
			return fmt.Errorf("failed to open event url: %v", err)
		}
		return err
	} else {
		err := exec.Command("open", "-a", browserApp, url).Run()
		if err != nil {
			return fmt.Errorf("failed to open event url: %v", err)
		}
		return err
	}
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}
