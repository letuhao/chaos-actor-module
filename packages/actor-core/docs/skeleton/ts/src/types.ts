// Actor Core v3 — Types (skeleton)
export type Bucket = 'FLAT'|'MULT'|'POST_ADD'|'OVERRIDE';
export type CapMode = 'BASELINE'|'ADDITIVE'|'OVERRIDE'|'HARD_MAX'|'HARD_MIN';

export interface Contribution { dimension:string; bucket:Bucket; value:number; system:string; priority?:number }
export interface CapContribution { system:string; dimension:string; mode:CapMode; kind:'max'|'min'; value:number; priority?:number; scope?:string; realm?:string; tags?:string[] }
export interface SubsystemMeta { system:string; stage?:string; version?:number }
export interface SubsystemOutput { primary:Contribution[]; derived:Contribution[]; caps:CapContribution[]; context?:Record<string, ModifierPack>; meta:SubsystemMeta }
export interface ClampSpec { Min:number; Max:number }
export type EffectiveCaps = Record<string, ClampSpec>;
export interface MergeRule { usePipeline:boolean; operator?:'SUM'|'MAX'|'MIN'; clampDefault:{min:number; max:number} }
export interface Snapshot { Primary:Record<string,number>; Derived:Record<string,number>; CapsUsed:EffectiveCaps; Version:number }
export interface ModifierPack { additive_percent?:number; multipliers?:number[]; post_add?:number }
