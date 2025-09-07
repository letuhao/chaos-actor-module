# Cursor AI Execution Guide — Actor Core v3 (Aggregator-only)
**Updated:** 2025-09-07 17:55

This guide instructs Cursor AI (and human devs) to implement Actor Core v3 **step-by-step**, using the **merged design pack**:
`actor-core-v3-english-design-pro-merged.zip`.

---

## 0) Prerequisites
- Read these docs in order (from the merged pack):
  1. `01_Executive_Summary.md`
  2. `02_Architecture_and_Source_Structure.md`
  3. `04_Constants_Enums_Interfaces.md`
  4. `05_Data_Schemas.md` (skim schemas list)
  5. `06_Aggregation_Algorithm.md`
  6. `07_Caps_and_Layers.md`
  7. `08_Combiner_Registry.md`
  8. `10_Implementation_Guide.md`
  9. `14_Testing_Strategy.md`
- Keep nearby:
  - `appendix/Registry.example.yml`
  - `appendix/CapLayerRegistry.example.yml`
  - `E1_TestVectors/*`

---

## 1) Repository layout (create now)
```
/docs/                       # copy the merged design pack here (optional for devs)
/src/core/                   # interfaces + types only
/src/caps/                   # caps engine (within-layer + across-layer)
/src/registry/               # combiner & layer registry loaders
/tests/golden/               # JSON vectors -> expected snapshots
/tests/property/             # order/clamp invariants
```

---

## 2) Implement Core Interfaces & Types
Create in `/src/core/`:
- `Actor` (metadata only): GUID, Name, Race, LifeSpan, Age, CreatedAt, UpdatedAt, Version, `Subsystems[]`.
- `Subsystem` interface: `systemId()`, `priority()`, `contribute(actor, ctx) -> SubsystemOutput`.
- `Contribution`: `{ dimension, bucket: FLAT|MULT|POST_ADD|OVERRIDE, value, system, priority? }`.
- `CapContribution`: `{ system, dimension, mode: BASELINE|ADDITIVE|HARD_MAX|HARD_MIN|OVERRIDE, kind: max|min, value, priority?, scope?, realm?, tags? }`.
- `SubsystemOutput`: `primary[]`, `derived[]`, `caps[]`, optional `context`, `meta.system`.
- `MergeRule`: `{ usePipeline, operator?, clampDefault{min,max} }`.
- `CombinerRegistry` interface: `ruleFor(dimension)`.
- `CapLayerRegistry` data: `{ order[], across_policy }`.
- `CapsProvider` interface: `effectiveCaps(actor, outputs[], layer) -> EffectiveCaps` for within-layer; plus helper to reduce across layers.
- `Aggregator` interface: `resolve(actor) -> Snapshot`.
- `Snapshot`: `{ Primary, Derived, CapsUsed, Version }`.

> Conform exactly to **Section 04** in the merged design.

---

## 3) Implement Registries
### 3.1 Combiner Registry (`/src/registry/combiner.*`)
- Load YAML/JSON -> map `dimension -> MergeRule`.
- Default if missing: `usePipeline=true`, `clampDefault: {min:0, max:1e12}`.

### 3.2 Cap Layer Registry (`/src/registry/layers.*`)
- Load YAML/JSON matching `CapLayerRegistry.schema.json`.
- Default (if not provided):
  ```yaml
  order: [REALM, WORLD, EVENT, TOTAL]
  across_policy: INTERSECT
  ```

---

## 4) Implement Caps Provider (Layered)
### 4.1 Within-Layer Merge
For each layer L and each dimension d:
- Accumulators: baseline/additive/hard/override (max & min sides).
- Candidate: `override ? override : baseline + additive`.
- Final per layer: clip by `hard` and registry clampDefault -> `LayerCaps[L][d] = [Min,Max]`.

### 4.2 Across-Layer Reduction
- Let `order = CapLayerRegistry.order`.
- Start range = [-∞, +∞]; for each L in order:
  - If `LayerCaps[L][d]` exists, intersect: `Min=max(Min, L.Min)`, `Max=min(Max, L.Max)`.
