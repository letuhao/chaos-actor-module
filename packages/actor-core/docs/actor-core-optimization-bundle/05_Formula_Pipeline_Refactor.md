# 05 — Formula Pipeline Refactor (Immutable)

## Goals
- Deterministic, **order-guaranteed** evaluation per stat.
- Pipeline per stat: **Flat → Mult → Clamp → Version++**.
- No in-place mutation; return **new** `DerivedStats` each resolve.

## Steps
1. Introduce `type Step struct {Apply(ds *DerivedStats, pc *PrimaryCore) (float64, error)}` with implementations: `FlatStep`, `MultStep`, `ClampStep`.
2. `Formula` becomes `struct { Name string; Deps []string; Steps []Step }`.
3. Implement **topological sort** across all formulas based on `Deps` to get `order[]`.
4. `ResolveDerivedStats`:
   - Clone previous snapshot (if any) or build fresh.
   - For each stat in `order`, run steps against inputs, **never touching inputs**.
   - After each stat finalize, set value in output.
5. Bump snapshot `Version++` once after all stats computed; set `UpdatedAt = now`.
6. Add property tests ensuring `Flat` then `Mult` is not commuted; and clamps apply last.
