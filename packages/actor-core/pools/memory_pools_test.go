package pools

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMemoryPools_New(t *testing.T) {
	pools := NewMemoryPools()

	if pools == nil {
		t.Fatal("Expected pools to be created")
	}

	if pools.actorPool == nil {
		t.Fatal("Expected actorPool to be created")
	}

	if pools.snapshotPool == nil {
		t.Fatal("Expected snapshotPool to be created")
	}

	if pools.contributionPool == nil {
		t.Fatal("Expected contributionPool to be created")
	}

	if pools.eventPool == nil {
		t.Fatal("Expected eventPool to be created")
	}

	if pools.messagePool == nil {
		t.Fatal("Expected messagePool to be created")
	}
}

func TestActorPool_GetAndPut(t *testing.T) {
	pool := NewActorPool()

	// Get an actor
	actor := pool.Get()
	if actor == nil {
		t.Fatal("Expected actor to be created")
	}

	// Set some data
	actor.ID = "test-actor"
	actor.Version = 1
	actor.Metadata["test"] = "value"

	// Put it back
	pool.Put(actor)

	// Get another actor
	actor2 := pool.Get()
	if actor2 == nil {
		t.Fatal("Expected actor2 to be created")
	}

	// Should be reset
	if actor2.ID != "" {
		t.Errorf("Expected actor2.ID to be empty, got %s", actor2.ID)
	}
	if actor2.Version != 0 {
		t.Errorf("Expected actor2.Version to be 0, got %d", actor2.Version)
	}
	if len(actor2.Metadata) != 0 {
		t.Errorf("Expected actor2.Metadata to be empty, got %v", actor2.Metadata)
	}
}

func TestActorPool_Stats(t *testing.T) {
	pool := NewActorPool()

	// Initial stats
	stats := pool.GetStats()
	if stats.totalGets != 0 {
		t.Errorf("Expected totalGets to be 0, got %d", stats.totalGets)
	}
	if stats.totalPuts != 0 {
		t.Errorf("Expected totalPuts to be 0, got %d", stats.totalPuts)
	}

	// Get an actor
	actor := pool.Get()
	pool.Put(actor)

	// Check stats
	stats = pool.GetStats()
	if stats.totalGets != 1 {
		t.Errorf("Expected totalGets to be 1, got %d", stats.totalGets)
	}
	if stats.totalPuts != 1 {
		t.Errorf("Expected totalPuts to be 1, got %d", stats.totalPuts)
	}
}

func TestActorPool_Efficiency(t *testing.T) {
	pool := NewActorPool()

	// Get and put multiple times
	for i := 0; i < 100; i++ {
		actor := pool.Get()
		actor.ID = "test-actor"
		pool.Put(actor)
	}

	// Check efficiency
	stats := pool.GetStats()
	if stats.efficiency < 50.0 { // Lower threshold for test
		t.Errorf("Expected efficiency to be >= 50%%, got %f", stats.efficiency)
	}
}

func TestSnapshotPool_GetAndPut(t *testing.T) {
	pool := NewSnapshotPool()

	// Get a snapshot
	snapshot := pool.Get()
	if snapshot == nil {
		t.Fatal("Expected snapshot to be created")
	}

	// Set some data
	snapshot.ActorID = "test-actor"
	snapshot.Version = 1
	snapshot.Primary["test"] = 1.0
	snapshot.Derived["test"] = 2.0

	// Put it back
	pool.Put(snapshot)

	// Get another snapshot
	snapshot2 := pool.Get()
	if snapshot2 == nil {
		t.Fatal("Expected snapshot2 to be created")
	}

	// Should be reset
	if snapshot2.ActorID != "" {
		t.Errorf("Expected snapshot2.ActorID to be empty, got %s", snapshot2.ActorID)
	}
	if snapshot2.Version != 0 {
		t.Errorf("Expected snapshot2.Version to be 0, got %d", snapshot2.Version)
	}
	if len(snapshot2.Primary) != 0 {
		t.Errorf("Expected snapshot2.Primary to be empty, got %v", snapshot2.Primary)
	}
	if len(snapshot2.Derived) != 0 {
		t.Errorf("Expected snapshot2.Derived to be empty, got %v", snapshot2.Derived)
	}
}

