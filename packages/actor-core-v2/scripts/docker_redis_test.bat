@echo off
echo Starting Redis with Docker...
echo =============================

REM Check if Docker is running
docker version >nul 2>&1
if %errorlevel% neq 0 (
    echo Docker is not running. Please start Docker Desktop first.
    pause
    exit /b 1
)

REM Start Redis container
echo Starting Redis container...
docker-compose up -d redis

REM Wait for Redis to be ready
echo Waiting for Redis to be ready...
timeout /t 10 /nobreak >nul

REM Test connection
echo Testing Redis connection...
go run ./scripts/simple_redis.go

REM Show Redis info
echo.
echo Redis container info:
docker ps | findstr redis

echo.
echo Redis logs:
docker logs actor-core-redis --tail 10

echo.
echo Press any key to stop Redis container...
pause

REM Stop Redis container
echo Stopping Redis container...
docker-compose down

echo Done!
