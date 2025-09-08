package pools

import (
	"sync"
	"sync/atomic"
	"time"
)

// MemoryPools manages all memory pools for zero-allocation operations
type MemoryPools struct {
	actorPool        *ActorPool
	snapshotPool     *SnapshotPool
	contributionPool *ContributionPool
	eventPool        *EventPool
	messagePool      *MessagePool
	stats            *PoolStats
}

// PoolStats represents overall pool statistics
type PoolStats struct {
	totalGets     int64
	totalPuts     int64
	totalHits     int64
	totalMisses   int64
	totalAllocs   int64
	totalDeallocs int64
	efficiency    float64
}

// ActorPool manages Actor objects
type ActorPool struct {
	pool  *sync.Pool
	stats *PoolStats
}

// SnapshotPool manages Snapshot objects
type SnapshotPool struct {
	pool  *sync.Pool
	stats *PoolStats
}

// ContributionPool manages Contribution objects
type ContributionPool struct {
	pool  *sync.Pool
	stats *PoolStats
}

// EventPool manages Event objects
type EventPool struct {
	pool  *sync.Pool
	stats *PoolStats
}

// MessagePool manages Message objects
type MessagePool struct {
	pool  *sync.Pool
	stats *PoolStats
}

// Actor represents a cached actor (placeholder - will be replaced with actual Actor type)
type Actor struct {
	ID         string
	Version    int64
	Subsystems []interface{} // placeholder
	Metadata   map[string]interface{}
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsReused   bool // Flag to detect if object was reused from pool
}

// Snapshot represents a cached snapshot (placeholder - will be replaced with actual Snapshot type)
type Snapshot struct {
	ActorID   string
	Primary   map[string]float64
	Derived   map[string]float64
	CapsUsed  map[string]interface{} // placeholder
	Version   int64
	CreatedAt time.Time
}

// Contribution represents a cached contribution (placeholder - will be replaced with actual Contribution type)
type Contribution struct {
	Dimension string
	Bucket    string
	Value     float64
	System    string
	Priority  int64
}

// Event represents a cached event (placeholder - will be replaced with actual Event type)
type Event struct {
	ID        string
	Type      string
	ActorID   string
	Data      interface{}
	Timestamp time.Time
	Source    string
	Priority  int
}

// Message represents a cached message (placeholder - will be replaced with actual Message type)
type Message struct {
	ID        string
	Topic     string
	Data      interface{}
	Timestamp time.Time
	Priority  int
	TTL       time.Duration
}

// NewMemoryPools creates a new memory pools manager
func NewMemoryPools() *MemoryPools {
	return &MemoryPools{
		actorPool:        NewActorPool(),
		snapshotPool:     NewSnapshotPool(),
		contributionPool: NewContributionPool(),
		eventPool:        NewEventPool(),
		messagePool:      NewMessagePool(),
		stats:            &PoolStats{},
	}
}

// NewActorPool creates a new actor pool
func NewActorPool() *ActorPool {
	return &ActorPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &Actor{
					ID:         "",
					Version:    0,
					Subsystems: make([]interface{}, 0, 10),
					Metadata:   make(map[string]interface{}, 10),
					CreatedAt:  time.Time{},
					UpdatedAt:  time.Time{},
					IsReused:   false,
				}
			},
		},
		stats: &PoolStats{},
	}
}

// Get retrieves an Actor from the pool
func (p *ActorPool) Get() *Actor {
	actor := p.pool.Get().(*Actor)
	atomic.AddInt64(&p.stats.totalGets, 1)

	// Check if this is a hit (reused object) or miss (new object)
	if actor.IsReused {
		// Reused object from pool
		atomic.AddInt64(&p.stats.totalHits, 1)
		actor.IsReused = false // Reset for next use
	} else {
		// New object from pool
		atomic.AddInt64(&p.stats.totalMisses, 1)
		atomic.AddInt64(&p.stats.totalAllocs, 1)
	}

	return actor
}

