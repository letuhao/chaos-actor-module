# Start Memurai Redis Server and Test Connection
Write-Host "Starting Memurai Redis Server..." -ForegroundColor Green
Write-Host "=================================" -ForegroundColor Green

# Check if Memurai is already running
$portCheck = netstat -an | Select-String ":6379"
if ($portCheck) {
    Write-Host "Memurai is already running on port 6379" -ForegroundColor Yellow
} else {
    Write-Host "Starting Memurai..." -ForegroundColor Cyan
    
    # Try to start Memurai (adjust path as needed)
    $memuraiPath = "C:\Program Files\Memurai\memurai-server.exe"
    if (Test-Path $memuraiPath) {
        Start-Process -FilePath $memuraiPath -ArgumentList "--port", "6379", "--maxmemory", "100mb", "--save", "60", "1" -WindowStyle Minimized
    } else {
        Write-Host "Memurai not found at $memuraiPath" -ForegroundColor Red
        Write-Host "Please install Memurai or update the path in this script" -ForegroundColor Red
        exit 1
    }
    
    # Wait for server to start
    Write-Host "Waiting for Memurai to start..." -ForegroundColor Cyan
    Start-Sleep -Seconds 5
}

# Test connection
Write-Host "Testing Redis connection..." -ForegroundColor Cyan
go run ./scripts/test_redis.go

Write-Host "Press any key to continue..." -ForegroundColor Yellow
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")
