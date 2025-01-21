package installer

import (
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func (d *DaemonCtrl) GeneratePlistStr() string {
	logPath := common.GetLogPath(common.GetAppDir())
	binPath := common.GetBinPath(common.GetAppDir())
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
	</array>

	<key>StandardErrorPath</key>
	<string>%s</string>

	<key>StandardOutPath</key>
	<string>%s</string>
</dict>
</plist>
`, d.GetDaemonName(), binPath, logPath, logPath)
}

func (d *DaemonCtrl) CreatePListFile(confirmOverwite bool) error {
	if common.FileExists(d.GetPListPath()) {
		fmt.Printf("常駐プロセス(LaunchAgents)ファイルが既に存在します: %s", d.GetPListPath())
		if confirmOverwite && PrintAndScanStdInput("上書きしますか？ (y/n) > ") != "y" {
			fmt.Println("常駐プロセス(LaunchAgents)ファイルの上書きを中止しました")
			return nil
		}
	}
	err := os.WriteFile(d.GetPListPath(), []byte(d.GeneratePlistStr()), 0644)
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
	stdout, err := exec.Command("launchctl", "list").Output()
	if err != nil {
		return false, fmt.Errorf("launchctl list のコマンドの実行に失敗しました: %v", err)
	}
	stdoutStr := string(stdout)
	// \nで分割して、行ごとに処理
	lines := strings.Split(stdoutStr, "\n")
	a := d.GetDaemonName()
	fmt.Println(a)
	for _, line := range lines {
		if strings.Contains(line, d.GetDaemonName()) {
			// 正常の場合、lineは以下のようになっている
			// 3113    -       com.github.fetaro.gcal_run
			// 異常の場合、lineは以下のようになっている
			// -       78      com.github.fetaro.gcal_run
			pid := strings.Split(line, "\t")[0]
			errorStr := strings.Split(line, "\t")[1]
			if errorStr == "78" {
				fmt.Println("常駐プロセス(LaunchAgents)が起動していません。おそらくバイナリファイルに実行権限がありません。errorNo=", errorStr)
				return false, nil
			} else if pid == "-" {
				fmt.Println("常駐プロセス(LaunchAgents)が起動していません。errorNo=", errorStr)
				return false, nil
			} else {
				fmt.Printf("常駐プロセス(LaunchAgents)が起動しています。PID: %s\n", pid)
				return true, nil
			}
		}
	}
	fmt.Println("常駐プロセス(LaunchAgents)が登録されておらず、起動していません")
	return false, nil

}
