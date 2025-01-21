package installer

import (
	"fmt"
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
	var installedVersionStr string

	// binFilePathのバイナリを --version の引数を付けて実行し、バージョンを取得する
	binFilePath := filepath.Join(installDir, "gcal_run")
	fmt.Printf("インストールされているツールのバージョンをバイナリファイルから取得します: %s --version\n", binFilePath)
	stdOutErr, err := exec.Command(binFilePath, "--version").CombinedOutput()
	if err != nil {
		panic(fmt.Errorf("インストールされているツールのバージョンの取得に失敗しました。エラー： %v\n", err))
	}
	installedVersionStr = string(stdOutErr)

	// 最後に改行が入っている場合は削除
	installedVersionStr = strings.TrimRight(installedVersionStr, "\n")
	fmt.Printf("インストールされているバージョン: %s\n", installedVersionStr)
	installedVersion, err := ParseVersionStr(installedVersionStr)
	if err != nil {
		panic(fmt.Errorf("インストールされているツールのバージョンのパースに失敗しました。エラー: %v\n", err))
	}
	// Githubのリリースのバージョンを取得する
	githubService := NewGithubService()
	gitVersion, err := githubService.GetLatestVersion()
	if err != nil {
		panic(fmt.Errorf("GitHubの最新のバージョンの取得に失敗しました: %v\n", err))
	}
	fmt.Printf("GitHubの最新のバージョン: %s\n", gitVersion)

	if gitVersion.IsNewer(installedVersion) {
		if PrintAndScanStdInput("プログラムを更新しますか (y/n) >") == "y" {
			fmt.Println("新しいバージョンをインストールディレクトリにコピーします")
			NewDownloader().DownloadAndCopy(gitVersion, installDir)

			//plistファイルを更新します
			err = NewDaemonCtrl().CreatePListFile(true)
			if err != nil {
				panic(err)
			}

			// 再起動
			daemonCtrl := NewDaemonCtrl()
			err = daemonCtrl.StopDaemon()
			if err != nil {
				panic(fmt.Errorf("常駐プロセスの停止に失敗しました: %v\n", err))
			}
			err = daemonCtrl.StartDaemon()
			if err != nil {
				panic(fmt.Errorf("常駐プロセスの起動に失敗しました。マニュアルに従って手動で起動してください: %v\n", err))
			}
			fmt.Println("常駐プロセスを再起動しました")
			fmt.Println("アップデート正常終了")
		} else {
			fmt.Println("中止しました")
		}
	} else {
		fmt.Println("インストールされているバージョンは最新のバージョンなので、アップデートは不要です")
		os.Exit(0)
	}

}
