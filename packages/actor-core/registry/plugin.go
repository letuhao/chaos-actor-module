package registry

import (
	"chaos-actor-module/packages/actor-core/interfaces"
	"fmt"
	"sort"
	"sync"
)

// PluginRegistryImpl implements the PluginRegistry interface
type PluginRegistryImpl struct {
	subsystems map[string]interfaces.Subsystem
	mu         sync.RWMutex
}

// NewPluginRegistry creates a new plugin registry
func NewPluginRegistry() interfaces.PluginRegistry {
	return &PluginRegistryImpl{
		subsystems: make(map[string]interfaces.Subsystem),
	}
}

// Register registers a subsystem
func (pr *PluginRegistryImpl) Register(subsystem interfaces.Subsystem) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	if subsystem == nil {
		return fmt.Errorf("subsystem cannot be nil")
	}

	systemID := subsystem.SystemID()
	if systemID == "" {
		return fmt.Errorf("subsystem system ID cannot be empty")
	}

	if _, exists := pr.subsystems[systemID]; exists {
		return fmt.Errorf("subsystem %s already registered", systemID)
	}

	pr.subsystems[systemID] = subsystem
	return nil
}

// Unregister unregisters a subsystem
func (pr *PluginRegistryImpl) Unregister(systemID string) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	if systemID == "" {
		return fmt.Errorf("system ID cannot be empty")
	}

	if _, exists := pr.subsystems[systemID]; !exists {
		return fmt.Errorf("subsystem %s not found", systemID)
	}

	delete(pr.subsystems, systemID)
	return nil
}

// Get returns a subsystem by ID
func (pr *PluginRegistryImpl) Get(systemID string) (interfaces.Subsystem, bool) {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	subsystem, exists := pr.subsystems[systemID]
	return subsystem, exists
}

// GetAll returns all registered subsystems
func (pr *PluginRegistryImpl) GetAll() []interfaces.Subsystem {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	subsystems := make([]interfaces.Subsystem, 0, len(pr.subsystems))
	for _, subsystem := range pr.subsystems {
		subsystems = append(subsystems, subsystem)
	}

	return subsystems
}

// GetByPriority returns subsystems ordered by priority
func (pr *PluginRegistryImpl) GetByPriority() []interfaces.Subsystem {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	subsystems := make([]interfaces.Subsystem, 0, len(pr.subsystems))
	for _, subsystem := range pr.subsystems {
		subsystems = append(subsystems, subsystem)
	}

	// Sort by priority (higher priority first)
	sort.Slice(subsystems, func(i, j int) bool {
		return subsystems[i].Priority() > subsystems[j].Priority()
	})

	return subsystems
}

// Clear clears all registered subsystems
func (pr *PluginRegistryImpl) Clear() {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	pr.subsystems = make(map[string]interfaces.Subsystem)
}

// Count returns the number of registered subsystems
func (pr *PluginRegistryImpl) Count() int {
	pr.mu.RLock()
	defer pr.mu.RUnlock()
	
	return len(pr.subsystems)
}

// HasSubsystem checks if a subsystem is registered
func (pr *PluginRegistryImpl) HasSubsystem(systemID string) bool {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	_, exists := pr.subsystems[systemID]
	return exists
}

// GetSystemIDs returns all registered system IDs
func (pr *PluginRegistryImpl) GetSystemIDs() []string {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	systemIDs := make([]string, 0, len(pr.subsystems))
	for systemID := range pr.subsystems {
		systemIDs = append(systemIDs, systemID)
	}

	return systemIDs
}

// GetByPriorityRange returns subsystems within a priority range
func (pr *PluginRegistryImpl) GetByPriorityRange(minPriority, maxPriority int64) []interfaces.Subsystem {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	subsystems := make([]interfaces.Subsystem, 0)
	for _, subsystem := range pr.subsystems {
		priority := subsystem.Priority()
		if priority >= minPriority && priority <= maxPriority {
			subsystems = append(subsystems, subsystem)
		}
	}

	// Sort by priority (higher priority first)
	sort.Slice(subsystems, func(i, j int) bool {
		return subsystems[i].Priority() > subsystems[j].Priority()
	})

	return subsystems
}

// GetHighestPriority returns the subsystem with the highest priority
func (pr *PluginRegistryImpl) GetHighestPriority() (interfaces.Subsystem, bool) {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	if len(pr.subsystems) == 0 {
		return nil, false
	}

	var highest interfaces.Subsystem
	var highestPriority int64 = -1

	for _, subsystem := range pr.subsystems {
		priority := subsystem.Priority()
		if priority > highestPriority {
			highest = subsystem
			highestPriority = priority
		}
	}

	return highest, true
}

// GetLowestPriority returns the subsystem with the lowest priority
func (pr *PluginRegistryImpl) GetLowestPriority() (interfaces.Subsystem, bool) {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	if len(pr.subsystems) == 0 {
		return nil, false
	}

	var lowest interfaces.Subsystem
	var lowestPriority int64 = 999999999

	for _, subsystem := range pr.subsystems {
		priority := subsystem.Priority()
		if priority < lowestPriority {
			lowest = subsystem
			lowestPriority = priority
		}
	}

	return lowest, true
}

// GetPriorityDistribution returns the distribution of priorities
func (pr *PluginRegistryImpl) GetPriorityDistribution() map[int64]int64 {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	distribution := make(map[int64]int64)
	for _, subsystem := range pr.subsystems {
		priority := subsystem.Priority()
		distribution[priority]++
	}

	return distribution
}

// Validate validates all registered subsystems
func (pr *PluginRegistryImpl) Validate() error {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	for systemID, subsystem := range pr.subsystems {
		if subsystem == nil {
			return fmt.Errorf("subsystem %s is nil", systemID)
		}

		if subsystem.SystemID() != systemID {
			return fmt.Errorf("subsystem system ID mismatch: expected %s, got %s",
				systemID, subsystem.SystemID())
		}

		if subsystem.SystemID() == "" {
			return fmt.Errorf("subsystem %s has empty system ID", systemID)
		}
	}

	return nil
}

// GetSubsystemNames returns the names of all registered subsystems
func (pr *PluginRegistryImpl) GetSubsystemNames() []string {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	names := make([]string, 0, len(pr.subsystems))
	for _, subsystem := range pr.subsystems {
		names = append(names, subsystem.SystemID())
	}

	sort.Strings(names)
	return names
}

// IsEmpty checks if the registry is empty
func (pr *PluginRegistryImpl) IsEmpty() bool {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	return len(pr.subsystems) == 0
}

// GetSubsystemCount returns the number of subsystems
func (pr *PluginRegistryImpl) GetSubsystemCount() int64 {
	return int64(pr.Count())
}
