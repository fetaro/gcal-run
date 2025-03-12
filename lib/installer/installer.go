package installer

import (
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Installer struct {
}

func NewInstaller() *Installer {
	return &Installer{}
}
func (i *Installer) MakeConfigFromUserInput() (*common.Config, error) {
	var err error
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	credPath, err := NewInputCredPath().Run(currentDir)
	if err != nil {
		return nil, err
	}
	minutesAgo, err := NewInputMinutes().Run()
	if err != nil {
		return nil, err
	}
	browserPath, err := NewBrowserPicker().Run()
	if err != nil {
		return nil, err
	}
	return common.NewConfig(credPath, minutesAgo, browserPath), nil
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
	config, err := i.MakeConfigFromUserInput()
	if err != nil {
		return err
	}
	err = i.InstallFiles(config, appDir)
	if err != nil {
		return err
	}
	if common.IsWindows() {
		return i.StartAtWindows(appDir)
	} else {
		return i.StartAtMac(appDir)
	}
}
