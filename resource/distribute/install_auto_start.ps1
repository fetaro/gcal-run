Write-Host "Select operation"
Write-Host "1) Install Auto-Start"
Write-Host "2) Delete Auto-Start"
Write-Host ""
$operation = Read-Host "Enter operation number"

if($operation -eq "1"){
    $script_dir = Split-Path -Parent $MyInvocation.MyCommand.Definition
    $argument_list = "-ExecutionPolicy Unrestricted -File ${script_dir}\scripts\set_up_auto_start.ps1"
    Start-Process powershell -ArgumentList $argument_list -Verb runas -Wait
}elseif($operation -eq "2"){
    $script_dir = Split-Path -Parent $MyInvocation.MyCommand.Definition
    $argument_list = "-ExecutionPolicy Unrestricted -File ${script_dir}\scripts\delete_auto_start.ps1"
    Start-Process powershell -ArgumentList $argument_list -Verb runas -Wait
}else{
    Write-Host "Invalid operation"
    Read-Host "Press Enter to exit"
}