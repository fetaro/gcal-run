package main

import (
	"fmt"
	"github.com/alecthomas/kingpin/v2"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"github.com/fetaro/gcal_forcerun_go/lib/installer"
	"os"
)

var version string // ビルドスクリプトで埋め込む
var (
	app = kingpin.New("installer", "GoogleカレンダーTV会議強制起動ツールのインストラー")

	installCommand = app.Command("install", "インストール")
	credentialPath = installCommand.Flag("credential", "GoogleAPIのクレデンシャルファイル").Short('c').Required().ExistingFile()

	updateCommand = app.Command("update", "アップデート")

	uninstallCommand = app.Command("uninstall", "アンインストール")
)

func main() {
	app.Version(version)
	appDir := common.GetAppDir()
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case installCommand.FullCommand():
		inst := installer.NewInstaller()
		// ユーザから入力を受けて、設定を作る
		minutesAgo, browserApp := inst.ScanInput()
		config := common.NewConfig(*credentialPath, minutesAgo, browserApp)
		// インストールする
		inst.Install(config, appDir)
		if installer.PrintAndScanStdInput("常駐プロセスを起動しますか？ (y/n) > ") == "y" {
			err := installer.NewDaemonCtrl().StartDaemon()
			if err != nil {
				panic(err)
			}
			fmt.Println("常駐プロセスを起動しました")
		}

	case updateCommand.FullCommand():
		// binファイルが存在するかチェック
		_, err := os.Stat(appDir)
		if os.IsNotExist(err) {
			fmt.Printf("インストールしたディレクトリが見つかりません. 探したパス: %s\n", appDir)
			os.Exit(1)
		}
		installer.NewUpdator().Update(appDir)
	case uninstallCommand.FullCommand():
		// binファイルが存在するかチェック
		_, err := os.Stat(appDir)
		if os.IsNotExist(err) {
			fmt.Printf("インストールしたディレクトリが見つかりません. 探したパス: %s\n", appDir)
			os.Exit(1)
		}
		installer.NewUninstaller().Uninstall(appDir)
	}
}
