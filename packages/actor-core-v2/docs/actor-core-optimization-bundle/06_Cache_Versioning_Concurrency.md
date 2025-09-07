# 06 — Cache, Versioning & Concurrency

## Add to `StatResolver`
```go
type StatResolver struct {
    formulas map[string]Formula
    cache    map[string]float64
    mu       sync.RWMutex
    version  int64 // FormulaRegistryVersion
}
```
- **Key format:** `stat:<name>:v{version}` (optionally include hash of primary stats bucketized).
- Wrap `Get/Set/Clear` under locks.
- Increment `version` when **AddFormula/RemoveFormula** is called; also when constants/config changes.

## Dependency order
- Build DAG once and cache the order per `version`. Rebuild when `version` bumps.

## Race safety
- Run tests with `-race`. Add benchmarks for hot functions (ResolveStat/ResolveDerivedStats).
