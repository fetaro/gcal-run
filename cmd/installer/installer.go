package main

import (
	"bufio"
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"github.com/fetaro/gcal_forcerun_go/lib/installer"
	"os"
)

var version string // ビルドスクリプトで埋め込む

func main() {

	fmt.Println("------------------------------------------------")
	fmt.Println(common.ToolName + "インストラ―")
	fmt.Println("バージョン: " + version)
	fmt.Println("------------------------------------------------")

	var err error

	appDir := common.GetAppDir()
	if !common.FileExists(appDir) {
		fmt.Println("現状、" + common.ToolName + "はまだインストールされていません")
		fmt.Println("")
		yOrN := installer.PrintAndScanStdInput("ツールをインストールしますか？ (y/n) > ")
		fmt.Println("")
		if yOrN == "y" {
			i := installer.NewInstaller()
			config := i.ScanUserInput()
			err = i.Install(config, appDir)
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
		fmt.Printf("エラー発生: %v\n", err.Error())
	}
	fmt.Println("")
	fmt.Printf("終了するには何かキーを押してください... ")
	scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
	scanner.Scan()
}
