package installer

import (
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"github.com/fetaro/gcal_forcerun_go/lib/gcal_run"
	"os"
	"strconv"
)

type Installer struct {
}

func NewInstaller() *Installer {
	return &Installer{}
}
func (i *Installer) ScanInput() (int, string) {

	var browserApp string
	for {
		browserApp = PrintAndScanStdInput(fmt.Sprintf("ブラウザアプリケーションのパスを指定してください\nデフォルトは「%s」です。デフォルトで良い場合は何も入力せずにEnterを押してください\n> ", common.DefaultBrowserApp))
		_, err := os.Stat(browserApp)
		if browserApp == "" {
			browserApp = common.DefaultBrowserApp
			break
		}
		if os.IsNotExist(err) {
			fmt.Println("ブラウザアプリケーションが存在しません。再度入力してください")
		} else {
			break
		}
	}
	var err error
	var minutesAgoStr string
	var minutesAgo int
	for {
		minutesAgoStr = PrintAndScanStdInput(fmt.Sprintf("会議の何分前に起動するか指定してください\nデフォルトは「%d分」です。デフォルトで良い場合は何も入力せずにEnterを押してください\n> ", common.DefaultMinutesAgo))
		if minutesAgoStr == "" {
			minutesAgo = common.DefaultMinutesAgo
			break
		}
		minutesAgo, err = strconv.Atoi(minutesAgoStr)
		if err != nil {
			fmt.Println("数値を入力してください")
			continue
		} else {
			break
		}
	}
	return minutesAgo, browserApp
}

func (i *Installer) Install(config *common.Config, appDir string) {
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		err := os.MkdirAll(appDir, 0755)
		if err != nil {
			panic(fmt.Errorf("ディレクトリを作成できませんでした: %v\n", err))
		}
		fmt.Printf("インストール先ディレクトリを作成しました: %s\n", appDir)
	} else {
		fmt.Printf("インストール先ディレクトリが既に存在します。: %s\n", appDir)
		if PrintAndScanStdInput("中身を空にして、インストールしますか？ (y/n) > ") == "y" {
			// installDirの中身を空にする
			err := os.RemoveAll(appDir)
			if err != nil {
				panic(fmt.Errorf("ディレクトリを空にできませんでした: %v\n", err))
			}
			err = os.MkdirAll(appDir, 0755)
			if err != nil {
				panic(fmt.Errorf("ディレクトリを作成できませんでした: %v\n", err))
			}
			fmt.Printf("ディレクトリを空にして再作成しました: %s\n", appDir)

		} else {
			fmt.Println("インストールを中止します")
			os.Exit(1)
		}
	}
	// 設定の保存
	err := config.Save()
	if err != nil {
		panic(fmt.Errorf("設定の保存に失敗しました: %v\n", err))
	}
	// ツールのダウンロード
	fmt.Println("ツールをダウンロードし、インストールディレクトリに展開します")
	latestVersion, err := NewGithubService().GetLatestVersion()
	if err != nil {
		panic(err)
	}
	NewDownloader().DownloadAndCopy(latestVersion, appDir)

	// plistファイルを作成
	err = NewDaemonCtrl().CreatePListFile()
	if err != nil {
		panic(err)
	}

	// トークンの取得
	tokenPath := common.GetTokenPath(appDir)
	_, err = gcal_run.NewOAuthTokenGetter().GetAndSaveToken(config.CredentialPath, tokenPath, config.BrowserApp)
	if err != nil {
		panic(err)
	}

	fmt.Printf(`
============================================
インストールが完了しました。

インストールディレクトリ : %s
常駐プロセス(LaunchAgents)ファイル : %s
`, appDir, NewDaemonCtrl().GetPListPath())

	return
}
