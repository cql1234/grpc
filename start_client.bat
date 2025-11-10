@echo off
REM 启动客户端的批处理脚本

echo ========================================
echo 启动 gRPC 客户端...
echo ========================================
echo.
echo 请确保服务器已经在另一个终端启动！
echo.

cd /d %~dp0
go run client/main.go

echo.
pause
@echo off
REM 启动服务器的批处理脚本

echo ========================================
echo 启动 gRPC 服务器...
echo ========================================
echo.

cd /d %~dp0
go run server/main.go

