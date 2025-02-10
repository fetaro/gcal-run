package installer

import (
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"os"
)

type Uninstaller struct{}

func NewUninstaller() *Uninstaller {
	return &Uninstaller{}
}
func (u *Uninstaller) Uninstall(installDir string) error {
	fmt.Printf("ツールは %s にインストールされています\n", installDir)
	var err error
	if PrintAndScanStdInput("アンインストールしますか(y/n) >  ") == "y" {
		if !common.IsWindows() {
			// 常駐プロセスの停止
			err = NewDaemonCtrl().StopDaemon()
			if err != nil {
				return err
			}
			// 常駐プロセスファイルの削除
			err = NewDaemonCtrl().DeletePListFile()
			if err != nil {
				fmt.Printf("ファイルの削除に失敗しましたが続行します: %v\n", err)
			}
		}
		// ファイルの削除
		fmt.Sprintln("インストールディレクトリの削除")
		err = os.RemoveAll(installDir)
		if err != nil {
			return err
		}
		fmt.Printf("ディレクトリを削除しました: %s\n", installDir)
		fmt.Println("アンインストールが完了しました")
	}
	return nil
}