// Put returns an Actor to the pool
func (p *ActorPool) Put(actor *Actor) {
	if actor == nil {
		return
	}

	// Reset actor for reuse
	actor.Reset()
	actor.IsReused = true // Mark as reused for next Get
	p.pool.Put(actor)

	atomic.AddInt64(&p.stats.totalPuts, 1)
	atomic.AddInt64(&p.stats.totalDeallocs, 1)
}

// Reset resets an Actor for reuse
func (a *Actor) Reset() {
	a.ID = ""
	a.Version = 0
	a.Subsystems = a.Subsystems[:0] // Reset slice length but keep capacity
	a.CreatedAt = time.Time{}
	a.UpdatedAt = time.Time{}
	// Don't reset IsReused - it will be set in Put function

	// Clear metadata map
	for k := range a.Metadata {
		delete(a.Metadata, k)
	}
}

// GetStats returns pool statistics
func (p *ActorPool) GetStats() *PoolStats {
	return &PoolStats{
		totalGets:     atomic.LoadInt64(&p.stats.totalGets),
		totalPuts:     atomic.LoadInt64(&p.stats.totalPuts),
		totalHits:     atomic.LoadInt64(&p.stats.totalHits),
		totalMisses:   atomic.LoadInt64(&p.stats.totalMisses),
		totalAllocs:   atomic.LoadInt64(&p.stats.totalAllocs),
		totalDeallocs: atomic.LoadInt64(&p.stats.totalDeallocs),
		efficiency:    p.calculateEfficiency(),
	}
}

// calculateEfficiency calculates pool efficiency
func (p *ActorPool) calculateEfficiency() float64 {
	gets := atomic.LoadInt64(&p.stats.totalGets)
	hits := atomic.LoadInt64(&p.stats.totalHits)

	if gets == 0 {
		return 0.0
	}

	return float64(hits) / float64(gets) * 100.0
}

// NewSnapshotPool creates a new snapshot pool
func NewSnapshotPool() *SnapshotPool {
	return &SnapshotPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &Snapshot{
					ActorID:   "",
					Primary:   make(map[string]float64, 50),
					Derived:   make(map[string]float64, 50),
					CapsUsed:  make(map[string]interface{}, 20),
					Version:   0,
					CreatedAt: time.Time{},
				}
			},
		},
		stats: &PoolStats{},
	}
}

// Get retrieves a Snapshot from the pool
func (p *SnapshotPool) Get() *Snapshot {
	snapshot := p.pool.Get().(*Snapshot)
	atomic.AddInt64(&p.stats.totalGets, 1)

	if snapshot.ActorID == "" {
		atomic.AddInt64(&p.stats.totalMisses, 1)
		atomic.AddInt64(&p.stats.totalAllocs, 1)
	} else {
		atomic.AddInt64(&p.stats.totalHits, 1)
	}

	return snapshot
}

// Put returns a Snapshot to the pool
func (p *SnapshotPool) Put(snapshot *Snapshot) {
	if snapshot == nil {
		return
	}

	snapshot.Reset()
	p.pool.Put(snapshot)

	atomic.AddInt64(&p.stats.totalPuts, 1)
	atomic.AddInt64(&p.stats.totalDeallocs, 1)
}

// Reset resets a Snapshot for reuse
func (s *Snapshot) Reset() {
	s.ActorID = ""
	s.Version = 0
	s.CreatedAt = time.Time{}

	// Clear maps but keep capacity
	for k := range s.Primary {
		delete(s.Primary, k)
	}
	for k := range s.Derived {
		delete(s.Derived, k)
	}
	for k := range s.CapsUsed {
		delete(s.CapsUsed, k)
	}
}

// GetStats returns pool statistics
func (p *SnapshotPool) GetStats() *PoolStats {
	return &PoolStats{
		totalGets:     atomic.LoadInt64(&p.stats.totalGets),
		totalPuts:     atomic.LoadInt64(&p.stats.totalPuts),
		totalHits:     atomic.LoadInt64(&p.stats.totalHits),
		totalMisses:   atomic.LoadInt64(&p.stats.totalMisses),
		totalAllocs:   atomic.LoadInt64(&p.stats.totalAllocs),
		totalDeallocs: atomic.LoadInt64(&p.stats.totalDeallocs),
		efficiency:    p.calculateEfficiency(),
	}
}

