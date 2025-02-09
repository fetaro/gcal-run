package installer

import (
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

type Installer struct {
}

func NewInstaller() *Installer {
	return &Installer{}
}
func (i *Installer) ScanUserInput() *common.Config {
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
	fmt.Println("")

	var browserApp string
	for {
		browserApp = PrintAndScanStdInput(fmt.Sprintf("ブラウザアプリケーションのパスを指定してください\nデフォルトは「%s」です。デフォルトで良い場合は何も入力せずにEnterを押してください\n> ", common.GetDefaultBrowserApp()))
		if browserApp == "" {
			browserApp = common.GetDefaultBrowserApp()
			break
		} else if !common.FileExists(browserApp) {
			fmt.Println("ブラウザアプリケーションが存在しません。再度入力してください")
		} else {
			break
		}
	}
	fmt.Println("")

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
	fmt.Println("")

	return common.NewConfig(credPath, minutesAgo, browserApp)
}

func (i *Installer) Install(config *common.Config, appDir string) error {
	if !common.FileExists(appDir) {
		err := os.MkdirAll(appDir, 0755)
		if err != nil {
			return fmt.Errorf("ディレクトリを作成できませんでした: %v\n", err)
		}
		fmt.Printf("インストール先ディレクトリを作成しました: %s\n", appDir)
	} else {
		fmt.Printf("インストール先ディレクトリが既に存在します。: %s\n", appDir)
		if PrintAndScanStdInput("ここにインストールしますか？ (y/n) > ") != "y" {
			fmt.Println("入力された文字列が'y'ではないため、インストールを中止しました")
			return nil
		}
	}
	// 設定の保存
	err := config.Save()
	if err != nil {
		return fmt.Errorf("設定の保存に失敗しました: %v\n", err)
	}
	fmt.Printf("設定ファイルを作成しました: %s\n", common.GetConfigPath(appDir))

	// ツールのダウンロード
	fmt.Println("ツールをインストールディレクトリにコピーします")
	var filePaths []string
	if common.IsWindows() {
		filePaths = []string{
			"gcal_run.exe",
			"gcal_run.ico",
			"installer.exe",
			"install_startup.ps1",
			"install_desktop_shortcut.ps1"}
	} else {
		filePaths = []string{
			"gcal_run",
			"installer"}
	}
	for _, fileName := range filePaths {
		if !common.FileExists(fileName) {
			return fmt.Errorf("実行ファイル「%s」が見つかりません: %s\n", fileName, err)
		}
		// ファイルをコピー
		dstBinFile := filepath.Join(appDir, fileName)
		fmt.Printf("ファイル %s を %s にコピーします\n", fileName, dstBinFile)
		err = CopyFile(fileName, dstBinFile)

		if !common.IsWindows() {
			// バイナリファイルに実行権限を付与
			fmt.Printf("ファイル %s に実行権限を付与します\n", dstBinFile)
			stdOutErr, err := exec.Command("chmod", "+x", dstBinFile).CombinedOutput()
			fmt.Println(string(stdOutErr))
			if err != nil {
				return err
			}
		}
	}
	if !common.IsWindows() {
		// plistファイルを作成
		err = NewDaemonCtrl().CreatePListFile(true)
		if err != nil {
			return err
		}
	}

	// トークンの取得
	tokenPath := common.GetTokenPath(appDir)
	_, err = NewOAuthTokenGetter(true).GetAndSaveToken(config.CredentialPath, tokenPath, config.BrowserApp)
	if err != nil {
		return err
	}

	fmt.Println("インストールが完了しました。")
	fmt.Println("")

	if common.IsWindows() {
		fmt.Println("プログラムを動かすには %s をダブルクリックして起動してください", common.GetBinPath(appDir))
		fmt.Println("")
		powershellPath := path.Join(common.GetAppDir(), "register_startup.ps1")
		if PrintAndScanStdInput("デスクトップにショートカットを作りますか？ (y/n) > ") == "y" {

		}
	} else {
		if PrintAndScanStdInput("Macの常駐プロセスを起動しますか？ (y/n) > ") == "y" {
			daemonCtrl := NewDaemonCtrl()
			err := daemonCtrl.StartDaemon()
			if err != nil {
				return err
			}
			fmt.Println("常駐プロセスを起動しました")
			fmt.Println("2秒待ちます")
			time.Sleep(2 * time.Second)
			isRunning, err := daemonCtrl.IsDaemonRunning()
			if err != nil {
				return err
			}
			if !isRunning {
				return fmt.Errorf("常駐プロセスが起動していません")
			}
			fmt.Println("常駐プロセスが動いていることを確認しました")
			fmt.Println("常駐プロセスのログは以下のコマンドで確認できます")
			fmt.Printf("tail -f %s\n", common.GetLogPath(appDir))
		}
	}
	return nil
}
