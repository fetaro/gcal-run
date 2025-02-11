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
func (u *Uninstaller) Uninstall(installDir string, interactive bool) error {
	var err error
	if interactive && PrintAndScanStdInput("アンインストールしますか(y/n) >  ") == "y" {
		if common.IsWindows() {
			// ショートカットの削除
			err = os.Remove(common.GetWinDesktopShortcutPath())
			if err != nil {
				return fmt.Errorf("デスクトップショートカットの削除に失敗しました: %v\n", err)
			}
			fmt.Printf("デスクトップショートカットを削除しました: %s\n", common.GetWinDesktopShortcutPath())
			if common.FileExists(common.GetWinStartupShortcutPath()) {
				err = os.Remove(common.GetWinStartupShortcutPath())
				if err != nil {
					return fmt.Errorf("スタートアップの登録削除に失敗しましたが続行します: %v\n", err)
				}
				fmt.Printf("スタートアップの登録を削除しました: %s\n", common.GetWinStartupShortcutPath())
			}
		} else {
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
