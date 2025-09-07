# CI Script for Actor Core
# Runs: build, lint, tests, race, bench

param(
    [switch]$SkipBuild,
    [switch]$SkipLint,
    [switch]$SkipTests,
    [switch]$SkipRace,
    [switch]$SkipBench,
    [switch]$Verbose
)

$ErrorActionPreference = "Stop"

Write-Host "🚀 Starting CI Pipeline for Actor Core" -ForegroundColor Green

# Check if golangci-lint is installed
if (-not $SkipLint) {
    try {
        golangci-lint --version | Out-Null
        Write-Host "✅ golangci-lint found" -ForegroundColor Green
    } catch {
        Write-Host "❌ golangci-lint not found. Installing..." -ForegroundColor Yellow
        go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    }
}

# Check if gotestsum is installed
if (-not $SkipTests) {
    try {
        gotestsum --version | Out-Null
        Write-Host "✅ gotestsum found" -ForegroundColor Green
    } catch {
        Write-Host "❌ gotestsum not found. Installing..." -ForegroundColor Yellow
        go install gotest.tools/gotestsum@latest
    }
}

# Check if benchstat is installed
if (-not $SkipBench) {
    try {
        benchstat -h | Out-Null
        Write-Host "✅ benchstat found" -ForegroundColor Green
    } catch {
        Write-Host "❌ benchstat not found. Installing..." -ForegroundColor Yellow
        go install golang.org/x/perf/cmd/benchstat@latest
    }
}

# Build
if (-not $SkipBuild) {
    Write-Host "`n🔨 Building..." -ForegroundColor Cyan
    go build ./...
    if ($LASTEXITCODE -ne 0) {
        Write-Host "❌ Build failed" -ForegroundColor Red
        exit 1
    }
    Write-Host "✅ Build successful" -ForegroundColor Green
}

# Lint
if (-not $SkipLint) {
    Write-Host "`n🔍 Linting..." -ForegroundColor Cyan
    golangci-lint run
    if ($LASTEXITCODE -ne 0) {
        Write-Host "❌ Linting failed" -ForegroundColor Red
        exit 1
    }
    Write-Host "✅ Linting passed" -ForegroundColor Green
}

# Tests
if (-not $SkipTests) {
    Write-Host "`n🧪 Running tests..." -ForegroundColor Cyan
    if ($Verbose) {
        gotestsum --format testname -- -race ./tests/...
    } else {
        gotestsum --format short-verbose -- -race ./tests/...
    }
    if ($LASTEXITCODE -ne 0) {
        Write-Host "❌ Tests failed" -ForegroundColor Red
        exit 1
    }
    Write-Host "✅ Tests passed" -ForegroundColor Green
}

# Race detection
if (-not $SkipRace) {
    Write-Host "`n🏃 Running race detection..." -ForegroundColor Cyan
    go test -race ./tests/...
    if ($LASTEXITCODE -ne 0) {
        Write-Host "❌ Race detection failed" -ForegroundColor Red
        exit 1
    }
    Write-Host "✅ No race conditions detected" -ForegroundColor Green
}

# Benchmarks
if (-not $SkipBench) {
    Write-Host "`n⚡ Running benchmarks..." -ForegroundColor Cyan
    $benchFile = "benchmark_results.txt"
    go test -bench=. -benchmem ./tests/... > $benchFile
    if ($LASTEXITCODE -ne 0) {
        Write-Host "❌ Benchmarks failed" -ForegroundColor Red
        exit 1
    }
    
    Write-Host "📊 Benchmark results:" -ForegroundColor Yellow
    Get-Content $benchFile | Select-Object -First 20
    
    Write-Host "`n📈 Benchmark analysis:" -ForegroundColor Yellow
    benchstat $benchFile
    
    Write-Host "✅ Benchmarks completed" -ForegroundColor Green
}

Write-Host "`n🎉 CI Pipeline completed successfully!" -ForegroundColor Green
