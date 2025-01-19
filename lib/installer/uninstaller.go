package installer

import (
	"bufio"
	"fmt"
	"os"
)

type Uninstaller struct{}

func NewUninstaller() *Uninstaller {
	return &Uninstaller{}
}
func (u *Uninstaller) Uninstall(installDir string) {
	var err error
	fmt.Printf("ツールは %s にインストールされています\n", installDir)
	fmt.Printf("アンインストールしますか(y/n) >  ")
	scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
	scanner.Scan()
	yOrN := scanner.Text()
	if yOrN == "y" {
		// デーモンの停止
		err = NewDaemonCtrl().StopDaemon()
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		// 常駐プロセスファイルの削除
		err = NewDaemonCtrl().DeletePListFile()
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
