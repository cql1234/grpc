@echo off
REM ???????????
echo ========================================
echo ?? gRPC ???...
echo ========================================
echo.
cd /d %~dp0
go run server/main.go
