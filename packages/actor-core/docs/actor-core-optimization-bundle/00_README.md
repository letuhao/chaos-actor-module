# Actor Core — Optimization Bundle (Cursor-ready)
**Generated:** 2025-09-07 14:33

This bundle contains a detailed, step-by-step optimization plan and *per-file, per-method* task lists so Cursor can implement reliably on Windows.

## Files in this bundle
- `01_Optimization_Plan.md` — High-level plan & priorities.
- `02_Codebase_Walkthrough.md` — Map of modules & responsibilities.
- `03_File-by-File_Steps.md` — **Concrete steps for each file & method** (Cursor feed).
- `04_StatKeys_Codegen_Spec.md` — Single-source-of-truth schema + codegen plan (Go/TS).
- `05_Formula_Pipeline_Refactor.md` — Immutable pipeline (Flat → Mult → Clamp) with versioning.
- `06_Cache_Versioning_Concurrency.md` — Cache invalidation keys, locks, race-safety.
- `07_Tests_Checklist.md` — Golden/property/cross-lang/race tests.
- `08_CI_Lint_Benchmarks.md` — Windows-friendly scripts, golangci-lint, bench.
- `09_TS_Parity_Bridge.md` — TypeScript parity generation & parity tests.
- `10_Migration_Playbook.md` — Ordered migration (safe, shippable after each phase).
- `11_Known_Issues.md` — Incomplete areas & how to fix.
- `tools/StatSchema.example.yml` — Example schema for codegen (edit and extend).
- `tools/PowerShell/CursorRun.ps1` — Convenience script for Windows devs.

> **Usage order:** Start at `10_Migration_Playbook.md`, then open each referenced file.
