package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fetaro/gcal_forcerun_go/lib"
)

func main() {
	// 引数の数をチェック
	if len(os.Args) != 2 {
		fmt.Println("第一引数にクレデンシャルファイルのパスを指定してください")
		fmt.Println("使い方 : installer /path/to/credential.json")
		os.Exit(1)
	}
	// 第一引数を取得
	credential := os.Args[1]
	credentialPath, err := filepath.Abs(credential)
	if err != nil {
		fmt.Printf("クレデンシャルファイルのフルパスの取得に失敗しました: %v\n", err)
		os.Exit(1)
	}
	_, err = os.Stat(credentialPath)
	if os.IsNotExist(err) {
		fmt.Println("クレデンシャルファイルを読み取れません")
		os.Exit(1)
	}
	fmt.Println("クレデンシャルファイルを読み取りました. ファイルパス: ", credentialPath)

	scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ

	fmt.Printf("インストール先ディレクトリを指定してください\nデフォルトは「%s」です。デフォルトで良い場合は何も入力せずにEnterを押してください\n> ", lib.DefaultInstallDir())
	scanner.Scan()
	installDir := scanner.Text()
	if installDir == "" {
		installDir = lib.DefaultInstallDir()
	}
	// installDirが存在しない場合は作る
	if _, err := os.Stat(installDir); os.IsNotExist(err) {
		err := os.MkdirAll(installDir, 0755)
		if err != nil {
			fmt.Printf("ディレクトリを作成できませんでした: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("ディレクトリを作成しました: %s\n", installDir)
	} else {
		fmt.Printf("ディレクトリが既に存在します。: %s\n", installDir)
		fmt.Printf("中身を空にして、インストールしますか？ (y/n) > ")
		scanner.Scan()
		yOrN := scanner.Text()
		if yOrN == "y" {
			// installDirの中身を空にする
			err := os.RemoveAll(installDir)
			if err != nil {
				fmt.Printf("ディレクトリを空にできませんでした: %v\n", err)
				os.Exit(1)
			} else {
				err := os.MkdirAll(installDir, 0755)
				if err != nil {
					fmt.Printf("ディレクトリを作成できませんでした: %v\n", err)
					os.Exit(1)
				}
				fmt.Printf("ディレクトリを空にして再作成しました: %s\n", installDir)
			}
		} else {
			fmt.Println("インストールを中止します")
			os.Exit(1)
		}
	}

	var browserApp string
	for {
		fmt.Printf("ブラウザアプリケーションのパスを指定してください\nデフォルトは「%s」です。デフォルトで良い場合は何も入力せずにEnterを押してください\n> ", lib.DefaultBrowserApp)
		scanner.Scan()
		browserApp = scanner.Text()
		_, err := os.Stat(browserApp)
		if browserApp == "" {
			browserApp = lib.DefaultBrowserApp
			break
		}
		if os.IsNotExist(err) {
			fmt.Println("ブラウザアプリケーションが存在しません。再度入力してください")
		} else {
			break
		}
	}

	var minutesAgoStr string
	var minutesAgo int
	for {
		fmt.Printf("会議の何分前に起動するか指定してください\nデフォルトは「%d分」です。デフォルトで良い場合は何も入力せずにEnterを押してください\n> ", lib.DefaultMinutesAgo)
		scanner.Scan()
		minutesAgoStr = scanner.Text()
		_, err := os.Stat(minutesAgoStr)
		if minutesAgoStr == "" {
			minutesAgo = lib.DefaultMinutesAgo
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

	installer := lib.NewInstaller()
	config := lib.NewConfig(credentialPath, installDir, minutesAgo, browserApp)
	err = installer.Install(config)
	if err != nil {
		fmt.Printf("インストールに失敗しました: %v\n", err)
		os.Exit(1)
	}
}
