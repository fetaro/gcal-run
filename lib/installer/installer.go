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
func (i *Installer) MakeConfigFromUserInput() *common.Config {
	var err error

	var credPath string
	for {
		credPath = PrintAndScanStdInput("GoogleカレンダーAPIのクレデンシャルパスを指定してください > ")
		if !common.FileExists(credPath) {
			fmt.Println("ファイルが見つかりません。再度入力してください")
			continue
		}
		_, err := os.ReadFile(credPath)
		if err != nil {
			fmt.Printf("ファイルが読み込めません。再度入力してください: %v\n", err)
			continue
		}
		break
	}
	fmt.Println("")

	var browserApp string
	for {
		browserApp = PrintAndScanStdInput(fmt.Sprintf("ブラウザアプリケーションのパスを指定してください\nデフォルトはGoogle Chromeでパスは「%s」です。\nデフォルトで良い場合は何も入力せずにEnterを押してください\n> ", common.GetDefaultBrowserApp()))
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

func (i *Installer) InstallFiles(config *common.Config, appDir string) error {
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

	// ツールのコピー
	fmt.Printf("ツールをインストールディレクトリにコピーします. \".\" -> \"%s\"\n", appDir)
	CopyDir(".", appDir)

	if common.IsWindows() {
		err := NewWinShortcutMaker(appDir).MakeShortCut(common.GetWinDesktopShortcutPath())
		if err != nil {
			return err
		}
		fmt.Println("デスクトップにショートカットを作成しました")
	} else {
		// バイナリファイルに実行権限を付与
		for _, fileName := range []string{"gcal_run", "installer"} {
			filePath := filepath.Join(appDir, fileName)
			fmt.Printf("ファイル %s に実行権限を付与します\n", filePath)
			err := exec.Command("chmod", "+x", filePath).Run()
			if err != nil {
				return err
			}
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

	return nil
}

func (i *Installer) StartAtWindows(appDir string) error {
	if PrintAndScanStdInput("自動で起動されるように、スタートアップに登録しますか？ (y/n) > ") == "y" {
		err := NewWinShortcutMaker(appDir).MakeShortCut(common.GetWinStartupShortcutPath())
		if err != nil {
			return err
		}
	}

	fmt.Println("----------------------------------------------------------------------------------------------")
	fmt.Println("")
	fmt.Println("プログラムを手動で動かすにはデスクトップにある gcal_run のショートカットをダブルクリックして起動してください")
	fmt.Println("")
	fmt.Println("----------------------------------------------------------------------------------------------")
	fmt.Println("")

	return nil
}

func (i *Installer) StartAtMac(appDir string) error {
	var err error
	if PrintAndScanStdInput("自動で起動されるように、Macの常駐プロセスを登録して起動しますか？ (y/n) > ") == "y" {

		daemonCtrl := NewDaemonCtrl()
		// plistファイルを作成
		err = daemonCtrl.CreatePListFile(true)
		if err != nil {
			return err
		}

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
	return nil
}

func (i *Installer) Install(appDir string) error {
	config := i.MakeConfigFromUserInput()
	err := i.InstallFiles(config, appDir)
	if err != nil {
		return err
	}
	if common.IsWindows() {
		return i.StartAtWindows(appDir)
	} else {
		return i.StartAtMac(appDir)
	}
}
