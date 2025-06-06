@echo off
echo Compiling pwd-tool...

rem Set CGO environment variable
set CGO_ENABLED=1

rem Build project
go build -o pwd-tool.exe

if %ERRORLEVEL% EQU 0 (
    echo Build successful!
    echo You can now run install.bat to install pwd-tool
) else (
    echo Build failed, please make sure gcc compiler (MinGW-w64) is installed
    echo You can download MinGW-w64 from:
    echo https://winlibs.com/ or https://www.mingw-w64.org/downloads/
)

pause