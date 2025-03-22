Unregister-ScheduledTask -TaskName "GCAL-TUN-TEST" -Confirm:$false

if($? -eq "True"){
    Write-Host "Success to delete AUTO START"
}else{
    Write-Host "Fail to delete up AUTO START"
}
Write-Host ""
Read-Host "Press Enter to exit"