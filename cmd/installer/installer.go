package main

import (
	"bufio"
	"fmt"
	"github.com/alecthomas/kingpin/v2"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"github.com/fetaro/gcal_forcerun_go/lib/installer"
	"os"
	"path/filepath"
)

var (
	app = kingpin.New("installer", "GoogleカレンダーTV会議強制起動ツールのインストラー")

	installCommand = app.Command("install", "インストール")
	credentialPath = installCommand.Flag("credential", "GoogleAPIのクレデンシャルファイル").Short('c').Required().ExistingFile()

	updateCommand    = app.Command("update", "アップデート")
	updateInstallDir = updateCommand.Flag("dir", "インストールディレクトリ").Default(common.DefaultInstallDir()).ExistingDir()

	uninstallCommand    = app.Command("uninstall", "アンインストール")
	uninstallInstallDir = uninstallCommand.Flag("dir", "インストールディレクトリ").Default(common.DefaultInstallDir()).ExistingDir()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case installCommand.FullCommand():
		inst := installer.NewInstaller()
		// ユーザから入力を受けて、設定を作る
		config, err := inst.ScanInput(*credentialPath)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		// インストールする
		err = inst.Install(config)
		if err != nil {
			fmt.Printf("インストールに失敗しました: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("常駐プロセスを起動しますか？ (y/n) > ")
		scanner2 := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
		scanner2.Scan()
		yOrN := scanner2.Text()
		if yOrN == "y" {
			err = installer.NewDaemonCtrl().StartDaemon()
			if err != nil {
				fmt.Printf("常駐プロセスの起動に失敗しました: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("常駐プロセスを起動しました")
		}

	case updateCommand.FullCommand():
		binFilePath := filepath.Join(*updateInstallDir, "gcal_run")
		// binファイルが存在するかチェック
		_, err := os.Stat(binFilePath)
		if os.IsNotExist(err) {
			fmt.Printf("インストールされているバイナリが見つかりません. 探したパス: %s\n", binFilePath)
			fmt.Println("インストールディレクトリをデフォルトから変更している場合は、引数にインストールディレクトリを指定してください")
			os.Exit(1)
		}
		installer.NewUpdator().Update(*updateInstallDir)
	case uninstallCommand.FullCommand():
		binFilePath := filepath.Join(*uninstallInstallDir, "gcal_run")
		// binファイルが存在するかチェック
		_, err := os.Stat(binFilePath)
		if os.IsNotExist(err) {
			fmt.Printf("インストールされているバイナリが見つかりません. 探したパス: %s\n", binFilePath)
			fmt.Println("インストールディレクトリをデフォルトから変更している場合は、引数にインストールディレクトリを指定してください")
			os.Exit(1)
		}
		installer.NewUninstaller().Uninstall(*uninstallInstallDir)
	}
}
