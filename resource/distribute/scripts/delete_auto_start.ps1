Unregister-ScheduledTask -TaskName "GCAL-RUN" -Confirm:$false

if($? -eq "True"){
    Write-Host "Success to uninstall AUTO-START"
}else{
    Write-Host "Error!"
}
Write-Host ""
Read-Host "Press Enter to exit"