// calculateEfficiency calculates pool efficiency
func (p *SnapshotPool) calculateEfficiency() float64 {
	gets := atomic.LoadInt64(&p.stats.totalGets)
	hits := atomic.LoadInt64(&p.stats.totalHits)

	if gets == 0 {
		return 0.0
	}

	return float64(hits) / float64(gets) * 100.0
}

// NewContributionPool creates a new contribution pool
func NewContributionPool() *ContributionPool {
	return &ContributionPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &Contribution{
					Dimension: "",
					Bucket:    "",
					Value:     0.0,
					System:    "",
					Priority:  0,
				}
			},
		},
		stats: &PoolStats{},
	}
}

// Get retrieves a Contribution from the pool
func (p *ContributionPool) Get() *Contribution {
	contribution := p.pool.Get().(*Contribution)
	atomic.AddInt64(&p.stats.totalGets, 1)

	if contribution.Dimension == "" {
		atomic.AddInt64(&p.stats.totalMisses, 1)
		atomic.AddInt64(&p.stats.totalAllocs, 1)
	} else {
		atomic.AddInt64(&p.stats.totalHits, 1)
	}

	return contribution
}

// Put returns a Contribution to the pool
func (p *ContributionPool) Put(contribution *Contribution) {
	if contribution == nil {
		return
	}

	contribution.Reset()
	p.pool.Put(contribution)

	atomic.AddInt64(&p.stats.totalPuts, 1)
	atomic.AddInt64(&p.stats.totalDeallocs, 1)
}

// Reset resets a Contribution for reuse
func (c *Contribution) Reset() {
	c.Dimension = ""
	c.Bucket = ""
	c.Value = 0.0
	c.System = ""
	c.Priority = 0
}

// GetStats returns pool statistics
func (p *ContributionPool) GetStats() *PoolStats {
	return &PoolStats{
		totalGets:     atomic.LoadInt64(&p.stats.totalGets),
		totalPuts:     atomic.LoadInt64(&p.stats.totalPuts),
		totalHits:     atomic.LoadInt64(&p.stats.totalHits),
		totalMisses:   atomic.LoadInt64(&p.stats.totalMisses),
		totalAllocs:   atomic.LoadInt64(&p.stats.totalAllocs),
		totalDeallocs: atomic.LoadInt64(&p.stats.totalDeallocs),
		efficiency:    p.calculateEfficiency(),
	}
}

// calculateEfficiency calculates pool efficiency
func (p *ContributionPool) calculateEfficiency() float64 {
	gets := atomic.LoadInt64(&p.stats.totalGets)
	hits := atomic.LoadInt64(&p.stats.totalHits)

	if gets == 0 {
		return 0.0
	}

	return float64(hits) / float64(gets) * 100.0
}

// NewEventPool creates a new event pool
func NewEventPool() *EventPool {
	return &EventPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &Event{
					ID:        "",
					Type:      "",
					ActorID:   "",
					Data:      nil,
					Timestamp: time.Time{},
					Source:    "",
					Priority:  0,
				}
			},
		},
		stats: &PoolStats{},
	}
}

// Get retrieves an Event from the pool
func (p *EventPool) Get() *Event {
	event := p.pool.Get().(*Event)
	atomic.AddInt64(&p.stats.totalGets, 1)

	if event.ID == "" {
		atomic.AddInt64(&p.stats.totalMisses, 1)
		atomic.AddInt64(&p.stats.totalAllocs, 1)
	} else {
		atomic.AddInt64(&p.stats.totalHits, 1)
	}

	return event
}

// Put returns an Event to the pool
func (p *EventPool) Put(event *Event) {
	if event == nil {
		return
	}

	event.Reset()
	p.pool.Put(event)

	atomic.AddInt64(&p.stats.totalPuts, 1)
	atomic.AddInt64(&p.stats.totalDeallocs, 1)
}

