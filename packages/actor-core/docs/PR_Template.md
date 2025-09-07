# PR: Actor Core v3 — Core Implementation

## Summary
- What was implemented? (interfaces, registries, caps provider, aggregator)

## Checklists
- [ ] Interfaces conform to design (Section 04)
- [ ] Combiner registry loader + defaults
- [ ] Cap layer registry loader + defaults
- [ ] Caps provider: within-layer merge + across-layer reduction (order respected)
- [ ] Aggregator: deterministic pipeline/operator + clamp
- [ ] JSON schema validation
- [ ] Golden tests (case1, world+realm) pass
- [ ] Property tests (shuffle, clamp invariants)
- [ ] Docs updated (dimension catalog, registries)

## Notes for Reviewers
- Determinism guarantees shown (sorting code)
- Any deviations from spec? Why?
