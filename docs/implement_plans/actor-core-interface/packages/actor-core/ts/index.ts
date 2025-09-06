// packages/actor-core/ts/index.ts
// Interfaces only; no implementations. Keep functions pure.

export type PrimaryCore = {
  HPMax: number;
  LifeSpan: number;
  Attack: number;
  Defense: number;
  Speed: number;
};

export type Derived = {
  HPMax: number; MPMax: number;
  ATK: number; MAG: number;
  DEF: number; RES: number;
  Haste: number; CritChance: number; CritMulti: number;
  MoveSpeed: number; RegenHP: number; RegenMP: number;
  Resists: Record<string, number>;
  Amplifiers: Record<string, number>;
  Version: bigint; // bump on recompute
};

export type CoreContribution = {
  Primary?: Partial<PrimaryCore>;
  Flat?: Record<string, number>;
  Mult?: Record<string, number>;
  Tags?: string[];
};

export type ResourceState = { current: number; max: number; regen: number; epoch: bigint };
export type ActorResources = Record<string, ResourceState>;

export interface ActorCore {
  ComposeCore(buckets: Record<string, CoreContribution>): CoreContribution;
  BaseFromPrimary(p: PrimaryCore, level: number): Derived;
  FinalizeDerived(base: Derived, flat: Record<string, number>, mult: Record<string, number>): Derived;
  ClampDerived(d: Derived): Derived;
}

// NOTE: Implementations must:
// - merge bucket keys in lexicographic order
// - treat missing Mult entries as 1.0
// - apply flats then mults then clamp
// - bump Version (+1) on FinalizeDerived
