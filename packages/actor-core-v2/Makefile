.PHONY: build test lint race bench clean install-tools ci

# Default target
all: build test lint race

# Build the project
build:
	@echo "🔨 Building..."
	go build ./...

# Run tests
test:
	@echo "🧪 Running tests..."
	go test ./tests/... -v

# Run tests with race detection
race:
	@echo "🏃 Running race detection..."
	go test -race ./tests/...

# Run benchmarks
bench:
	@echo "⚡ Running benchmarks..."
	go test -bench=. -benchmem ./tests/benchmark/...

# Run benchmarks and save results
bench-save:
	@echo "⚡ Running benchmarks and saving results..."
	go test -bench=. -benchmem ./tests/benchmark/... > benchmark_results.txt
	@echo "📊 Results saved to benchmark_results.txt"

# Run benchmarks with comparison
bench-compare:
	@echo "⚡ Running benchmarks with comparison..."
	go test -bench=. -benchmem ./tests/benchmark/... > benchmark_new.txt
	@if [ -f benchmark_old.txt ]; then \
		benchstat benchmark_old.txt benchmark_new.txt; \
	else \
		benchstat benchmark_new.txt; \
	fi

# Run linting
lint:
	@echo "🔍 Running linter..."
	golangci-lint run

# Install required tools
install-tools:
	@echo "📦 Installing tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install gotest.tools/gotestsum@latest
	go install golang.org/x/perf/cmd/benchstat@latest

# Run CI pipeline
ci: build lint test race bench

# Clean generated files
clean:
	@echo "🧹 Cleaning..."
	go clean
	rm -f benchmark_results.txt benchmark_old.txt benchmark_new.txt

# Run specific benchmark
bench-stat-resolver:
	@echo "⚡ Running StatResolver benchmarks..."
	go test -bench=BenchmarkStatResolver -benchmem ./tests/benchmark/...

bench-cache:
	@echo "⚡ Running Cache benchmarks..."
	go test -bench=BenchmarkMemCache -benchmem ./tests/benchmark/...

bench-performance-monitor:
	@echo "⚡ Running PerformanceMonitor benchmarks..."
	go test -bench=BenchmarkPerformanceMonitor -benchmem ./tests/benchmark/...

bench-formula-pipeline:
	@echo "⚡ Running FormulaPipeline benchmarks..."
	go test -bench=BenchmarkFormulaPipeline -benchmem ./tests/benchmark/...

# Run tests with coverage
test-coverage:
	@echo "🧪 Running tests with coverage..."
	go test -coverprofile=coverage.out ./tests/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "📊 Coverage report generated: coverage.html"

# Run specific test package
test-unit:
	@echo "🧪 Running unit tests..."
	go test ./tests/unit/... -v

test-golden:
	@echo "🧪 Running golden tests..."
	go test ./tests/golden/... -v

test-property:
	@echo "🧪 Running property tests..."
	go test ./tests/property/... -v

test-race:
	@echo "🧪 Running race tests..."
	go test ./tests/race/... -v

test-parity:
	@echo "🧪 Running parity tests..."
	go test ./tests/parity/... -v

# Help
help:
	@echo "Available targets:"
	@echo "  build          - Build the project"
	@echo "  test           - Run all tests"
	@echo "  race           - Run tests with race detection"
	@echo "  bench          - Run benchmarks"
	@echo "  bench-save     - Run benchmarks and save results"
	@echo "  bench-compare  - Run benchmarks with comparison"
	@echo "  lint           - Run linter"
	@echo "  install-tools  - Install required tools"
	@echo "  ci             - Run full CI pipeline"
	@echo "  clean          - Clean generated files"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  test-unit      - Run unit tests only"
	@echo "  test-golden    - Run golden tests only"
	@echo "  test-property  - Run property tests only"
	@echo "  test-race      - Run race tests only"
	@echo "  test-parity    - Run parity tests only"
	@echo "  bench-stat-resolver     - Run StatResolver benchmarks"
	@echo "  bench-cache            - Run Cache benchmarks"
	@echo "  bench-performance-monitor - Run PerformanceMonitor benchmarks"
	@echo "  bench-formula-pipeline - Run FormulaPipeline benchmarks"
	@echo "  help           - Show this help message"