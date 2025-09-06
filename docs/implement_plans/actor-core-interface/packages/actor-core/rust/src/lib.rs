// packages/actor-core/rust/src/lib.rs
//! Actor Core interfaces only; no implementation. Keep functions pure.

use std::collections::BTreeMap;

#[derive(Clone, Debug, Default)]
pub struct PrimaryCore {
    pub hp_max: i32,
    pub life_span: i32,
    pub attack: i32,
    pub defense: i32,
    pub speed: i32,
}

#[derive(Clone, Debug, Default)]
pub struct Derived {
    pub hp_max: f64,
    pub mp_max: f64,
    pub atk: f64,
    pub mag: f64,
    pub def_: f64,
    pub res: f64,
    pub haste: f64,
    pub crit_chance: f64,
    pub crit_multi: f64,
    pub move_speed: f64,
    pub regen_hp: f64,
    pub regen_mp: f64,
    pub resists: BTreeMap<String, f64>,
    pub amplifiers: BTreeMap<String, f64>,
    pub version: u64,
}

#[derive(Clone, Debug, Default)]
pub struct CoreContribution {
    pub primary: Option<PrimaryCore>,
    pub flat: BTreeMap<String, f64>,
    pub mult: BTreeMap<String, f64>,
    pub tags: Vec<String>,
}

#[derive(Clone, Debug, Default)]
pub struct ResourceState { pub current: f64, pub max: f64, pub regen: f64, pub epoch: u64 }
pub type ActorResources = BTreeMap<String, ResourceState>;

pub trait ActorCore {
    fn compose_core(&self, buckets: BTreeMap<String, CoreContribution>) -> CoreContribution;
    fn base_from_primary(&self, p: PrimaryCore, level: i32) -> Derived;
    fn finalize_derived(&self, base: Derived, flat: BTreeMap<String, f64>, mult: BTreeMap<String, f64>) -> Derived;
    fn clamp_derived(&self, d: Derived) -> Derived;
}

// Laws:
// - Merge buckets by lexicographic key; flats sum; mults multiply (missing -> 1.0).
// - finalize_derived: apply flat -> mult -> clamp, then increment version.
