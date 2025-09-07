// Actor Core v3 — Interfaces (skeleton)
import { SubsystemOutput, EffectiveCaps, MergeRule, Snapshot } from './types';

export interface Subsystem {
  systemId(): string;
  priority(): number;
  contribute(actor: Actor, ctx?: unknown): Promise<SubsystemOutput> | SubsystemOutput;
}

export interface CapLayerRegistry { order: string[]; across_policy: 'INTERSECT'|'PRIORITIZED_OVERRIDE' }
export interface CombinerRegistry { ruleFor(dimension:string): MergeRule }
export interface CapsProvider {
  // within-layer result
  effectiveCapsWithinLayer(actor: Actor, outputs: SubsystemOutput[], layer: string): EffectiveCaps;
  // across-layer reduction
  reduceAcrossLayers(layerCaps: Record<string, EffectiveCaps>, combiner: CombinerRegistry, layerReg: CapLayerRegistry): EffectiveCaps;
}

export interface Actor {
  GUID: string; Name: string; Race?: { id():string; name():string };
  LifeSpan?: number; Age?: number; CreatedAt?: Date; UpdatedAt?: Date; Version: number;
  Subsystems: Subsystem[];
}

export interface Aggregator {
  resolve(actor: Actor, combiner: CombinerRegistry, caps: CapsProvider, layers: CapLayerRegistry): Promise<Snapshot> | Snapshot;
}
