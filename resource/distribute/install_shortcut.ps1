Param(
    [ValidateSet("desktop","startup")][String]$Arg1
)

$WsShell = New-Object -ComObject WScript.Shell
$appdata = $Env:APPDATA

if ($Arg1 -eq "desktop") {
    $homepath = $Env:HOMEPATH
    $ShortcutPath = "${homepath}\Desktop\gcal_run.lnk"
}else{ # startup
    $ShortcutPath = "${appdata}\Microsoft\Windows\Start Menu\Programs\Startup\gcal_run.lnk"
}

$Shortcut = $WsShell.CreateShortcut($ShortcutPath)
$Shortcut.TargetPath = "${appdata}\gcal_run\gcal_run.exe"
$Shortcut.IconLocation = "${appdata}\gcal_run\gcal_run.ico"
$Shortcut.Save()