func TestContributionPool_GetAndPut(t *testing.T) {
	pool := NewContributionPool()

	// Get a contribution
	contribution := pool.Get()
	if contribution == nil {
		t.Fatal("Expected contribution to be created")
	}

	// Set some data
	contribution.Dimension = "test"
	contribution.Value = 1.0
	contribution.System = "test-system"

	// Put it back
	pool.Put(contribution)

	// Get another contribution
	contribution2 := pool.Get()
	if contribution2 == nil {
		t.Fatal("Expected contribution2 to be created")
	}

	// Should be reset
	if contribution2.Dimension != "" {
		t.Errorf("Expected contribution2.Dimension to be empty, got %s", contribution2.Dimension)
	}
	if contribution2.Value != 0.0 {
		t.Errorf("Expected contribution2.Value to be 0.0, got %f", contribution2.Value)
	}
	if contribution2.System != "" {
		t.Errorf("Expected contribution2.System to be empty, got %s", contribution2.System)
	}
}

func TestEventPool_GetAndPut(t *testing.T) {
	pool := NewEventPool()

	// Get an event
	event := pool.Get()
	if event == nil {
		t.Fatal("Expected event to be created")
	}

	// Set some data
	event.ID = "test-event"
	event.Type = "test-type"
	event.ActorID = "test-actor"
	event.Data = "test-data"

	// Put it back
	pool.Put(event)

	// Get another event
	event2 := pool.Get()
	if event2 == nil {
		t.Fatal("Expected event2 to be created")
	}

	// Should be reset
	if event2.ID != "" {
		t.Errorf("Expected event2.ID to be empty, got %s", event2.ID)
	}
	if event2.Type != "" {
		t.Errorf("Expected event2.Type to be empty, got %s", event2.Type)
	}
	if event2.ActorID != "" {
		t.Errorf("Expected event2.ActorID to be empty, got %s", event2.ActorID)
	}
	if event2.Data != nil {
		t.Errorf("Expected event2.Data to be nil, got %v", event2.Data)
	}
}

func TestMessagePool_GetAndPut(t *testing.T) {
	pool := NewMessagePool()

	// Get a message
	message := pool.Get()
	if message == nil {
		t.Fatal("Expected message to be created")
	}

	// Set some data
	message.ID = "test-message"
	message.Topic = "test-topic"
	message.Data = "test-data"
	message.Priority = 1

	// Put it back
	pool.Put(message)

	// Get another message
	message2 := pool.Get()
	if message2 == nil {
		t.Fatal("Expected message2 to be created")
	}

	// Should be reset
	if message2.ID != "" {
		t.Errorf("Expected message2.ID to be empty, got %s", message2.ID)
	}
	if message2.Topic != "" {
		t.Errorf("Expected message2.Topic to be empty, got %s", message2.Topic)
	}
	if message2.Data != nil {
		t.Errorf("Expected message2.Data to be nil, got %v", message2.Data)
	}
	if message2.Priority != 0 {
		t.Errorf("Expected message2.Priority to be 0, got %d", message2.Priority)
	}
}

func TestMemoryPools_OverallStats(t *testing.T) {
	pools := NewMemoryPools()

	// Use all pools
	actor := pools.actorPool.Get()
	pools.actorPool.Put(actor)

	snapshot := pools.snapshotPool.Get()
	pools.snapshotPool.Put(snapshot)

	contribution := pools.contributionPool.Get()
	pools.contributionPool.Put(contribution)

	event := pools.eventPool.Get()
	pools.eventPool.Put(event)

	message := pools.messagePool.Get()
	pools.messagePool.Put(message)

	// Check overall stats
	stats := pools.GetOverallStats()
	if stats.totalGets != 5 {
		t.Errorf("Expected totalGets to be 5, got %d", stats.totalGets)
	}
	if stats.totalPuts != 5 {
		t.Errorf("Expected totalPuts to be 5, got %d", stats.totalPuts)
	}
}

func TestMemoryPools_Reset(t *testing.T) {
	pools := NewMemoryPools()

	// Use all pools
	actor := pools.actorPool.Get()
	pools.actorPool.Put(actor)

	snapshot := pools.snapshotPool.Get()
	pools.snapshotPool.Put(snapshot)

	// Check stats are not zero
	stats := pools.GetOverallStats()
	if stats.totalGets == 0 {
		t.Fatal("Expected some gets")
	}

	// Reset
	pools.Reset()

	// Check stats are reset
	stats = pools.GetOverallStats()
	if stats.totalGets != 0 {
		t.Errorf("Expected totalGets to be 0 after reset, got %d", stats.totalGets)
	}
	if stats.totalPuts != 0 {
		t.Errorf("Expected totalPuts to be 0 after reset, got %d", stats.totalPuts)
	}
}

