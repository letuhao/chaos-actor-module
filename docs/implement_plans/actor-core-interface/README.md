# Actor Core Interface Bundle (v1.0.0)

**Purpose:** Provide a language-agnostic, highest-level Actor Core interfaces to be used by the Combat system and other subsystems.
This bundle contains **interfaces only** (no implementations). It defines:
- Core data contracts (`PrimaryCore`, `Derived`, `CoreContribution`, `ActorResources`)
- Composition function signatures (`ComposeCore`, `BaseFromPrimary`, `FinalizeDerived`, `ClampDerived`)
- Deterministic merge rules (stable lexicographic by bucket key)
- Versioning fields and invariants

See `SPEC.md` for the full contract and rules.

Generated: 2025-09-06T20:25:40.054841Z
