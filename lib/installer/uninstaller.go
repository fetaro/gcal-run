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

func (i *Uninstaller) UninstallAutoStart(appDir string) {
	var err error
	if common.IsWindows() {
		// ショートカットの削除
		err = os.Remove(common.GetWinDesktopShortcutPath())
		if err != nil {
			fmt.Printf("注意：デスクトップショートカットの削除に失敗しましたが続行します: %v\n", err)
		}else{
        	fmt.Printf("デスクトップショートカットを削除しました: %s\n", common.GetWinDesktopShortcutPath())
		}
		if common.FileExists(common.GetWinStartupShortcutPath()) {
			err = os.Remove(common.GetWinStartupShortcutPath())
			if err != nil {
				fmt.Printf("注意：スタートアップの登録削除に失敗しましたが続行します: %v\n", err)
			}else{
    			fmt.Printf("スタートアップの登録を削除しました: %s\n", common.GetWinStartupShortcutPath())
			}
		}
	} else {
		// 常駐プロセスの停止
		err = NewDaemonCtrl().StopDaemon()
		if err != nil {
			fmt.Printf("注意：常駐プロセスの停止に失敗しましたが続行します: %v\n", err)
		}else{
			fmt.Println("常駐プロセスを停止しました")
		}
		// 常駐プロセスファイルの削除
		err = NewDaemonCtrl().DeletePListFile()
		if err != nil {
			fmt.Printf("注意：常駐プロセスファイルの削除に失敗しましたが続行します: %v\n", err)
		}else{
			fmt.Println("常駐プロセスファイルを削除しました")
		}
	}
}	

func (u *Uninstaller) Uninstall(installDir string, interactive bool) {
	var err error
	if interactive && PrintAndScanStdInput("アンインストールしますか(y/n) >  ") == "y" {
		// 自動起動解除
		u.UninstallAutoStart(installDir)
		// ファイルの削除
		fmt.Sprintln("インストールディレクトリの削除")
		err = os.RemoveAll(installDir)
		if err != nil {
			fmt.Printf("注意：インストールディレクトリの削除に失敗しました: %v\n", err)
		}else{
			fmt.Printf("インストールディレクトリを削除しました: %s\n", installDir)
		}
	}
}
