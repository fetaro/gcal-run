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
$triggerOnWorkstationLock.Delay = "PT30S"

# make trigger for "at logon"
$triggerAtLogon = New-ScheduledTaskTrigger -AtLogon
$triggerAtLogon.Delay = "PT30S"

$triggers = @($triggerOnWorkstationLock, $triggerAtLogon)

$action = New-ScheduledTaskAction -Execute "${env:APPDATA}\gcal_run\gcal_run.exe"

$userId = (Get-WmiObject -Query "SELECT * FROM Win32_UserAccount WHERE Name='$env:USERNAME'").SID
$principal = New-ScheduledTaskPrincipal -UserId $userId -LogonType Interactive -RunLevel Limited

$task = New-ScheduledTask -Action $action -Trigger $triggers -Principal $principal -Description "Google Calendar Auto Run Tool"

Register-ScheduledTask -InputObject $task -TaskName "GCAL-RUN"

if($? -eq "True"){
    Write-Host "Register ScheduledTask named GCAL-RUN"
    Write-Host "Success to install AUTO-START"
}else{
    Write-Host "Error!"
}
Write-Host ""
Read-Host "Press Enter to exit"