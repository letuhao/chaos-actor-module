# 11 — Known Issues & Fix Paths

- **Inconsistent stat keys** across JSON tags (`hp_max`), constants (`hpMax`), and resolver map (`"hp_max"`). → Normalize to **snake_case** and generate.
- **Placeholders** present in: models/effects/combat_effect.go, services/cache/redis_cache.go, services/monitoring/performance_monitor.go → Complete them first.
- **Cache not synchronized** (plain map). → Add `sync.RWMutex`, versioned keys, optional TTL.
- **PrimaryCore.New*** has ellipses (metadata init). → Fill timestamps, version, defaults.
- **PerformanceMonitor** lacks context cancellation in long-running collectors. → Add `ctx` and `select` on `Done()`.
