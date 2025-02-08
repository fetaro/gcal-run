package installer

import (
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

type Installer struct {
}

func NewInstaller() *Installer {
	return &Installer{}
}
func (i *Installer) ScanInput() *common.Config {
	var err error

	var credPath string
	for {
		credPath = PrintAndScanStdInput("GoogleカレンダーAPIのクレデンシャルパスを指定してください > ")
		if !common.FileExists(credPath) {
			fmt.Println("GoogleカレンダーAPIのクレデンシャルパスを指定してください。再度入力してください")
		} else {
			break
		}
	}

	var browserApp string
	for {
		browserApp = PrintAndScanStdInput(fmt.Sprintf("ブラウザアプリケーションのパスを指定してください\nデフォルトは「%s」です。デフォルトで良い場合は何も入力せずにEnterを押してください\n> ", common.DefaultBrowserApp))
		if browserApp == "" {
			browserApp = common.DefaultBrowserApp
			break
		} else if !common.FileExists(browserApp) {
			fmt.Println("ブラウザアプリケーションが存在しません。再度入力してください")
		} else {
			break
		}
	}

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
	return common.NewConfig(credPath, minutesAgo, browserApp)
}

func (i *Installer) Install(config *common.Config, appDir string) {
	if !common.FileExists(appDir) {
		err := os.MkdirAll(appDir, 0755)
		if err != nil {
			panic(fmt.Errorf("ディレクトリを作成できませんでした: %v\n", err))
		}
		fmt.Printf("インストール先ディレクトリを作成しました: %s\n", appDir)
	} else {
		fmt.Printf("インストール先ディレクトリが既に存在します。: %s\n", appDir)
		if PrintAndScanStdInput("ここにインストールしますか？ (y/n) > ") != "y" {
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
	fmt.Println("ツールを、インストールディレクトリにコピーし、実行権限を付与します")
	// ./gcal_run, ./installerをコピー
	var binPaths []string
	if common.IsWindows() {
		binPaths = []string{"gcal_run.exe", "installer.exe"}
	} else {
		binPaths = []string{"gcal_run", "installer"}
	}
	for _, binFileName := range binPaths {
		if !common.FileExists(binFileName) {
			panic(fmt.Errorf("installerのファイルの隣にあるはずの実行ファイル「gcal_run」が見つかりません: %s\n", err))
		}
		// ファイルをコピー
		dstBinFile := filepath.Join(appDir, binFileName)
		fmt.Printf("バイナリファイル %s を %s にコピーします\n", binFileName, dstBinFile)
		err = CopyFile(binFileName, dstBinFile)

		if !common.IsWindows() {
			// バイナリファイルに実行権限を付与
			fmt.Printf("バイナリファイル %s に実行権限を付与します\n", dstBinFile)
			stdOutErr, err := exec.Command("chmod", "+x", dstBinFile).CombinedOutput()
			fmt.Println(string(stdOutErr))
			if err != nil {
				panic(err)
			}
		}
	}
	if !common.IsWindows() {
		// plistファイルを作成
		err = NewDaemonCtrl().CreatePListFile(true)
		if err != nil {
			panic(err)
		}
	}

	// トークンの取得
	tokenPath := common.GetTokenPath(appDir)
	_, err = NewOAuthTokenGetter(true).GetAndSaveToken(config.CredentialPath, tokenPath, config.BrowserApp)
	if err != nil {
		panic(err)
	}

	fmt.Println("インストールが完了しました。")

	if PrintAndScanStdInput("常駐プロセスを起動しますか？ (y/n) > ") == "y" {
		daemonCtrl := NewDaemonCtrl()
		err := daemonCtrl.StartDaemon()
		if err != nil {
			panic(err)
		}
		fmt.Println("常駐プロセスを起動しました")
		fmt.Println("2秒待ちます")
		time.Sleep(2 * time.Second)
		isRunning, err := daemonCtrl.IsDaemonRunning()
		if err != nil {
			panic(err)
		}
		if !isRunning {
			panic("常駐プロセスが起動していません")
		}
		fmt.Println("常駐プロセスが動いていることを確認しました")
		fmt.Println("常駐プロセスのログは以下のコマンドで確認できます")
		fmt.Printf("tail -f %s\n", common.GetLogPath(appDir))
	}
	return
}