- Intersect once more with registry clampDefault.
- Output `EffectiveCapsFinal[d]`.

> Determinism: sort cap contributions as `(dimension, mode, -priority, system)` before reducing.

---

## 5) Implement Aggregator
For each of `Primary` and `Derived`:
1) Gather all contributions grouped by `dimension`.
2) Sort `(bucket asc, priority desc, system asc)`.
3) Get `MergeRule` for the dimension.
4) If `usePipeline` -> compute `(sumFlat * product(1+MULT) + post)`; if any `OVERRIDE`, use its value.
5) Else apply `SUM | MAX | MIN`.
6) Clamp with `EffectiveCapsFinal[dimension]` (from Caps Provider).
7) Write into `Snapshot`.

> Base value is zero unless contributed; systems that don't train a stat won't change it.

---

## 6) Validation & Tests
### 6.1 JSON Schema Validation
- Validate `SubsystemOutput`, `CapContribution`, and registries against provided schemas.

### 6.2 Golden Tests (`/tests/golden/`)
- Load vectors from the merged pack:
  - `E1_TestVectors/case1_subsystems.json` -> expect `E1_TestVectors/case1_expected.json`
  - `E1_TestVectors/case_world_caps_subsystems.json` -> expect `E1_TestVectors/case_world_caps_expected.json`

### 6.3 Property Tests
- Shuffle contributions/subsystems order -> Snapshot invariant.
- Clamp invariants: `min ≤ value ≤ max` always.

### 6.4 Parity (optional)
- If you implement in a second language (e.g., TS), run the same vectors and assert equality.

---

## 7) Subsystems (later phases)
- Effects, Items, Race, Talent, Fate, Karma, Cultivation -> do not add formulas in the core.
- Each subsystem returns `SubsystemOutput`. Emit caps with proper `scope` (REALM/WORLD/EVENT/TOTAL) only for active layers.

---

## 8) PR Template (summary)
- Interfaces match design
- Registries load & default correctly
- Caps provider merges within-layer and across-layer (order respected)
- Aggregator pipeline/operator + deterministic sort
- JSON schema validation added
- Golden + property tests passing
- Docs: list new dimensions, update registries

---

## 9) Performance Optimizations (Critical for Games)
- **Read [23_Performance_Optimizations.md](23_Performance_Optimizations.md)** before implementation
- Implement caching from the start (L1 cache for snapshots)
- Add dirty tracking for incremental updates
- Use object pooling to reduce GC pressure
- Implement parallel processing for subsystems
- Add performance monitoring and benchmarks

---

## 10) Subsystem Development
- **Read [24_Subsystem_Development_Guide.md](24_Subsystem_Development_Guide.md)** for comprehensive guidance
- Implement core `Subsystem` interface with proper error handling
- Add optional interfaces (Configurable, Validating, Caching, Lifecycle)
- Use configuration management for subsystem settings
- Write comprehensive unit and integration tests
- Follow performance best practices from the start

---

## 11) Production Deployment
- **Read [25_Production_Deployment_Guide.md](25_Production_Deployment_Guide.md)** before going live
- Set up monitoring and observability (Prometheus, Jaeger)
- Configure caching layer (Redis) and database (PostgreSQL)
- Implement security measures (authentication, encryption)
- Set up backup and disaster recovery procedures
- Use blue-green or rolling deployment strategies

---

## 12) Common Pitfalls (avoid)
- Putting formulas in the core.
- Using factor for MULT instead of delta (must be k-1).
- Skipping sort -> non-determinism.
- Ignoring registry clampDefault.
- Emitting inactive realm/world caps.
- **Not implementing performance optimizations from the start**

---

## 13) Done Criteria
- All golden/property tests pass.
- Deterministic across runs (order shuffled).
- Snapshot respects layered caps (realm/world/event/total) exactly.
- No core math tied to any domain formula.
- **Performance benchmarks meet requirements (see 23_Performance_Optimizations.md)**
