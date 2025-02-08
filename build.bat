set version=v0.0.2
go build  -ldflags "-X main.version=%version%" -o "dist\gcal-run_windows_amd64_%version%\gcal_run.exe" .\cmd\gcal_run\gcal_run.go
go build  -ldflags "-X main.version=%version%" -o "dist\gcal-run_windows_amd64_%version%\installer.exe" .\cmd\installer\installer.go
