# 05 — XP & Leveling

## Sources
- **Kill** (`KILL`, `PVP_KILL`): XP dựa trên `targetLevel`, `rarity`, `elite/boss`.
- **Quest** (`QUEST`): XP từ nhiệm vụ; table dữ liệu.
- **Explore** (`EXPLORE`): XP khi khám phá mốc bản đồ.
- **Event** (`EVENT`): XP theo sự kiện thời gian thực.

## XP Formula (configurable)
1) **Exponential**: `XP_to_next(L) = XP_BASE * (XP_GROWTH ^ L)`
2) **Polynomial**: `XP_to_next(L) = XP_BASE * (L + 1) ^ XP_EXP`
3) **Hybrid**: poly đến L* rồi exponent (softcaps).

### Kill XP
```
XP = BaseByTier(targetTier) * DisparityModifier(level, targetLevel) * (1 + ΣdecoratorPerc) + ΣdecoratorFlat
XP_final = min(XP, capMax_from_decorators? else +∞)
```
- Anti‑exploit: DiminishingReturn(sameTargetCount), DailyCap, PartyShare, Rested bonus (qua decorators).

## Level Up
- Khi XP ≥ `XP_to_next(L)`: tăng L, trừ XP.
- **Points per level**: random `[MIN..MAX]` hoặc theo bảng; **player tự phân phối**.
- **Lifespan bonus**: mỗi level emit `CapContribution`: `{ dimension:'lifespan_years', mode:'ADDITIVE', kind:'max', value: LIFESPAN_PER_LEVEL, scope:'TOTAL' }`.
