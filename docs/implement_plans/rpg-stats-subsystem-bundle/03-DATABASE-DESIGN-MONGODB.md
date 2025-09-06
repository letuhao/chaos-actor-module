# Database Design — MongoDB

## Collections
### `player_progress`
- `_id`: `actorId` (string)
- `level`: number
- `xp`: number
- `unspentPoints`: number
- `allocations`: object of `{ [StatKey]: number }`
- `updatedAt`: ISO date

**Indexes**
- `{ _id: 1 }` (PK)

---

### `player_effects_active`
- `_id`: ObjectId
- `actorId`: string
- `sourceKind`: "ITEM" | "TITLE" | "BUFF" | "DEBUFF" | "PASSIVE" | "AURA" | "ENV"
- `sourceId`: string
- `modifiers`: `StatModifier[]`
- `expiresAt`: ISO date | null
- `stackId`: string | null
- `stacks`: number (default 1)
- `appliedAt`: ISO date

**Indexes**
- `{ actorId: 1 }`
- `{ actorId: 1, expiresAt: 1 }` (for cleanup jobs)

---

### `player_equipment`
- `_id`: ObjectId
- `actorId`: string
- `slot`: string ("HEAD","CHEST","MAIN_HAND","OFF_HAND",...)
- `itemId`: string
- `modifiers`: `StatModifier[]`
- `updatedAt`: ISO date

**Indexes**
- `{ actorId: 1 }`
- `{ actorId: 1, slot: 1 }` (unique)

---

### `titles_owned`
- `_id`: ObjectId
- `actorId`: string
- `titleId`: string
- `isActive`: boolean
- `modifiers`: `StatModifier[]`

**Indexes**
- `{ actorId: 1 }`
- `{ actorId: 1, isActive: 1 }`

---

### `content_stat_registry`
- `_id`: string (content build id)
- `version`: number
- `stats`: `StatDef[]`
- `baseCurves`: object (serialized curve defs)
- `createdAt`: ISO date

**Indexes**
- `{ _id: 1 }`

## Example Documents
See `db/examples/*.json` in this bundle.

## Index Script
See `db/indexes.js` for `createIndex` calls.

## Notes
- Use **lean** projections for hot paths.
- Consider a **cache layer** for content registry.
- TTL optional: add TTL index on `expiresAt` in `player_effects_active` for expired cleanup.
