package registry

import (
	"chaos-actor-module/packages/actor-core/interfaces"
	"fmt"
)

// RegistryFactory creates registry instances
type RegistryFactory struct{}

// NewRegistryFactory creates a new registry factory
func NewRegistryFactory() *RegistryFactory {
	return &RegistryFactory{}
}

// CreateCombinerRegistry creates a new combiner registry
func (rf *RegistryFactory) CreateCombinerRegistry() interfaces.CombinerRegistry {
	return NewCombinerRegistry()
}

// CreateCombinerRegistryFromFile creates a new combiner registry from a file
func (rf *RegistryFactory) CreateCombinerRegistryFromFile(filePath string) (interfaces.CombinerRegistry, error) {
	return NewCombinerRegistryFromFile(filePath)
}

// CreateCapLayerRegistry creates a new cap layer registry
func (rf *RegistryFactory) CreateCapLayerRegistry() interfaces.CapLayerRegistry {
	return NewCapLayerRegistry()
}

// CreateCapLayerRegistryFromFile creates a new cap layer registry from a file
func (rf *RegistryFactory) CreateCapLayerRegistryFromFile(filePath string) (interfaces.CapLayerRegistry, error) {
	return NewCapLayerRegistryFromFile(filePath)
}

// CreatePluginRegistry creates a new plugin registry
func (rf *RegistryFactory) CreatePluginRegistry() interfaces.PluginRegistry {
	return NewPluginRegistry()
}

// CreateConfigLoader creates a new config loader
func (rf *RegistryFactory) CreateConfigLoader() interfaces.ConfigLoader {
	return NewConfigLoader()
}

// CreateCache creates a new cache
func (rf *RegistryFactory) CreateCache(maxSize int64, evictionPolicy string) interfaces.Cache {
	return NewCache(maxSize, evictionPolicy)
}

// CreateAllRegistries creates all registry types
func (rf *RegistryFactory) CreateAllRegistries() (*RegistrySet, error) {
	combinerRegistry := rf.CreateCombinerRegistry()
	capLayerRegistry := rf.CreateCapLayerRegistry()
	pluginRegistry := rf.CreatePluginRegistry()
	configLoader := rf.CreateConfigLoader()
	cache := rf.CreateCache(1000, "lru")
	
	return &RegistrySet{
		CombinerRegistry: combinerRegistry,
		CapLayerRegistry: capLayerRegistry,
		PluginRegistry:   pluginRegistry,
		ConfigLoader:     configLoader,
		Cache:            cache,
	}, nil
}

// CreateAllRegistriesFromFiles creates all registry types from files
func (rf *RegistryFactory) CreateAllRegistriesFromFiles(combinerFile, layerFile string) (*RegistrySet, error) {
	combinerRegistry, err := rf.CreateCombinerRegistryFromFile(combinerFile)
	if err != nil {
		return nil, fmt.Errorf("failed to create combiner registry: %w", err)
	}
	
	capLayerRegistry, err := rf.CreateCapLayerRegistryFromFile(layerFile)
	if err != nil {
		return nil, fmt.Errorf("failed to create cap layer registry: %w", err)
	}
	
	pluginRegistry := rf.CreatePluginRegistry()
	configLoader := rf.CreateConfigLoader()
	cache := rf.CreateCache(1000, "lru")
	
	return &RegistrySet{
		CombinerRegistry: combinerRegistry,
		CapLayerRegistry: capLayerRegistry,
		PluginRegistry:   pluginRegistry,
		ConfigLoader:     configLoader,
		Cache:            cache,
	}, nil
}

// RegistrySet contains all registry instances
type RegistrySet struct {
	CombinerRegistry interfaces.CombinerRegistry
	CapLayerRegistry interfaces.CapLayerRegistry
	PluginRegistry   interfaces.PluginRegistry
	ConfigLoader     interfaces.ConfigLoader
	Cache            interfaces.Cache
}

