package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"github.com/fetaro/gcal_forcerun_go/lib/installer"
)

var version string // ビルドスクリプトで埋め込む
var (
	app       = kingpin.New("gcal_run", "gcal_run: "+common.ToolName)
	operation = app.Arg("operation", "[install|update|uninstall").String()
)

func main() {
	app.Version(version)
	app.Parse(os.Args[1:])
	fmt.Println(*operation)
	fmt.Println("------------------------------------------------")
	fmt.Println(common.ToolName + "インストラ―")
	fmt.Println("バージョン: " + version)
	fmt.Println("------------------------------------------------")

	var err error

	appDir := common.GetAppDir()
	switch *operation {
	case "install":
		err = installer.NewInstaller().Install(appDir)
		if err != nil {
			fmt.Printf("インストールに失敗しました: %v\n", err)
			fmt.Println("インストールに失敗したため、クリーンアップします")
			installer.NewUninstaller().Uninstall(appDir, false)
		}
	case "update":
		err = installer.NewUpdator().Update(appDir)
	case "uninstall":
		installer.NewUninstaller().Uninstall(appDir, true)
	case "install-auto-start":
		fmt.Println("自動起動登録のみを行います")
		err = installer.NewInstaller().InstallAutoStart(appDir)
	case "uninstall-auto-start":
		fmt.Println("自動起動登録の解除のみを行います")
		installer.NewUninstaller().UninstallAutoStart(appDir)
	case "":
		// インタラクティブモード
		if !common.FileExists(appDir) {
			fmt.Println("現状、" + common.ToolName + "はまだインストールされていません")
			fmt.Println("")
			yOrN := installer.PrintAndScanStdInput("ツールをインストールしますか？ (y/n) > ")
			fmt.Println("")
			if yOrN == "y" {
				err = installer.NewInstaller().Install(appDir)
			} else {
				fmt.Println("インストールをキャンセルしました")
			}
			if err != nil {
				fmt.Printf("インストールに失敗しました: %v\n", err)
				fmt.Println("インストールに失敗したため、クリーンアップします")
				installer.NewUninstaller().Uninstall(appDir, false)
			}
		} else {
			fmt.Println(common.ToolName + "が既にインストールされています。インストールディレクトリ=" + appDir)
			commandNo := installer.PrintAndScanStdInput("実行できるコマンド\n " +
				"[1] バージョンアップ\n " +
				"[2] アンインストール\n " +
				"\n" +
				"実行したいコマンドの番号を指定してください > ")
			switch commandNo {
			case "1":
				err = installer.NewUpdator().Update(appDir)
			case "2":
				installer.NewUninstaller().Uninstall(appDir, true)
			default:
				fmt.Println("無効なコマンドです。終了します")
			}
		}
	default:
		fmt.Println("無効な引数です。引数は、install, update, uninstall, install-auto-start, uninstall-auto-start, [引数なし] のいずれかを指定してください")
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("エラー発生: %v\n", err.Error())
		fmt.Println("")
		fmt.Printf("終了するには何かキーを押してください... ")
		scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
		scanner.Scan()
		os.Exit(1)
	} else {
		fmt.Println("正常終了")
		fmt.Println("")
		fmt.Printf("終了するには何かキーを押してください... ")
		scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
		scanner.Scan()
		os.Exit(0)
	}
}
