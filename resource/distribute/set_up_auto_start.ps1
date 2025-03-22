# make trigger for "on workstation lock"
$stateChangeTrigger = Get-CimClass `
    -Namespace ROOT\Microsoft\Windows\TaskScheduler `
    -ClassName MSFT_TaskSessionStateChangeTrigger

$triggerOnWorkstationLock = New-CimInstance `
    -CimClass $stateChangeTrigger `
    -Property @{
      StateChange = 8  # TASK_SESSION_STATE_CHANGE_TYPE.TASK_SESSION_UNLOCK (taskschd.h)
    } `
    -ClientOnly

# make trigger for "at logon"
$triggerAtLogon = New-ScheduledTaskTrigger -AtLogon

$triggers = @($triggerOnWorkstationLock, $triggerAtLogon)

$action = New-ScheduledTaskAction -Execute "${env:APPDATA}\gcal_run\gcal_run.exe"

$userId = (Get-WmiObject -Query "SELECT * FROM Win32_UserAccount WHERE Name='$env:USERNAME'").SID
$principal = New-ScheduledTaskPrincipal -UserId $userId -LogonType Interactive -RunLevel Limited

$task = New-ScheduledTask -Action $action -Trigger $triggers -Principal $principal

Register-ScheduledTask -InputObject $task -TaskName "GCAL-TUN-TEST"

Write-Host ""
Write-Host "Set up to AUTO START"
Write-Host ""
