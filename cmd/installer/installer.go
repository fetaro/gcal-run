package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/alecthomas/kingpin/v2"
	"github.com/fetaro/gcal_forcerun_go/lib"
)

var (
	credential = kingpin.Arg("credential", "GoogleAPIのクレデンシャルファイル").Required().String()
)

func main() {
	kingpin.Parse()
	credentialPath, err := filepath.Abs(*credential)
	if err != nil {
		fmt.Printf("クレデンシャルファイルのフルパスの取得に失敗しました: %v\n", err)
		os.Exit(1)
	}
	_, err = os.Stat(credentialPath)
	if os.IsNotExist(err) {
		fmt.Println("クレデンシャルファイルを読み取れません")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ

	DefaultAppHome := path.Join(os.Getenv("HOME"), ".gcal_run")
	fmt.Printf("インストール先ディレクトリを指定してください\nデフォルトは「%s」です。デフォルトで良い場合は何も入力せずにEnterを押してください\n> ", DefaultAppHome)
	scanner.Scan()
	installDir := scanner.Text()
	if installDir == "" {
		installDir = DefaultAppHome
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
		fmt.Printf("ディレクトリが既に存在します: %s\n", installDir)
		os.Exit(1)
	}

	var browserApp string
	for {
		// 標準入力からsearchMinutesを取得
		fmt.Printf("ブラウザアプリケーションのパスを指定してください\nデフォルトは「%s」です。デフォルトで良い場合は何も入力せずにEnterを押してください\n> ", lib.DefaultBrowserApp)
		scanner.Scan()
		browserApp = scanner.Text()
		_, err := os.Stat(browserApp)
		if browserApp == "" {
			browserApp = lib.DefaultBrowserApp
			break
		}
		if !os.IsNotExist(err) {
			fmt.Println("ブラウザアプリケーションが存在しません。再度入力してください")
		} else {
			break
		}
	}

	var searchMinuteStr string
	var searchMinute int
	for {
		// 標準入力からsearchMinutesを取得
		fmt.Printf("会議の何分前に起動するか指定してください\nデフォルトは「%d分」です。デフォルトで良い場合は何も入力せずにEnterを押してください\n> ", lib.DefaultSearchMinutes)
		scanner.Scan()
		searchMinuteStr = scanner.Text()
		_, err := os.Stat(searchMinuteStr)
		if searchMinuteStr == "" {
			searchMinute = lib.DefaultSearchMinutes
			break
		}
		searchMinute, err = strconv.Atoi(searchMinuteStr)
		if err != nil {
			fmt.Println("数値を入力してください")
			continue
		} else {
			break
		}
	}

	installer := lib.NewInstaller()
	config := lib.NewConfig(credentialPath, installDir, searchMinute, browserApp)
	err = installer.Install(config)
	if err != nil {
		os.Exit(1)
	}
}