// Reset resets an Event for reuse
func (e *Event) Reset() {
	e.ID = ""
	e.Type = ""
	e.ActorID = ""
	e.Data = nil
	e.Timestamp = time.Time{}
	e.Source = ""
	e.Priority = 0
}

// GetStats returns pool statistics
func (p *EventPool) GetStats() *PoolStats {
	return &PoolStats{
		totalGets:     atomic.LoadInt64(&p.stats.totalGets),
		totalPuts:     atomic.LoadInt64(&p.stats.totalPuts),
		totalHits:     atomic.LoadInt64(&p.stats.totalHits),
		totalMisses:   atomic.LoadInt64(&p.stats.totalMisses),
		totalAllocs:   atomic.LoadInt64(&p.stats.totalAllocs),
		totalDeallocs: atomic.LoadInt64(&p.stats.totalDeallocs),
		efficiency:    p.calculateEfficiency(),
	}
}

// calculateEfficiency calculates pool efficiency
func (p *EventPool) calculateEfficiency() float64 {
	gets := atomic.LoadInt64(&p.stats.totalGets)
	hits := atomic.LoadInt64(&p.stats.totalHits)

	if gets == 0 {
		return 0.0
	}

	return float64(hits) / float64(gets) * 100.0
}

// NewMessagePool creates a new message pool
func NewMessagePool() *MessagePool {
	return &MessagePool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &Message{
					ID:        "",
					Topic:     "",
					Data:      nil,
					Timestamp: time.Time{},
					Priority:  0,
					TTL:       0,
				}
			},
		},
		stats: &PoolStats{},
	}
}

// Get retrieves a Message from the pool
func (p *MessagePool) Get() *Message {
	message := p.pool.Get().(*Message)
	atomic.AddInt64(&p.stats.totalGets, 1)

	if message.ID == "" {
		atomic.AddInt64(&p.stats.totalMisses, 1)
		atomic.AddInt64(&p.stats.totalAllocs, 1)
	} else {
		atomic.AddInt64(&p.stats.totalHits, 1)
	}

	return message
}

// Put returns a Message to the pool
func (p *MessagePool) Put(message *Message) {
	if message == nil {
		return
	}

	message.Reset()
	p.pool.Put(message)

	atomic.AddInt64(&p.stats.totalPuts, 1)
	atomic.AddInt64(&p.stats.totalDeallocs, 1)
}

// Reset resets a Message for reuse
func (m *Message) Reset() {
	m.ID = ""
	m.Topic = ""
	m.Data = nil
	m.Timestamp = time.Time{}
	m.Priority = 0
	m.TTL = 0
}

// GetStats returns pool statistics
func (p *MessagePool) GetStats() *PoolStats {
	return &PoolStats{
		totalGets:     atomic.LoadInt64(&p.stats.totalGets),
		totalPuts:     atomic.LoadInt64(&p.stats.totalPuts),
		totalHits:     atomic.LoadInt64(&p.stats.totalHits),
		totalMisses:   atomic.LoadInt64(&p.stats.totalMisses),
		totalAllocs:   atomic.LoadInt64(&p.stats.totalAllocs),
		totalDeallocs: atomic.LoadInt64(&p.stats.totalDeallocs),
		efficiency:    p.calculateEfficiency(),
	}
}

// calculateEfficiency calculates pool efficiency
func (p *MessagePool) calculateEfficiency() float64 {
	gets := atomic.LoadInt64(&p.stats.totalGets)
	hits := atomic.LoadInt64(&p.stats.totalHits)

	if gets == 0 {
		return 0.0
	}

	return float64(hits) / float64(gets) * 100.0
}

