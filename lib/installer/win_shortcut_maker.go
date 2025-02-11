package installer

import (
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"os"
	"os/exec"
	"path/filepath"
)

type WinShortcutMaker struct {
	appDir string
}

func NewWinShortcutMaker(appDir string) *WinShortcutMaker {
	return &WinShortcutMaker{appDir: appDir}
}

func (w *WinShortcutMaker) MakeShortCut(shortcutPath string) error {
	tmpDir, err := os.MkdirTemp("", "gcal_run_install_shortcut")
	if err != nil {
		return err
	}
	tmpFilePath := filepath.Join(tmpDir, "gcal_run_install_shortcut.ps1")
	powerShellStr := fmt.Sprintf(`
$ShortcutPath = "%s"
$WsShell = New-Object -ComObject WScript.Shell
$Shortcut = $WsShell.CreateShortcut($ShortcutPath)
$Shortcut.TargetPath = "%s"
$Shortcut.IconLocation = "%s"
$Shortcut.Save()
Write-Host "Success to make shortcut: $ShortcutPath"
`,
		shortcutPath,
		common.GetBinPath(w.appDir),
		common.GetWinIconPath(w.appDir))

	file, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}
	_, err = file.WriteString(powerShellStr)
	if err != nil {
		return err
	}
	file.Close()

	out, err := exec.Command("cmd", "/c", "powershell.exe", tmpFilePath).Output()
	if err != nil {
		return err
	}
	sjisStr, err := common.SJisToUtf8(string(out))
	if err != nil {
		fmt.Println("SJISへの変換に失敗しましたが、動作に影響はないので、無視します")
	}
	fmt.Println(sjisStr)
	if err != nil {
		return err
	}
	return nil
}
