$script_dir = Split-Path -Parent $MyInvocation.MyCommand.Definition
$argument_list = "-ExecutionPolicy Unrestricted -File ${script_dir}\scripts\set_up_auto_start.ps1"
Start-Process powershell -ArgumentList $argument_list -Verb runas -Wait
