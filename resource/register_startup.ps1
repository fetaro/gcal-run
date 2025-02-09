Write-Host "GoogleカレンダーTV会議強制起動ツールをWindowsのスタートアップに登録します"
Write-Host "何かキーを押すと、登録を実行します..."
Read-Host

$WsShell = New-Object -ComObject WScript.Shell
$appdata = $Env:APPDATA
$ShortcutPath = "${appdata}\Microsoft\Windows\Start Menu\Programs\Startup\gcal_run.lnk"
$Shortcut = $WsShell.CreateShortcut($ShortcutPath)
$Shortcut.TargetPath = "${appdata}\gcal_run\gcal_run.exe"
$Shortcut.IconLocation = "${appdata}\gcal_run\gcal_run.ico"
$Shortcut.Save()
Write-Host "ショートカットファイルを作成しました $ShortcutPath"
Write-Host ""
Write-Host "正常終了"
Write-Host ""
Write-Host "何かキーを押すとこのウインドウを閉じます"
Read-Host