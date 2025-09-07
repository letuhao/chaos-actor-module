# 01 — Optimization Plan & Priorities

## Snapshot
- Total Go LOC (incl. tests): **22105**
- Top heavy files:
- services/core/stat_resolver.go — 934 LOC
- tests/unit/effects/non_combat_effect_test.go — 905 LOC
- tests/unit/effects/effect_examples_test.go — 892 LOC
- tests/unit/config_manager_test.go — 873 LOC
- tests/unit/performance_monitor_test.go — 809 LOC
- tests/unit/flexible_stats_test.go — 800 LOC
- models/effects/combat_effect_examples.go — 793 LOC
- tests/unit/state_tracker_test.go — 707 LOC
- models/core/derived_stats.go — 702 LOC
- tests/unit/effects/combat_effect_manager_test.go — 697 LOC

## Critical priorities (do in order)
1. **Finish placeholders & stubs** so `go build ./...` and unit tests are green.
- models/effects/combat_effect.go
- services/cache/redis_cache.go
- services/monitoring/performance_monitor.go
2. **Unify stat keys** (JSON, constants, formula map): choose **snake_case** everywhere (JSON-friendly) and generate Go/TS constants via a schema.
3. **Refactor formula execution** into an **immutable pipeline** (Flat → Mult → Clamp) with **version bump** per finalize; enforce **determinism**.
4. **Concurrency-safe cache** in `StatResolver` (add `sync.RWMutex`, versioned cache keys, optional TTL); add **dependency DAG topological sort**.
5. **Finish PrimaryCore & DerivedStats** constructors and mutators; ensure **JSON tags**, **timestamps**, **version increments**, and **Clone** copies deeply.
6. **Observability**: expand `PerformanceMonitor` to be context-driven; avoid races; add Prometheus-style counters (behind interface).
7. **Cross-lang parity**: generate TypeScript enums/types & pure-formulas from the same schema; add parity tests.
8. **CI & linting**: `golangci-lint`, race detector, benches; Windows-safe scripts.

## Success criteria
- Reproducible output: Go and TS return identical numbers for all stats for the same inputs.
- Zero data races at `-race`; solid cache hit ratio; predictable allocation profile in hot paths.
- Golden tests ensure formula changes are intentional (diff-based review).
