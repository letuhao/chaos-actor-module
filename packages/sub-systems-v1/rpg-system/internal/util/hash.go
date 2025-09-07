package util

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sort"
)

// BuildHash creates a deterministic hash from a compute input for caching
func BuildHash(input interface{}) (string, error) {
	// Convert to JSON for consistent serialization
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return "", fmt.Errorf("failed to marshal input: %w", err)
	}

	// Create SHA-256 hash
	hash := sha256.Sum256(jsonBytes)
	return fmt.Sprintf("%x", hash), nil
}

// BuildModifierHash creates a deterministic hash from a list of modifiers
func BuildModifierHash(modifiers []interface{}) (string, error) {
	// Sort modifiers by key and source for consistent ordering
	sortedModifiers := make([]interface{}, len(modifiers))
	copy(sortedModifiers, modifiers)

	// Sort by a combination of key and source for deterministic ordering
	sort.Slice(sortedModifiers, func(i, j int) bool {
		// This is a simplified sort - in practice you'd want to sort by actual fields
		// For now, we'll just use the JSON representation
		iBytes, _ := json.Marshal(sortedModifiers[i])
		jBytes, _ := json.Marshal(sortedModifiers[j])
		return string(iBytes) < string(jBytes)
	})

	// Convert to JSON
	jsonBytes, err := json.Marshal(sortedModifiers)
	if err != nil {
		return "", fmt.Errorf("failed to marshal modifiers: %w", err)
	}

	// Create SHA-256 hash
	hash := sha256.Sum256(jsonBytes)
	return fmt.Sprintf("%x", hash), nil
}

// BuildStatHash creates a deterministic hash from stat allocations
func BuildStatHash(allocations map[string]int) (string, error) {
	// Convert map to sorted slice for consistent ordering
	type statEntry struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	}

	var entries []statEntry
	for key, value := range allocations {
		entries = append(entries, statEntry{Key: key, Value: value})
	}

	// Sort by key
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Key < entries[j].Key
	})

	// Convert to JSON
	jsonBytes, err := json.Marshal(entries)
	if err != nil {
		return "", fmt.Errorf("failed to marshal stat allocations: %w", err)
	}

	// Create SHA-256 hash
	hash := sha256.Sum256(jsonBytes)
	return fmt.Sprintf("%x", hash), nil
}

// BuildSnapshotHash creates a deterministic hash from a stat snapshot
func BuildSnapshotHash(snapshot interface{}) (string, error) {
	// Convert to JSON
	jsonBytes, err := json.Marshal(snapshot)
	if err != nil {
		return "", fmt.Errorf("failed to marshal snapshot: %w", err)
	}

	// Create SHA-256 hash
	hash := sha256.Sum256(jsonBytes)
	return fmt.Sprintf("%x", hash), nil
}

// StableStringify creates a stable string representation of any value
func StableStringify(value interface{}) (string, error) {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("failed to marshal value: %w", err)
	}
	return string(jsonBytes), nil
}
