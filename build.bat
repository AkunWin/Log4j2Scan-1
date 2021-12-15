@echo off
echo start build
if exist %cd%\Log4j2Scan-Linux (
    del Log4j2Scan-Linux
)
if exist %cd%\Log4j2Scan-Darwin% (
    del Log4j2Scan-Darwin
)
if exist %cd%\Log4j2Scan-Windows.exe (
    del Log4j2Scan-Windows.exe
)

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o Log4j2Scan-Linux
echo start linux success

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o Log4j2Scan-Darwin
echo start darwin success

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o Log4j2Scan-Windows.exe
echo start windows success
