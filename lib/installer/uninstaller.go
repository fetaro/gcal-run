package installer

import (
	"fmt"
	"os"
)

type Uninstaller struct{}

func NewUninstaller() *Uninstaller {
	return &Uninstaller{}
}
func (u *Uninstaller) Uninstall(installDir string) {
	fmt.Printf("ツールは %s にインストールされています\n", installDir)
	if PrintAndScanStdInput("アンインストールしますか(y/n) >  ") == "y" {
		// 常駐プロセスの停止
		err := NewDaemonCtrl().StopDaemon()
		if err != nil {
			panic(err)
		}
		// 常駐プロセスファイルの削除
		err = NewDaemonCtrl().DeletePListFile()
		if err != nil {
			panic(err)
		}
		// ファイルの削除
		fmt.Sprintln("インストールディレクトリの削除")
		err = os.RemoveAll(installDir)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ディレクトリを削除しました: %s\n", installDir)
		fmt.Println("アンインストールが完了しました")
	}
}
