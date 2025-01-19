package main

import (
	"bufio"
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib"
	"os"
	"path/filepath"
)

func main() {
	// 引数の数をチェック
	var installDir string
	if len(os.Args) == 2 {
		installDir = os.Args[1]
	} else {
		installDir = lib.DefaultInstallDir()
	}
	binFilePath := filepath.Join(installDir, "gcal_run")
	// binファイルが存在するかチェック
	_, err := os.Stat(binFilePath)
	if os.IsNotExist(err) {
		fmt.Printf("インストールされているバイナリが見つかりません. 探したパス: %s\n", binFilePath)
		fmt.Println("インストールディレクトリをデフォルトから変更している場合は、第一引数にインストールディレクトリを指定してください")
		fmt.Println("使い方 : updator /path/to/install/dir")
		os.Exit(1)
	}
	fmt.Printf("アンインストールしますか(インストールディレクトリ: %s) y/n: \n", installDir)
	scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
	scanner.Scan()
	yOrN := scanner.Text()
	if yOrN == "y" {
		// デーモンの停止
		err = lib.NewDaemonCtrl().StopDaemon()
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		// 常駐プロセスファイルの削除
		err = lib.NewDaemonCtrl().DeletePListFile()
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		// ファイルの削除
		fmt.Sprintln("インストールディレクトリの削除")
		err = os.RemoveAll(installDir)
		if err != nil {
			fmt.Printf("ディレクトリを削除できませんでした: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("ディレクトリを削除しました: %s\n", installDir)
		fmt.Println("アンインストールが完了しました")
	}
}
