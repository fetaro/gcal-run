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

var (
	tmpDir = "/tmp"
)

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (u *Downloader) downaloadRelease(version *Version) string {
	fileName := fmt.Sprintf("gcal-run_%s_%s_v%s.tar.gz", runtime.GOOS, runtime.GOARCH, version.String())
	downloadGzPath := filepath.Join(tmpDir, fileName)
	url := fmt.Sprintf("https://github.com/fetaro/gcal-run/releases/download/v%s/%s", version.String(), fileName)
	fmt.Printf("GitHubからプログラムのダウンロード. URL: %s\n", url)
	fmt.Printf("ダウンロード先: %s\n", downloadGzPath)
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

	fmt.Printf("%s に %s を展開します\n", tmpDir, downloadedGzPath)
	stdOutErr, err := exec.Command("tar", "zxvf", downloadedGzPath, "-C", tmpDir).CombinedOutput()
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
		fmt.Printf("バイナリファイル %s に実行権限を付与します\n", binPath)
		stdOutErr, err = exec.Command("chmod", "+x", binPath).CombinedOutput()
		fmt.Println(string(stdOutErr))
		if err != nil {
			panic(err)
		}
	}

	// decompressedDirを削除
	fmt.Printf("展開したディレクトリ %s を削除します\n", decompressedDir)
	err = os.RemoveAll(decompressedDir)
	if err != nil {
		fmt.Printf("展開したディレクトリを削除できませんでしたが、続行します: %v\n", err)
	}

	// ダウンロードしたファイルを削除
	fmt.Printf("ダウンロードしたファイル %s を削除します\n", downloadedGzPath)
	err = os.Remove(downloadedGzPath)
	if err != nil {
		fmt.Printf("ダウンロードしたファイルを削除できませんでしたが、続行します: %v\n", err)
	}
}
