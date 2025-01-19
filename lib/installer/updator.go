package installer

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Updator struct {
}

func NewUpdator() *Updator {
	return &Updator{}
}

type Release struct {
	ID      int    `json:"id"`
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Body    string `json:"body"`
}

func (u *Updator) Update(installDir string) {
	versionFilePath := filepath.Join(installDir, "VERSION")
	var installedVersionStr string

	// VERSIONファイルが有るか確認
	if _, err := os.Stat(versionFilePath); os.IsNotExist(err) {
		// VERSIONファイルが無い場合は、バイナリからバージョンを取得する
		// binFilePathのバイナリを --version の引数を付けて実行し、バージョンを取得する
		binFilePath := filepath.Join(installDir, "gcal_run")
		fmt.Printf("インストールされているツールのバージョンをバイナリファイルから取得します: %s --version\n", binFilePath)
		stdOutErr, err := exec.Command(binFilePath, "--version").CombinedOutput()
		if err != nil {
			fmt.Printf("インストールされているツールのバージョンの取得に失敗しました。エラー： %v\n", err)
			os.Exit(1)
		}
		installedVersionStr = string(stdOutErr)
	} else {
		fmt.Printf("インストールされているバージョンをバージョンファイルから取得します: %s\n", versionFilePath)
		// ./VERSION ファイルの中身を読む
		file, err := os.Open(versionFilePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		binary, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		installedVersionStr = string(binary)
	}
	// 最後に改行が入っている場合は削除
	installedVersionStr = strings.TrimRight(installedVersionStr, "\n")
	fmt.Printf("インストールされているバージョン: %s\n", installedVersionStr)
	installedVersion, err := ParseVersionStr(installedVersionStr)
	if err != nil {
		fmt.Printf("インストールされているツールのバージョンのパースに失敗しました。エラー: %v\n", err)
		os.Exit(1)
	}
	// Githubのリリースのバージョンを取得する
	githubService := NewGithubService()
	gitVersion, err := githubService.GetLatestVersion()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GitHubの最新のバージョン: %s\n", gitVersion)

	if gitVersion.IsNewer(installedVersion) {
		// インストールする
		fmt.Printf("プログラムを更新しますか (y/n) >")
		scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
		scanner.Scan()
		yOrN := scanner.Text()
		if yOrN == "y" {
			// installDirの中身を空にする
			fmt.Printf("インストールディレクトリを空にします: %s\n", installDir)
			err := os.RemoveAll(installDir)
			if err != nil {
				panic(err)
			}
			// ダウンロードして解凍
			fmt.Println("新しいバージョンをインストールディレクトリにダウンロードします")
			NewDownloader().Download(gitVersion, installDir)
			fmt.Println("アップデート正常終了")
		} else {
			fmt.Println("中止しました")
		}
	} else {
		fmt.Println("インストールされているバージョンは最新のバージョンなので、アップデートは不要です")
		os.Exit(0)
	}

}
