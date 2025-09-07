@echo off
echo Starting Memurai Redis Server...
echo ================================

REM Check if Memurai is already running
netstat -an | findstr :6379 >nul
if %errorlevel% == 0 (
    echo Memurai is already running on port 6379
    goto :test
)

REM Try to start Memurai
echo Starting Memurai...
memurai-server.exe --port 6379 --maxmemory 100mb --save 60 1

REM Wait a moment for server to start
timeout /t 3 /nobreak >nul

:test
echo Testing connection...
go run ./scripts/test_redis.go

pause
