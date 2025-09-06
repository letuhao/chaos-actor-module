# RPG Stats Sub-System — Bundle
Generated: 2025-09-06T21:04:44.659661Z

This bundle contains **design docs** and **MongoDB schemas** for a modular RPG Stats Sub-System
that plugs into your **Core Actor Interface**. The core remains DB-agnostic; this subsystem
handles persistence (MongoDB), progression, allocations, item/title/buff/debuff modifiers,
and emits a deterministic **StatSnapshot** for combat & UI.

## Contents
- `00-COLLECTION-CURSOR.md` — One-file control script to drive Cursor: read docs, create code, run tests.
- `01-BASIC-DESIGN.md` — Goals, boundaries, architecture, flows.
- `02-DETAIL-DESIGN.md` — Interfaces, resolver pipeline, stacking rules, extension points.
- `03-DATABASE-DESIGN-MONGODB.md` — Collections, indexes, schemas, example docs.
- `04-IMPLEMENT-GUIDE.md` — Step-by-step file-by-file implementation plan.
- `05-TEST-GUIDE.md` — Unit + integration test plan and sample cases.
- `06-API-CONTRACTS.md` — Suggested service contracts (TypeScript) for progression & snapshot building.
- `07-STACKING-RULES.md` — Deterministic stacking order and examples.
- `08-STAT-REGISTRY-EXAMPLE.md` — A starter registry with example formulas.
- `db/` — JSON examples and index scripts for MongoDB.

---

## Quick Start
1) Open `00-COLLECTION-CURSOR.md` in Cursor.
2) Follow the sections to let Cursor **read docs → scaffold code → implement → test**.
3) Integrate by calling the subsystem's **SnapshotProvider** to get a `StatSnapshot` and feed it into your Core.


> **Update 2025-09-06T21:15:08.141940Z**: Primary stats expanded to **8** (Morrowind-style): STR, INT, WIL, AGI, SPD, END, PER, LUK.
