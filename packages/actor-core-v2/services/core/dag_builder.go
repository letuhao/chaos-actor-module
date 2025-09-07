package core

import (
	"fmt"
	"sync"
)

// DAGNode represents a node in the dependency graph
type DAGNode struct {
	Name         string
	Dependencies []string
	Visited      bool
	Visiting     bool
}

// DAGBuilder builds and manages dependency graphs
type DAGBuilder struct {
	nodes map[string]*DAGNode
	mu    sync.RWMutex
}

// NewDAGBuilder creates a new DAG builder
func NewDAGBuilder() *DAGBuilder {
	return &DAGBuilder{
		nodes: make(map[string]*DAGNode),
	}
}

// AddNode adds a node to the DAG
func (db *DAGBuilder) AddNode(name string, dependencies []string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.nodes[name] = &DAGNode{
		Name:         name,
		Dependencies: dependencies,
		Visited:      false,
		Visiting:     false,
	}
}

// BuildOrder performs topological sort to determine evaluation order
func (db *DAGBuilder) BuildOrder() ([]string, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	// Reset visit states
	for _, node := range db.nodes {
		node.Visited = false
		node.Visiting = false
	}

	var order []string
	var visit func(string) error

	visit = func(nodeName string) error {
		node, exists := db.nodes[nodeName]
		if !exists {
			// External dependency, skip
			return nil
		}

		if node.Visiting {
			return fmt.Errorf("circular dependency detected involving %s", nodeName)
		}

		if node.Visited {
			return nil
		}

		node.Visiting = true

		// Visit all dependencies first
		for _, dep := range node.Dependencies {
			if err := visit(dep); err != nil {
				return err
			}
		}

		node.Visiting = false
		node.Visited = true
		order = append(order, nodeName)

		return nil
	}

	// Visit all nodes
	for nodeName := range db.nodes {
		if err := visit(nodeName); err != nil {
			return nil, err
		}
	}

	return order, nil
}

// GetDependencies returns dependencies for a node
func (db *DAGBuilder) GetDependencies(nodeName string) ([]string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	node, exists := db.nodes[nodeName]
	if !exists {
		return nil, false
	}

	return node.Dependencies, true
}

// HasNode checks if a node exists
func (db *DAGBuilder) HasNode(nodeName string) bool {
	db.mu.RLock()
	defer db.mu.RUnlock()

	_, exists := db.nodes[nodeName]
	return exists
}

// GetNodeCount returns the number of nodes
func (db *DAGBuilder) GetNodeCount() int {
	db.mu.RLock()
	defer db.mu.RUnlock()

	return len(db.nodes)
}

// Clear removes all nodes
func (db *DAGBuilder) Clear() {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.nodes = make(map[string]*DAGNode)
}
