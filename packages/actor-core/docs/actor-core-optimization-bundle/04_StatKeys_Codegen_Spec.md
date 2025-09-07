# 04 — Stat Keys Single-Source Schema & Codegen

We will keep JSON keys and internal keys in **snake_case** and generate both Go and TS enums/consts from one YAML.

## 1) Schema (edit `tools/StatSchema.example.yml`)
```yaml
version: 1
stats:
  derived:
    - key: hp_max
      display: Maximum Health
      clamp: {min: 1, max: 1000000}
      formula:
        flat: ["vitality * 10", "constitution * 5"]
        mult: ["1 + mastery_hp * 0.02"]
    - key: move_speed
      display: Movement Speed
      clamp: {min: 0, max: 100}
      formula:
        flat: ["agility * 0.1"]
        mult: []
  primary:
    - key: vitality
    - key: constitution
    - key: agility
```

## 2) Codegen
- **Go**: generate `constants/derived_stats.go` & `constants/primary_stats.go` (string consts = snake_case).
- **TypeScript**: generate `packages/shared/derived.ts` (string literals), and types.
- **Docs**: generate a markdown table for README.

## 3) Enforcement
- Add a test that **scans for string literals** matching stat keys and ensures they are **only imported from constants**.
