$WsShell = New-Object -ComObject WScript.Shell
$appdata = $Env:APPDATA
$ShortcutPath = "${appdata}\Microsoft\Windows\Start Menu\Programs\Startup\gcal_run.lnk"
$Shortcut = $WsShell.CreateShortcut($ShortcutPath)
$Shortcut.TargetPath = "${appdata}\gcal_run\gcal_run.exe"
$Shortcut.IconLocation = "${appdata}\gcal_run\gcal_run.ico"
$Shortcut.Save()