# API Contracts (Go style)
type Resolver interface { ComputeSnapshot(in ComputeInput) StatSnapshot }
type SnapshotProvider interface { BuildForActor(actorID string) (StatSnapshot, error) }
type ProgressionService interface { GrantXP(...); TryBreakthrough(...); AdvanceStage(...) }

// RealmRank mở rộng: thêm HeavenTrampling1..9 → 12..20
