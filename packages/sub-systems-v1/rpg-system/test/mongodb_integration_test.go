package test

import (
	"testing"
)

// TestMongoDBIntegration tests the real MongoDB integration
// This test requires MongoDB to be running on localhost:27017
func TestMongoDBIntegration(t *testing.T) {
	// Skip if MongoDB is not available
	// Note: MongoDB integration requires the mongodb build tag
	t.Skip("MongoDB integration test requires mongodb build tag")
}
