package installer

import (
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"os"
	"os/exec"
	"path/filepath"
)

type DaemonCtrl struct{}

func NewDaemonCtrl() *DaemonCtrl {
	return &DaemonCtrl{}
}

func (d *DaemonCtrl) GetDaemonName() string {
	if os.Getenv("GCAL_RUN_TEST") == "1" {
		return "com.github.fetaro.gcal_run_test"
	} else {
		return "com.github.fetaro.gcal_run"
	}
}
func (d *DaemonCtrl) GetPListPath() string {
	// ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist を返す
	return filepath.Join(os.Getenv("HOME"), fmt.Sprintf("Library/LaunchAgents/%s.plist", d.GetDaemonName()))
}

func (d *DaemonCtrl) GeneratePlistStr(c *common.Config) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>RunAtLoad</key>
	<false/>

	<key>KeepAlive</key>
	<true/>

	<key>Label</key>
	<string>%s</string>

	<key>ProgramArguments</key>
	<array>
		<string>%s</string>
        <string>--credential</string>
		<string>%s</string>
        <string>--dir</string>
		<string>%s</string>
        <string>--minute</string>
		<string>%d</string>
        <string>--browser</string>
		<string>%s</string>
	</array>

	<key>StandardErrorPath</key>
	<string>%s</string>

	<key>StandardOutPath</key>
	<string>%s</string>
</dict>
</plist>
`, d.GetDaemonName(), c.BinPath, c.CredentialPath, c.InstallDir, c.MinutesAgo, c.BrowserApp, c.LogPath, c.LogPath)
}

func (d *DaemonCtrl) CreatePListFile(c *common.Config) error {
	err := os.WriteFile(d.GetPListPath(), []byte(d.GeneratePlistStr(c)), 0644)
	if err != nil {
		return fmt.Errorf("常駐プロセス(LaunchAgents)ファイルの作成に失敗しました. エラー: %v", err)
	} else {
		fmt.Printf("常駐プロセス(LaunchAgents)ファイルを作成しました: %s\n", d.GetPListPath())
		return nil
	}
}
func (d *DaemonCtrl) DeletePListFile() error {
	err := os.Remove(d.GetPListPath())
	if err != nil {
		return fmt.Errorf("常駐プロセス(LaunchAgents)ファイルの削除に失敗しました. エラー: %v", err)
	} else {
		fmt.Printf("常駐プロセス(LaunchAgents)ファイルを削除しました: %s\n", d.GetPListPath())
		return nil
	}
}

func (d *DaemonCtrl) StartDaemon() error {
	fmt.Println("常駐プロセス(LaunchAgents)を開始します")
	launchctlCmd := exec.Command("launchctl", "load", d.GetPListPath())
	fmt.Println(launchctlCmd)
	err := launchctlCmd.Run()
	if err != nil {
		return fmt.Errorf("常駐プロセス(LaunchAgents)の開始に失敗しました. エラー: %v", err)
	}
	isRunning, err := d.IsDaemonRunning()
	if err != nil {
		return fmt.Errorf("常駐プロセス(LaunchAgents)の起動確認に失敗しました. エラー: %v", err)
	}
	if !isRunning {
		return fmt.Errorf("常駐プロセス(LaunchAgents)のコマンドは成功しましたが、プロセスは起動できませんでした")
	}
	fmt.Println("常駐プロセス(LaunchAgents)を開始しました")
	return nil

}

func (d *DaemonCtrl) StopDaemon() error {
	fmt.Println("常駐プロセス(LaunchAgents)を停止します")
	launchctlCmd := exec.Command("launchctl", "unload", d.GetPListPath())
	fmt.Println(launchctlCmd)
	err := launchctlCmd.Run()
	if err != nil {
		return fmt.Errorf("常駐プロセス(LaunchAgents)の停止に失敗しました. エラー: %v", err)
	} else {
		fmt.Println("常駐プロセス(LaunchAgents)を停止しました")
		return nil
	}
}

func (d *DaemonCtrl) IsDaemonRunning() (bool, error) {
	stdout, err := exec.Command("launchctl", "list", d.GetDaemonName()).Output()
	if err != nil {
		fmt.Println("常駐プロセス(LaunchAgents)が起動していません")
		return false, nil
	}
	//stdoutの一文字目が「-」であれば、デーモンが起動していない
	pid := stdout[0]
	if pid == '-' {
		fmt.Println("常駐プロセス(LaunchAgents)が起動していません")
		return false, nil
	} else {
		fmt.Printf("常駐プロセス(LaunchAgents)が起動しています。PID: %v\n", pid)
		return true, nil
	}
}
