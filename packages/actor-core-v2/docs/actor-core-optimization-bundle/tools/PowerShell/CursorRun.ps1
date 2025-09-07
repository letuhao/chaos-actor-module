# PowerShell helper for Windows developers
param(
    [switch]$Race=$true
)
Write-Host "== Build =="
go build ./...
if ($LASTEXITCODE -ne 0) { exit 1 }

Write-Host "== Lint =="
golangci-lint run ./...

Write-Host "== Test =="
if ($Race) { go test ./... -race -count=1 } else { go test ./... -count=1 }

Write-Host "== Bench =="
go test ./services/core -bench=. -benchmem -run ^$
