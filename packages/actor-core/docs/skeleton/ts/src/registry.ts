// Actor Core v3 — Registries (skeleton)
import { MergeRule } from './types';
export const CombinerRegistryTODO = {
  ruleFor(_dimension: string): MergeRule {
    // TODO: look up dimension; default clamp to wide range
    throw new Error('TODO: combiner registry');
  }
};
export const CapLayerRegistryTODO = { order: ['REALM','WORLD','EVENT','TOTAL'], across_policy: 'INTERSECT' as const };
