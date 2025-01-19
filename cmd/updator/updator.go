package main

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fetaro/gcal_forcerun_go/lib"
)

type Release struct {
	ID      int    `json:"id"`
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Body    string `json:"body"`
}

func downaloadRelease(version *lib.Version) string {
	fileName := fmt.Sprintf("gcal-run_%s_%s_v%s.tar.gz", runtime.GOOS, runtime.GOARCH, version.String())
	url := fmt.Sprintf("https://github.com/fetaro/gcal-run/releases/download/v%s/%s", version.String(), fileName)
	fmt.Printf("GitHubからプログラムのダウンロード. URL: %s\n", url)
	// HTTPリクエストを作成
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// ファイルを作成
	out, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// レスポンスボディをファイルに書き込む
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	return fileName
}

func decompressTarGz(gzipPath string, dest string) {
	fmt.Printf("解凍 %s\n", gzipPath)
	// gzipファイルを開く
	gzipFile, err := os.Open(gzipPath)
	if err != nil {
		panic(err)
	}
	defer gzipFile.Close()

	// gzipリーダーを作成
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		panic(err)
	}
	defer gzipReader.Close()

	// tarリーダーを作成
	tarReader := tar.NewReader(gzipReader)

	// tarファイルの中の各ファイルを処理
	for {
		header, err := tarReader.Next()

		switch {
		case err == io.EOF:
			return
		case err != nil:
			panic(err)
		case header == nil:
			continue
		}

		target := filepath.Join(dest, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					panic(err)
				}
			}
		case tar.TypeReg:
			file, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				panic(err)
			}
			defer file.Close()
			if _, err := io.Copy(file, tarReader); err != nil {
				panic(err)
			}
		}
	}
}

func fileCopy(srcPath string, dstPath string) {
	// ファイルをコピー
	in, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.Create(dstPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		panic(err)
	}
}

func main() {
	// 引数の数をチェック
	var installDir string
	if len(os.Args) == 2 {
		installDir = os.Args[1]
	} else {
		installDir = lib.DefaultInstallDir()
	}
	binFilePath := filepath.Join(installDir, "gcal_run")
	// binファイルが存在するかチェック
	_, err := os.Stat(binFilePath)
	if os.IsNotExist(err) {
		fmt.Printf("インストールされているバイナリが見つかりません. 探したパス: %s\n", binFilePath)
		fmt.Println("インストールディレクトリをデフォルトから変更している場合は、第一引数にインストールディレクトリを指定してください")
		fmt.Println("使い方 : updator /path/to/install/dir")
		os.Exit(1)
	}
	// binFilePathのバイナリを --version の引数を付けて実行し、バージョンを取得する
	fmt.Printf("バイナリのバージョンを取得します: %s --version\n", binFilePath)
	stdOutErr, err := exec.Command(binFilePath, "--version").CombinedOutput()
	if err != nil {
		fmt.Printf("バージョンの取得に失敗しました。エラー： %v\n", err)
		os.Exit(1)
	}
	versonStr := string(stdOutErr)
	fmt.Printf("バージョン: %s\n", versonStr)
	installedVersion, err := lib.ParseVersionStr(versonStr)
	if err != nil {
		fmt.Printf("バージョンのパースに失敗しました。エラー: %v\n", err)
		os.Exit(1)
	}

	// Githubのリリースのバージョンを取得する
	url := "https://api.github.com/repos/fetaro/gcal-run/releases"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("GitHubのAPIからリリース情報の取得に失敗しました。URL:%s error:%v\n", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("GitHubのAPIからリリース情報のリクエストボディの取得に失敗しました。%v\n", err)
		os.Exit(1)
	}

	var releases []Release
	err = json.Unmarshal(body, &releases)
	if err != nil {
		fmt.Printf("GitHubのAPIからリリース情報のJSONのパースに失敗しました。json=%s, error=%v\n", err)
		os.Exit(1)
	}
	latestTagName := releases[0].TagName
	gitVersion, err := lib.ParseVersionStr(latestTagName)
	fmt.Printf("インストールされているバージョン: %s\n", installedVersion)
	fmt.Printf("最新のバージョン: %s\n", gitVersion)

	if gitVersion.IsNewer(installedVersion) {
		fmt.Printf("新しいバージョンがあります: %s\n", latestTagName)
		// インストールする
		fmt.Printf("プログラムを更新しますか y/n: %s\n")
		scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
		scanner.Scan()
		yOrN := scanner.Text()
		if yOrN == "y" {
			downloadTarGzFileName := downaloadRelease(gitVersion)
			// 解凍する
			decompressTarGz(downloadTarGzFileName, ".")
			downloadDirName := strings.ReplaceAll(downloadTarGzFileName, ".tar.gz", "")
			entries, err := ioutil.ReadDir(downloadDirName)
			if err != nil {
				panic(err)
			}
			// プログラムの停止

			// ファイルをコピー
			for _, entry := range entries {
				src := filepath.Join(downloadDirName, entry.Name())
				dst := entry.Name()
				fmt.Printf("cp %s %s\n", src, dst)
				fileCopy(src, dst)
			}
			fmt.Println("アップデート正常終了")
		} else {
			fmt.Println("中止しました")
		}
	} else {
		fmt.Printf("インストールされているバージョンは最新のバージョンです: %s\n", latestTagName)
		os.Exit(0)
	}

}
