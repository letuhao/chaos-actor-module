# Docker Redis Test Script
Write-Host "Starting Redis with Docker..." -ForegroundColor Green
Write-Host "=============================" -ForegroundColor Green

# Check if Docker is running
try {
    $dockerVersion = docker version 2>$null
    if ($LASTEXITCODE -ne 0) {
        throw "Docker is not running"
    }
} catch {
    Write-Host "Docker is not running. Please start Docker Desktop first." -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

# Start Redis container
Write-Host "Starting Redis container..." -ForegroundColor Cyan
docker-compose up -d redis

# Wait for Redis to be ready
Write-Host "Waiting for Redis to be ready..." -ForegroundColor Cyan
Start-Sleep -Seconds 10

# Check if Redis is healthy
$healthCheck = docker inspect actor-core-redis --format='{{.State.Health.Status}}' 2>$null
if ($healthCheck -eq "healthy") {
    Write-Host "✅ Redis is healthy!" -ForegroundColor Green
} else {
    Write-Host "⚠️ Redis health check: $healthCheck" -ForegroundColor Yellow
}

# Test connection
Write-Host "Testing Redis connection..." -ForegroundColor Cyan
go run ./scripts/simple_redis.go

# Show Redis info
Write-Host "`nRedis container info:" -ForegroundColor Cyan
docker ps | Select-String "redis"

Write-Host "`nRedis logs:" -ForegroundColor Cyan
docker logs actor-core-redis --tail 10

Write-Host "`nPress any key to stop Redis container..." -ForegroundColor Yellow
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")

# Stop Redis container
Write-Host "Stopping Redis container..." -ForegroundColor Cyan
docker-compose down

Write-Host "Done!" -ForegroundColor Green
