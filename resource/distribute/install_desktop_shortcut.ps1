$WsShell = New-Object -ComObject WScript.Shell
$appdata = $Env:APPDATA
$homepath = $Env:HOMEPATH
$ShortcutPath = "${homepath}\Desktop\gcal_run.lnk"
$Shortcut = $WsShell.CreateShortcut($ShortcutPath)
$Shortcut.TargetPath = "${appdata}\gcal_run\gcal_run.exe"
$Shortcut.IconLocation = "${appdata}\gcal_run\gcal_run.ico"
$Shortcut.Save()