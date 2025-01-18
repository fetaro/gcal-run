package main

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
	fmt.Printf("ダウンロード. URL: %s\n", url)
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

	// ./VERSION ファイルの中身を読む
	file, err := os.Open("./VERSION")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	binary, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	installedVersion, err := lib.ParseVersionStr(string(binary))
	if err != nil {
		panic(err)
	}

	// Githubのリリースのバージョンを取得する
	resp, err := http.Get("https://api.github.com/repos/fetaro/gcal-run/releases")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var releases []Release
	err = json.Unmarshal(body, &releases)
	if err != nil {
		panic(err)
	}
	latestTagName := releases[0].TagName
	gitVersion, err := lib.ParseVersionStr(latestTagName)
	fmt.Printf("インストールされているバージョン: %s\n", installedVersion)
	fmt.Printf("最新のバージョン: %s\n", gitVersion)

	if gitVersion.IsNewer(installedVersion) {
		fmt.Printf("新しいバージョンがあります: %s\n", latestTagName)
		downloadTarGzFileName := downaloadRelease(gitVersion)
		// 解凍する
		decompressTarGz(downloadTarGzFileName, ".")
		// インストールする
		fmt.Printf("プログラムを更新しますか y/n: %s\n")
		downloadDirName := strings.ReplaceAll(downloadTarGzFileName, ".tar.gz", "")
		entries, err := ioutil.ReadDir(downloadDirName)
		if err != nil {
			panic(err)
		}
		for _, entry := range entries {
			src := filepath.Join(downloadDirName, entry.Name())
			dst := entry.Name()
			fmt.Printf("copy src: %s, dst: %s\n", src, dst)
			//fileCopy(src, dst)
		}
	} else {
		fmt.Printf("インストールされているバージョンは最新のバージョンです: %s\n", latestTagName)
		os.Exit(0)
	}

}