// GetOverallStats returns overall pool statistics
func (mp *MemoryPools) GetOverallStats() *PoolStats {
	actorStats := mp.actorPool.GetStats()
	snapshotStats := mp.snapshotPool.GetStats()
	contributionStats := mp.contributionPool.GetStats()
	eventStats := mp.eventPool.GetStats()
	messageStats := mp.messagePool.GetStats()

	totalGets := actorStats.totalGets + snapshotStats.totalGets + contributionStats.totalGets + eventStats.totalGets + messageStats.totalGets
	totalHits := actorStats.totalHits + snapshotStats.totalHits + contributionStats.totalHits + eventStats.totalHits + messageStats.totalHits

	var efficiency float64
	if totalGets > 0 {
		efficiency = float64(totalHits) / float64(totalGets) * 100.0
	}

	return &PoolStats{
		totalGets:     totalGets,
		totalPuts:     actorStats.totalPuts + snapshotStats.totalPuts + contributionStats.totalPuts + eventStats.totalPuts + messageStats.totalPuts,
		totalHits:     totalHits,
		totalMisses:   actorStats.totalMisses + snapshotStats.totalMisses + contributionStats.totalMisses + eventStats.totalMisses + messageStats.totalMisses,
		totalAllocs:   actorStats.totalAllocs + snapshotStats.totalAllocs + contributionStats.totalAllocs + eventStats.totalAllocs + messageStats.totalAllocs,
		totalDeallocs: actorStats.totalDeallocs + snapshotStats.totalDeallocs + contributionStats.totalDeallocs + eventStats.totalDeallocs + messageStats.totalDeallocs,
		efficiency:    efficiency,
	}
}

// Reset resets all pool statistics
func (mp *MemoryPools) Reset() {
	atomic.StoreInt64(&mp.actorPool.stats.totalGets, 0)
	atomic.StoreInt64(&mp.actorPool.stats.totalPuts, 0)
	atomic.StoreInt64(&mp.actorPool.stats.totalHits, 0)
	atomic.StoreInt64(&mp.actorPool.stats.totalMisses, 0)
	atomic.StoreInt64(&mp.actorPool.stats.totalAllocs, 0)
	atomic.StoreInt64(&mp.actorPool.stats.totalDeallocs, 0)

	atomic.StoreInt64(&mp.snapshotPool.stats.totalGets, 0)
	atomic.StoreInt64(&mp.snapshotPool.stats.totalPuts, 0)
	atomic.StoreInt64(&mp.snapshotPool.stats.totalHits, 0)
	atomic.StoreInt64(&mp.snapshotPool.stats.totalMisses, 0)
	atomic.StoreInt64(&mp.snapshotPool.stats.totalAllocs, 0)
	atomic.StoreInt64(&mp.snapshotPool.stats.totalDeallocs, 0)

	atomic.StoreInt64(&mp.contributionPool.stats.totalGets, 0)
	atomic.StoreInt64(&mp.contributionPool.stats.totalPuts, 0)
	atomic.StoreInt64(&mp.contributionPool.stats.totalHits, 0)
	atomic.StoreInt64(&mp.contributionPool.stats.totalMisses, 0)
	atomic.StoreInt64(&mp.contributionPool.stats.totalAllocs, 0)
	atomic.StoreInt64(&mp.contributionPool.stats.totalDeallocs, 0)

	atomic.StoreInt64(&mp.eventPool.stats.totalGets, 0)
	atomic.StoreInt64(&mp.eventPool.stats.totalPuts, 0)
	atomic.StoreInt64(&mp.eventPool.stats.totalHits, 0)
	atomic.StoreInt64(&mp.eventPool.stats.totalMisses, 0)
	atomic.StoreInt64(&mp.eventPool.stats.totalAllocs, 0)
	atomic.StoreInt64(&mp.eventPool.stats.totalDeallocs, 0)

	atomic.StoreInt64(&mp.messagePool.stats.totalGets, 0)
	atomic.StoreInt64(&mp.messagePool.stats.totalPuts, 0)
	atomic.StoreInt64(&mp.messagePool.stats.totalHits, 0)
	atomic.StoreInt64(&mp.messagePool.stats.totalMisses, 0)
	atomic.StoreInt64(&mp.messagePool.stats.totalAllocs, 0)
	atomic.StoreInt64(&mp.messagePool.stats.totalDeallocs, 0)
}
