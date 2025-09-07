// Actor Core v3 — Caps Provider (skeleton placeholder)
import { CapsProvider, CombinerRegistry } from './interfaces';
import { EffectiveCaps } from './types';
export const CapsProviderTODO: CapsProvider = {
  effectiveCapsWithinLayer(_actor, _outs, _layer): EffectiveCaps {
    // TODO: implement within-layer merge (Section 07)
    throw new Error('TODO: within-layer caps');
  },
  reduceAcrossLayers(_layerCaps, _combiner: CombinerRegistry, _layerReg): EffectiveCaps {
    // TODO: implement across-layer reduction (Section 07)
    throw new Error('TODO: across-layer reduction');
  }
};
