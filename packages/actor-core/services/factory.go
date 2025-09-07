package services

import (
	"chaos-actor-module/packages/actor-core/interfaces"
	"chaos-actor-module/packages/actor-core/registry"
	"fmt"
)

// ServiceFactory creates service instances
type ServiceFactory struct{}

// NewServiceFactory creates a new service factory
func NewServiceFactory() *ServiceFactory {
	return &ServiceFactory{}
}

// CreateCapsProvider creates a new caps provider
func (sf *ServiceFactory) CreateCapsProvider(layerRegistry interfaces.CapLayerRegistry) interfaces.CapsProvider {
	return NewCapsProvider(layerRegistry)
}

// CreateAggregator creates a new aggregator
func (sf *ServiceFactory) CreateAggregator(
	combinerRegistry interfaces.CombinerRegistry,
	capsProvider interfaces.CapsProvider,
	pluginRegistry interfaces.PluginRegistry,
	cache interfaces.Cache,
) interfaces.Aggregator {
	return NewAggregator(combinerRegistry, capsProvider, pluginRegistry, cache)
}

// CreateAllServices creates all services
func (sf *ServiceFactory) CreateAllServices() (*ServiceSet, error) {
	// Create registries
	registryFactory := registry.NewRegistryFactory()
	registrySet, err := registryFactory.CreateAllRegistries()
	if err != nil {
		return nil, fmt.Errorf("failed to create registries: %w", err)
	}

	// Create services
	capsProvider := sf.CreateCapsProvider(registrySet.GetCapLayerRegistry())
	aggregator := sf.CreateAggregator(
		registrySet.GetCombinerRegistry(),
		capsProvider,
		registrySet.GetPluginRegistry(),
		registrySet.GetCache(),
	)

	return &ServiceSet{
		CapsProvider: capsProvider,
		Aggregator:   aggregator,
		Registries:   registrySet,
	}, nil
}

// CreateAllServicesFromRegistries creates all services from existing registries
func (sf *ServiceFactory) CreateAllServicesFromRegistries(registrySet *registry.RegistrySet) (*ServiceSet, error) {
	if registrySet == nil {
		return nil, fmt.Errorf("registry set cannot be nil")
	}

	// Create services
	capsProvider := sf.CreateCapsProvider(registrySet.GetCapLayerRegistry())
	aggregator := sf.CreateAggregator(
		registrySet.GetCombinerRegistry(),
		capsProvider,
		registrySet.GetPluginRegistry(),
		registrySet.GetCache(),
	)

	return &ServiceSet{
		CapsProvider: capsProvider,
		Aggregator:   aggregator,
		Registries:   registrySet,
	}, nil
}

// ServiceSet contains all service instances
type ServiceSet struct {
	CapsProvider interfaces.CapsProvider
	Aggregator   interfaces.Aggregator
	Registries   *registry.RegistrySet
}

// Validate validates all services
func (ss *ServiceSet) Validate() error {
	if ss.CapsProvider == nil {
		return fmt.Errorf("caps provider is nil")
	}

	if ss.Aggregator == nil {
		return fmt.Errorf("aggregator is nil")
	}

	if ss.Registries == nil {
		return fmt.Errorf("registries is nil")
	}

	// Validate individual services
	// CapsProvider and Aggregator don't have Validate methods
	// if err := ss.CapsProvider.Validate(); err != nil {
	// 	return fmt.Errorf("caps provider validation failed: %w", err)
	// }

	// if err := ss.Aggregator.Validate(); err != nil {
	// 	return fmt.Errorf("aggregator validation failed: %w", err)
	// }

	if err := ss.Registries.Validate(); err != nil {
		return fmt.Errorf("registries validation failed: %w", err)
	}

	return nil
}

// GetCapsProvider returns the caps provider
func (ss *ServiceSet) GetCapsProvider() interfaces.CapsProvider {
	return ss.CapsProvider
}

// GetAggregator returns the aggregator
func (ss *ServiceSet) GetAggregator() interfaces.Aggregator {
	return ss.Aggregator
}

// GetRegistries returns the registries
func (ss *ServiceSet) GetRegistries() *registry.RegistrySet {
	return ss.Registries
}

// SetCapsProvider sets the caps provider
func (ss *ServiceSet) SetCapsProvider(provider interfaces.CapsProvider) {
	ss.CapsProvider = provider
}

// SetAggregator sets the aggregator
func (ss *ServiceSet) SetAggregator(aggregator interfaces.Aggregator) {
	ss.Aggregator = aggregator
}

// SetRegistries sets the registries
func (ss *ServiceSet) SetRegistries(registries *registry.RegistrySet) {
	ss.Registries = registries
}

// GetStatistics returns statistics for all services
func (ss *ServiceSet) GetStatistics() map[string]interface{} {
	stats := make(map[string]interface{})

	if ss.CapsProvider != nil {
		// Get caps provider statistics
		// This would require a method to get stats from caps provider
		stats["caps_provider"] = "available"
	}

	if ss.Aggregator != nil {
		stats["aggregator"] = ss.Aggregator.GetMetrics()
	}

	if ss.Registries != nil {
		stats["registries"] = ss.Registries.GetStats()
	}

	return stats
}

// Clear clears all services
func (ss *ServiceSet) Clear() {
	if ss.Registries != nil {
		ss.Registries.Clear()
	}

	if ss.Aggregator != nil {
		ss.Aggregator.ClearCache()
	}
}

// GetServiceStatus returns the status of all services
func (ss *ServiceSet) GetServiceStatus() map[string]string {
	status := make(map[string]string)

	if ss.CapsProvider != nil {
		status["caps_provider"] = "active"
	} else {
		status["caps_provider"] = "inactive"
	}

	if ss.Aggregator != nil {
		status["aggregator"] = "active"
	} else {
		status["aggregator"] = "inactive"
	}

	if ss.Registries != nil {
		status["registries"] = "active"
	} else {
		status["registries"] = "inactive"
	}

	return status
}

// IsHealthy checks if all services are healthy
func (ss *ServiceSet) IsHealthy() bool {
	if ss.CapsProvider == nil || ss.Aggregator == nil || ss.Registries == nil {
		return false
	}

	// Check if services can be validated
	if err := ss.Validate(); err != nil {
		return false
	}

	return true
}

// GetHealthReport returns a detailed health report
func (ss *ServiceSet) GetHealthReport() map[string]interface{} {
	report := make(map[string]interface{})

	report["healthy"] = ss.IsHealthy()
	report["status"] = ss.GetServiceStatus()
	report["statistics"] = ss.GetStatistics()

	// Add validation errors if any
	if err := ss.Validate(); err != nil {
		report["validation_error"] = err.Error()
	}

	return report
}
