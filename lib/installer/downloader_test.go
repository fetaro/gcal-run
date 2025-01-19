package installer

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloader_Download(t *testing.T) {
	d := NewDownloader()
	tempDir, err := os.MkdirTemp("", "example")
	if err != nil {
		fmt.Printf("Failed to create temporary directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir)

	d.Download(NewVersion(1, 1, 0), filepath.Join(tempDir))
}