// Validate validates all registries
func (rs *RegistrySet) Validate() error {
	if rs.CombinerRegistry == nil {
		return fmt.Errorf("combiner registry is nil")
	}
	
	if rs.CapLayerRegistry == nil {
		return fmt.Errorf("cap layer registry is nil")
	}
	
	if rs.PluginRegistry == nil {
		return fmt.Errorf("plugin registry is nil")
	}
	
	if rs.ConfigLoader == nil {
		return fmt.Errorf("config loader is nil")
	}
	
	if rs.Cache == nil {
		return fmt.Errorf("cache is nil")
	}
	
	// Validate individual registries
	if err := rs.CombinerRegistry.Validate(); err != nil {
		return fmt.Errorf("combiner registry validation failed: %w", err)
	}
	
	if err := rs.CapLayerRegistry.Validate(); err != nil {
		return fmt.Errorf("cap layer registry validation failed: %w", err)
	}
	
	// PluginRegistry doesn't have Validate method
	// if err := rs.PluginRegistry.Validate(); err != nil {
	// 	return fmt.Errorf("plugin registry validation failed: %w", err)
	// }
	
	return nil
}

// GetCombinerRegistry returns the combiner registry
func (rs *RegistrySet) GetCombinerRegistry() interfaces.CombinerRegistry {
	return rs.CombinerRegistry
}

// GetCapLayerRegistry returns the cap layer registry
func (rs *RegistrySet) GetCapLayerRegistry() interfaces.CapLayerRegistry {
	return rs.CapLayerRegistry
}

// GetPluginRegistry returns the plugin registry
func (rs *RegistrySet) GetPluginRegistry() interfaces.PluginRegistry {
	return rs.PluginRegistry
}

// GetConfigLoader returns the config loader
func (rs *RegistrySet) GetConfigLoader() interfaces.ConfigLoader {
	return rs.ConfigLoader
}

// GetCache returns the cache
func (rs *RegistrySet) GetCache() interfaces.Cache {
	return rs.Cache
}

// SetCombinerRegistry sets the combiner registry
func (rs *RegistrySet) SetCombinerRegistry(registry interfaces.CombinerRegistry) {
	rs.CombinerRegistry = registry
}

// SetCapLayerRegistry sets the cap layer registry
func (rs *RegistrySet) SetCapLayerRegistry(registry interfaces.CapLayerRegistry) {
	rs.CapLayerRegistry = registry
}

// SetPluginRegistry sets the plugin registry
func (rs *RegistrySet) SetPluginRegistry(registry interfaces.PluginRegistry) {
	rs.PluginRegistry = registry
}

// SetConfigLoader sets the config loader
func (rs *RegistrySet) SetConfigLoader(loader interfaces.ConfigLoader) {
	rs.ConfigLoader = loader
}

// SetCache sets the cache
func (rs *RegistrySet) SetCache(cache interfaces.Cache) {
	rs.Cache = cache
}

// Clear clears all registries
func (rs *RegistrySet) Clear() {
	// CombinerRegistry doesn't have Clear method
	// if rs.CombinerRegistry != nil {
	// 	rs.CombinerRegistry.Clear()
	// }
	
	if rs.PluginRegistry != nil {
		rs.PluginRegistry.Clear()
	}
	
	if rs.Cache != nil {
		rs.Cache.Clear()
	}
}

// GetStats returns statistics for all registries
func (rs *RegistrySet) GetStats() map[string]interface{} {
	stats := make(map[string]interface{})
	
	if rs.CombinerRegistry != nil {
		// CombinerRegistry doesn't have Count method
		// stats["combiner_rules"] = rs.CombinerRegistry.Count()
	}
	
	if rs.CapLayerRegistry != nil {
		// CapLayerRegistry doesn't have GetLayerCount method
		// stats["layer_count"] = rs.CapLayerRegistry.GetLayerCount()
		stats["across_policy"] = rs.CapLayerRegistry.GetAcrossLayerPolicy()
	}
	
	if rs.PluginRegistry != nil {
		stats["subsystem_count"] = rs.PluginRegistry.Count()
	}
	
	if rs.Cache != nil {
		stats["cache_stats"] = rs.Cache.GetStats()
	}
	
	return stats
}