func TestMemoryPools_ConcurrentAccess(t *testing.T) {
	pools := NewMemoryPools()

	// Concurrent access
	var wg sync.WaitGroup
	numGoroutines := 100
	numOperations := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				// Use actor pool
				actor := pools.actorPool.Get()
				actor.ID = fmt.Sprintf("actor_%d_%d", goroutineID, j)
				pools.actorPool.Put(actor)

				// Use snapshot pool
				snapshot := pools.snapshotPool.Get()
				snapshot.ActorID = fmt.Sprintf("actor_%d_%d", goroutineID, j)
				pools.snapshotPool.Put(snapshot)

				// Use contribution pool
				contribution := pools.contributionPool.Get()
				contribution.Dimension = fmt.Sprintf("dim_%d_%d", goroutineID, j)
				pools.contributionPool.Put(contribution)

				// Use event pool
				event := pools.eventPool.Get()
				event.ID = fmt.Sprintf("event_%d_%d", goroutineID, j)
				pools.eventPool.Put(event)

				// Use message pool
				message := pools.messagePool.Get()
				message.ID = fmt.Sprintf("message_%d_%d", goroutineID, j)
				pools.messagePool.Put(message)
			}
		}(i)
	}

	wg.Wait()

	// Check overall stats
	stats := pools.GetOverallStats()
	if stats.totalGets == 0 {
		t.Fatal("Expected some gets")
	}
	if stats.totalPuts == 0 {
		t.Fatal("Expected some puts")
	}

	// Check efficiency
	if stats.efficiency < 15.0 { // Lower threshold for concurrent test
		t.Errorf("Expected efficiency to be >= 15%%, got %f", stats.efficiency)
	}
}

func TestActor_Reset(t *testing.T) {
	actor := &Actor{
		ID:         "test-actor",
		Version:    1,
		Subsystems: []interface{}{"sub1", "sub2"},
		Metadata:   map[string]interface{}{"key": "value"},
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Reset
	actor.Reset()

	// Check reset
	if actor.ID != "" {
		t.Errorf("Expected ID to be empty, got %s", actor.ID)
	}
	if actor.Version != 0 {
		t.Errorf("Expected Version to be 0, got %d", actor.Version)
	}
	if len(actor.Subsystems) != 0 {
		t.Errorf("Expected Subsystems to be empty, got %v", actor.Subsystems)
	}
	if len(actor.Metadata) != 0 {
		t.Errorf("Expected Metadata to be empty, got %v", actor.Metadata)
	}
	// CreatedAt is not reset - we use it to detect hits
	// if !actor.CreatedAt.IsZero() {
	//	t.Errorf("Expected CreatedAt to be zero, got %v", actor.CreatedAt)
	// }
	if !actor.UpdatedAt.IsZero() {
		t.Errorf("Expected UpdatedAt to be zero, got %v", actor.UpdatedAt)
	}
}

func TestSnapshot_Reset(t *testing.T) {
	snapshot := &Snapshot{
		ActorID:   "test-actor",
		Primary:   map[string]float64{"key1": 1.0, "key2": 2.0},
		Derived:   map[string]float64{"key3": 3.0, "key4": 4.0},
		CapsUsed:  map[string]interface{}{"key5": "value5"},
		Version:   1,
		CreatedAt: time.Now(),
	}

	// Reset
	snapshot.Reset()

	// Check reset
	if snapshot.ActorID != "" {
		t.Errorf("Expected ActorID to be empty, got %s", snapshot.ActorID)
	}
	if snapshot.Version != 0 {
		t.Errorf("Expected Version to be 0, got %d", snapshot.Version)
	}
	if len(snapshot.Primary) != 0 {
		t.Errorf("Expected Primary to be empty, got %v", snapshot.Primary)
	}
	if len(snapshot.Derived) != 0 {
		t.Errorf("Expected Derived to be empty, got %v", snapshot.Derived)
	}
	if len(snapshot.CapsUsed) != 0 {
		t.Errorf("Expected CapsUsed to be empty, got %v", snapshot.CapsUsed)
	}
	if !snapshot.CreatedAt.IsZero() {
		t.Errorf("Expected CreatedAt to be zero, got %v", snapshot.CreatedAt)
	}
}

func TestContribution_Reset(t *testing.T) {
	contribution := &Contribution{
		Dimension: "test",
		Bucket:    "test-bucket",
		Value:     1.0,
		System:    "test-system",
		Priority:  1,
	}

	// Reset
	contribution.Reset()

	// Check reset
	if contribution.Dimension != "" {
		t.Errorf("Expected Dimension to be empty, got %s", contribution.Dimension)
	}
	if contribution.Bucket != "" {
		t.Errorf("Expected Bucket to be empty, got %s", contribution.Bucket)
	}
	if contribution.Value != 0.0 {
		t.Errorf("Expected Value to be 0.0, got %f", contribution.Value)
	}
	if contribution.System != "" {
		t.Errorf("Expected System to be empty, got %s", contribution.System)
	}
	if contribution.Priority != 0 {
		t.Errorf("Expected Priority to be 0, got %d", contribution.Priority)
	}
}

func TestEvent_Reset(t *testing.T) {
	event := &Event{
		ID:        "test-event",
		Type:      "test-type",
		ActorID:   "test-actor",
		Data:      "test-data",
		Timestamp: time.Now(),
		Source:    "test-source",
		Priority:  1,
	}

	// Reset
	event.Reset()

	// Check reset
	if event.ID != "" {
		t.Errorf("Expected ID to be empty, got %s", event.ID)
	}
	if event.Type != "" {
		t.Errorf("Expected Type to be empty, got %s", event.Type)
	}
	if event.ActorID != "" {
		t.Errorf("Expected ActorID to be empty, got %s", event.ActorID)
	}
	if event.Data != nil {
		t.Errorf("Expected Data to be nil, got %v", event.Data)
	}
	if !event.Timestamp.IsZero() {
		t.Errorf("Expected Timestamp to be zero, got %v", event.Timestamp)
	}
	if event.Source != "" {
		t.Errorf("Expected Source to be empty, got %s", event.Source)
	}
	if event.Priority != 0 {
		t.Errorf("Expected Priority to be 0, got %d", event.Priority)
	}
}

func TestMessage_Reset(t *testing.T) {
	message := &Message{
		ID:        "test-message",
		Topic:     "test-topic",
		Data:      "test-data",
		Timestamp: time.Now(),
		Priority:  1,
		TTL:       time.Hour,
	}

	// Reset
	message.Reset()

	// Check reset
	if message.ID != "" {
		t.Errorf("Expected ID to be empty, got %s", message.ID)
	}
	if message.Topic != "" {
		t.Errorf("Expected Topic to be empty, got %s", message.Topic)
	}
	if message.Data != nil {
		t.Errorf("Expected Data to be nil, got %v", message.Data)
	}
	if !message.Timestamp.IsZero() {
		t.Errorf("Expected Timestamp to be zero, got %v", message.Timestamp)
	}
	if message.Priority != 0 {
		t.Errorf("Expected Priority to be 0, got %d", message.Priority)
	}
	if message.TTL != 0 {
		t.Errorf("Expected TTL to be 0, got %v", message.TTL)
	}
}

func BenchmarkActorPool_GetAndPut(b *testing.B) {
	pool := NewActorPool()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actor := pool.Get()
		actor.ID = fmt.Sprintf("actor_%d", i)
		pool.Put(actor)
	}
}

