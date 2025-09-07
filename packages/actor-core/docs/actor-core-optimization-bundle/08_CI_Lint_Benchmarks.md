# 08 — CI, Lint & Benchmarks (Windows-friendly)

## Tools
- `golangci-lint` (enable `govet`, `ineffassign`, `gocritic`, `gosec`, `errcheck`, `dupl`).
- `gotestsum` for pretty output.
- `benchstat` to compare benches.

## Scripts
- `tools/PowerShell/CursorRun.ps1` runs: build, lint, tests, race, bench.
