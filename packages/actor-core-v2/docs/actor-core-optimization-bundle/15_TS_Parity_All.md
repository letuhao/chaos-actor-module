# 15 — TypeScript Parity for All Derived (template)

Create `packages/parity/parity_all.test.ts` with:

```ts
// NOTE: Replace formulas below with the generated TS implementations.
function clamp(x:number, min:number, max:number){ return Math.max(min, Math.min(max, x)); }

export function calculate(v:{
  strength:number; agility:number; vitality:number;
  intelligence:number; spirit:number; luck:number;
  mastery_hp:number; mastery_mp:number; mastery_stamina:number;
  mastery_attack:number; mastery_magic:number; mastery_defense:number; mastery_resist:number; mastery_speed:number;
}){
  const m: Record<string, number> = {};
  m["hp_max"] = clamp((v.vitality*20 + v.strength*5)*(1 + v.mastery_hp*0.02), 1, 2000000);
  m["mp_max"] = clamp((v.intelligence*18 + v.spirit*12)*(1 + v.mastery_mp*0.02), 0, 2000000);
  m["stamina_max"] = clamp((v.vitality*5 + v.agility*8)*(1 + v.mastery_stamina*0.02), 0, 100000);
  m["attack_power"] = clamp((v.strength*3 + v.agility*1)*(1 + v.mastery_attack*0.03), 1, 1000000);
  m["magic_power"] = clamp((v.intelligence*3 + v.spirit*1)*(1 + v.mastery_magic*0.03), 0, 1000000);
  m["defense"] = clamp((v.vitality*2 + v.strength*1)*(1 + v.mastery_defense*0.02), 0, 500000);
  m["resist"] = clamp((v.spirit*2 + v.intelligence*1)*(1 + v.mastery_resist*0.02), 0, 500000);
  m["crit_rate"] = clamp(v.luck*0.001 + v.agility*0.0005, 0, 1);
  m["crit_damage"] = clamp(1.5 + v.luck*0.001, 1, 3);
  m["accuracy"] = clamp(0.7 + v.agility*0.0008 + v.luck*0.0004, 0, 1);
  m["evasion"] = clamp(v.agility*0.0009 + v.luck*0.0003, 0, 0.8);
  m["block_chance"] = clamp(v.strength*0.0006 + v.vitality*0.0004, 0, 0.6);
  m["move_speed"] = clamp((5 + v.agility*0.01)*(1 + v.mastery_speed*0.01), 0, 12);
  m["hp_regen"] = clamp(v.vitality*0.1 + v.spirit*0.05, 0, 10000);
  m["mp_regen"] = clamp(v.spirit*0.12 + v.intelligence*0.08, 0, 10000);
  m["stamina_regen"] = clamp(v.vitality*0.1 + v.agility*0.1, 0, 10000);
  m["cooldown_reduction"] = clamp(v.intelligence*0.0005 + v.spirit*0.0003, 0, 0.5);
  m["lifesteal"] = clamp(v.luck*0.0004 + v.spirit*0.0002, 0, 0.3);
  return m;
}
```

Then write a Jest test that loads Go golden `all.golden.json` and compares outputs.
