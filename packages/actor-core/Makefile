# Actor Core v2.0 - Makefile

.PHONY: help test-redis test-redis-integration test-redis-script clean-redis docker-redis docker-redis-stop

# Default target
help:
	@echo "Actor Core v2.0 - Available Commands:"
	@echo "====================================="
	@echo "test-redis              - Run Redis integration tests"
	@echo "test-redis-integration  - Run Redis integration tests (verbose)"
	@echo "test-redis-script       - Run Redis test script"
	@echo "docker-redis            - Start Redis with Docker"
	@echo "docker-redis-stop       - Stop Redis Docker container"
	@echo "clean-redis             - Clean Redis test data"
	@echo "help                    - Show this help message"

# Test Redis connection and basic operations
test-redis:
	@echo "🚀 Testing Redis (Memurai) Connection..."
	@go test ./tests/integration/redis_integration_test.go -v -run TestRedisConnection

# Test Redis integration tests
test-redis-integration:
	@echo "🚀 Running Redis Integration Tests..."
	@go test ./tests/integration/redis_integration_test.go -v

# Run Redis test script
test-redis-script:
	@echo "🚀 Running Redis Test Script..."
	@go run ./scripts/test_redis.go

# Clean Redis test data
clean-redis:
	@echo "🧹 Cleaning Redis test data..."
	@go run -c 'package main; import "actor-core-v2/tests/integration"; func main() { helper := integration.NewRedisTestHelper(); helper.ClearAll() }'

# Test all cache systems
test-cache:
	@echo "🚀 Testing All Cache Systems..."
	@go test ./tests/unit/mem_cache_test.go -v
	@go test ./tests/unit/state_tracker_test.go -v
	@go test ./tests/integration/redis_integration_test.go -v

# Build all
build:
	@echo "🔨 Building Actor Core v2.0..."
	@go build ./...

# Run all tests
test:
	@echo "🧪 Running All Tests..."
	@go test ./... -v

# Start Redis with Docker
docker-redis:
	@echo "🐳 Starting Redis with Docker..."
	@docker-compose up -d redis
	@echo "✅ Redis started! Run 'make test-redis-script' to test"

# Stop Redis Docker container
docker-redis-stop:
	@echo "🐳 Stopping Redis Docker container..."
	@docker-compose down
	@echo "✅ Redis stopped!"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	@go clean -cache
	@go mod tidy
