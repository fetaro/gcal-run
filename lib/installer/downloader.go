package installer

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type Downloader struct {
}

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (u *Downloader) downaloadRelease(version *Version) string {
	fileName := fmt.Sprintf("gcal-run_%s_%s_v%s.tar.gz", runtime.GOOS, runtime.GOARCH, version.String())
	downloadGzPath := filepath.Join("/tmp/", fileName)
	url := fmt.Sprintf("https://github.com/fetaro/gcal-run/releases/download/v%s/%s", version.String(), fileName)
	fmt.Printf("GitHubからプログラムのダウンロード. URL: %s\n", url)
	// HTTPリクエストを作成
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// ファイルを作成
	out, err := os.Create(downloadGzPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// レスポンスボディをファイルに書き込む
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ダウンロード完了 %s\n", downloadGzPath)
	return downloadGzPath
}

func (u *Downloader) DownloadAndCopy(gitVersion *Version, appDir string) {
	downloadedGzPath := u.downaloadRelease(gitVersion)

	// tar zxvf downloadedGzPath -C /tmp のコマンドを実行
	fmt.Printf("tar zxvf %s -C /tmp\n", downloadedGzPath)
	stdOutErr, err := exec.Command("tar", "zxvf", downloadedGzPath, "-C", "/tmp").CombinedOutput()
	fmt.Println(string(stdOutErr))
	if err != nil {
		panic(err)
	}

	// decompressedDirの中身をinstallDirにコピー
	decompressedDir := strings.Replace(downloadedGzPath, ".tar.gz", "", 1)
	fmt.Printf("%s の中身を %s にコピーします\n", decompressedDir, appDir)
	err = CopyDir(decompressedDir, appDir)
	if err != nil {
		panic(err)
	}
	
	// バイナリファイルに実行権限を付与
	for _, binFileName := range []string{"gcal_run", "installer"} {
		binPath := filepath.Join(appDir, binFileName)
		fmt.Printf("chmod +x %s\n", binPath)
		stdOutErr, err = exec.Command("chmod", "+x", binPath).CombinedOutput()
		fmt.Println(string(stdOutErr))
		if err != nil {
			panic(err)
		}
	}

	// decompressedDirを削除
	err = os.RemoveAll(decompressedDir)

	// ダウンロードしたファイルを削除
	err = os.Remove(downloadedGzPath)
	if err != nil {
		panic(err)
	}
}
