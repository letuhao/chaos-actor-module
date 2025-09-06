# API Contracts (TypeScript)

## Resolver
```ts
export interface IStatResolver {
  computeSnapshot(input: {
    actorId: string;
    level: number;
    baseAllocations: Record<StatKey, number>;
    registry: StatDef[];
    baseCurves: Record<string, unknown>;
    items: StatModifier[];
    titles: StatModifier[];
    passives: StatModifier[];
    buffs: StatModifier[];
    debuffs: StatModifier[];
    auras: StatModifier[];
    environment: StatModifier[];
    withBreakdown?: boolean;
  }): StatSnapshot;
}
```

## Snapshot Provider
```ts
export interface ISnapshotProvider {
  buildForActor(actorId: string, options?: { withBreakdown?: boolean }): Promise<StatSnapshot>;
}
```

## Progression Service
```ts
export interface IRpgProgressionService {
  grantXp(actorId: string, amount: number): Promise<{ newLevel?: number; pointsGranted?: number } | null>;
  allocatePoints(actorId: string, delta: Partial<Record<StatKey, number>>): Promise<void>;
  respec(actorId: string): Promise<void>;
}
```
