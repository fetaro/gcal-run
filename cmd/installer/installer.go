package main

import (
	"bufio"
	"fmt"
	"github.com/alecthomas/kingpin/v2"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"github.com/fetaro/gcal_forcerun_go/lib/installer"
	"os"
)

func exitWithError(err error) {
	fmt.Printf("エラー発生: %v\n", err.Error())
	fmt.Println("")
	fmt.Printf("終了するには何かキーを押してください... ")
	scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
	scanner.Scan()
	os.Exit(1)
}

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
	case "update":
		err = installer.NewUpdator().Update(appDir)
	case "uninstall":
		err = installer.NewUninstaller().Uninstall(appDir)
	default:
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
				err = installer.NewUninstaller().Uninstall(appDir)
			default:
				fmt.Println("無効なコマンドです。終了します")
			}
		}
		if err != nil {
			exitWithError(err)
		}
	}
}
