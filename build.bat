set version=v0.0.2
set current_dir=%~dp0
echo %current_dir%

rem copy icon file
copy .\resorce\gcal_run.syso .\cmd\gcal_run\

rem build
cd .\cmd\gcal_run\
go build  -ldflags "-X main.version=%version%" -o "%current_dir%\dist\gcal-run_windows_amd64_%version%\gcal_run.exe"
cd %current_dir%

rem remove icon file
del .\cmd\gcal_run\gcal_run.syso

cd .\cmd\installer\
go build  -ldflags "-X main.version=%version%" -o "%current_dir%\dist\gcal-run_windows_amd64_%version%\installer.exe"
cd %current_dir%

