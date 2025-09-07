# 07 — Tests Checklist

## Golden tests (per stat)
- Input vector: generate grid for primary stats (e.g., vitality=1..100 step 10).
- Store expected outputs in `testdata/derived/<stat>.golden.json`.
- Diff on failures.

## Property-based tests
- Associativity/commutativity properties where applicable.
- Clamp invariants: always min ≤ x ≤ max.

## Race tests
- `go test ./... -race -run TestResolveDerived_Parallel`

## Cross-language parity
- Generate TS formulas from schema; run the same vectors and **assert byte-equal JSON**.
