Write-Host "Google�J�����_�[TV��c�����N���c�[����Windows�̃X�^�[�g�A�b�v�ɓo�^���܂�"
Write-Host "�����L�[�������ƁA�o�^�����s���܂�..."
Read-Host

$WsShell = New-Object -ComObject WScript.Shell
$appdata = $Env:APPDATA
$ShortcutPath = "${appdata}\Microsoft\Windows\Start Menu\Programs\Startup\gcal_run.lnk"
$Shortcut = $WsShell.CreateShortcut($ShortcutPath)
$Shortcut.TargetPath = "${appdata}\gcal_run\gcal_run.exe"
$Shortcut.IconLocation = "${appdata}\gcal_run\gcal_run.ico"
$Shortcut.Save()
Write-Host "�V���[�g�J�b�g�t�@�C�����쐬���܂��� $ShortcutPath"
Write-Host ""
Write-Host "����I��"
Write-Host ""
Write-Host "�����L�[�������Ƃ��̃E�C���h�E����܂�"
Read-Host