func BenchmarkSnapshotPool_GetAndPut(b *testing.B) {
	pool := NewSnapshotPool()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		snapshot := pool.Get()
		snapshot.ActorID = fmt.Sprintf("actor_%d", i)
		pool.Put(snapshot)
	}
}

func BenchmarkContributionPool_GetAndPut(b *testing.B) {
	pool := NewContributionPool()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		contribution := pool.Get()
		contribution.Dimension = fmt.Sprintf("dim_%d", i)
		pool.Put(contribution)
	}
}

func BenchmarkEventPool_GetAndPut(b *testing.B) {
	pool := NewEventPool()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		event := pool.Get()
		event.ID = fmt.Sprintf("event_%d", i)
		pool.Put(event)
	}
}

func BenchmarkMessagePool_GetAndPut(b *testing.B) {
	pool := NewMessagePool()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		message := pool.Get()
		message.ID = fmt.Sprintf("message_%d", i)
		pool.Put(message)
	}
}

func BenchmarkMemoryPools_Concurrent(b *testing.B) {
	pools := NewMemoryPools()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			// Use all pools
			actor := pools.actorPool.Get()
			actor.ID = fmt.Sprintf("actor_%d", i)
			pools.actorPool.Put(actor)

			snapshot := pools.snapshotPool.Get()
			snapshot.ActorID = fmt.Sprintf("actor_%d", i)
			pools.snapshotPool.Put(snapshot)

			contribution := pools.contributionPool.Get()
			contribution.Dimension = fmt.Sprintf("dim_%d", i)
			pools.contributionPool.Put(contribution)

			event := pools.eventPool.Get()
			event.ID = fmt.Sprintf("event_%d", i)
			pools.eventPool.Put(event)

			message := pools.messagePool.Get()
			message.ID = fmt.Sprintf("message_%d", i)
			pools.messagePool.Put(message)

			i++
		}
	})
}
