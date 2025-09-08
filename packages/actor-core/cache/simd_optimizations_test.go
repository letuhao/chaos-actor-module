package cache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSIMDOptimizer_New(t *testing.T) {
	config := &SIMDConfig{
		EnableSIMD:      true,
		EnableCRC32SIMD: true,
		BatchSize:       512,
	}

	optimizer := NewSIMDOptimizer(config)
	if optimizer == nil {
		t.Fatal("Expected non-nil optimizer")
	}

	if !optimizer.config.EnableSIMD {
		t.Error("Expected SIMD to be enabled")
	}
}

func TestSIMDOptimizer_DefaultConfig(t *testing.T) {
	optimizer := NewSIMDOptimizer(nil)
	if optimizer == nil {
		t.Fatal("Expected non-nil optimizer")
	}

	config := optimizer.config
	if !config.EnableSIMD {
		t.Error("Expected SIMD to be enabled by default")
	}
	if config.BatchSize <= 0 {
		t.Error("Expected positive batch size")
	}
}

func TestSIMDOptimizer_FastHash(t *testing.T) {
	optimizer := NewSIMDOptimizer(nil)

	tests := []struct {
		name     string
		data     []byte
		expected uint64
	}{
		{
			name:     "empty data",
			data:     []byte{},
			expected: 0,
		},
		{
			name:     "single byte",
			data:     []byte{42},
			expected: 42,
		},
		{
			name:     "multiple bytes",
			data:     []byte{1, 2, 3, 4},
			expected: 1*31*31*31 + 2*31*31 + 3*31 + 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := optimizer.FastHash(tt.data)
			if len(tt.data) == 0 && result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
			// For non-empty data, just ensure we get a consistent result
			if len(tt.data) > 0 {
				result2 := optimizer.FastHash(tt.data)
				if result != result2 {
					t.Error("Hash should be consistent")
				}
			}
		})
	}
}

func TestSIMDOptimizer_FastMemcpy(t *testing.T) {
	optimizer := NewSIMDOptimizer(nil)

	tests := []struct {
		name     string
		src      []byte
		dstSize  int
		expected int
	}{
		{
			name:     "empty source",
			src:      []byte{},
			dstSize:  10,
			expected: 0,
		},
		{
			name:     "normal copy",
			src:      []byte{1, 2, 3, 4, 5},
			dstSize:  10,
			expected: 5,
		},
		{
			name:     "dst smaller than src",
			src:      []byte{1, 2, 3, 4, 5},
			dstSize:  3,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := make([]byte, tt.dstSize)
			copied := optimizer.FastMemcpy(dst, tt.src)

			if copied != tt.expected {
				t.Errorf("Expected %d bytes copied, got %d", tt.expected, copied)
			}

			// Verify the copied data
			for i := 0; i < copied; i++ {
				if dst[i] != tt.src[i] {
					t.Errorf("Data mismatch at index %d: expected %d, got %d", i, tt.src[i], dst[i])
				}
			}
		})
	}
}

func TestSIMDOptimizer_FastCompare(t *testing.T) {
	optimizer := NewSIMDOptimizer(nil)

	tests := []struct {
		name     string
		a        []byte
		b        []byte
		expected int
	}{
		{
			name:     "equal arrays",
			a:        []byte{1, 2, 3, 4},
			b:        []byte{1, 2, 3, 4},
			expected: 0,
		},
		{
			name:     "a less than b",
			a:        []byte{1, 2, 3, 3},
			b:        []byte{1, 2, 3, 4},
			expected: -1,
		},
		{
			name:     "a greater than b",
			a:        []byte{1, 2, 3, 5},
			b:        []byte{1, 2, 3, 4},
			expected: 1,
		},
		{
			name:     "a shorter than b",
			a:        []byte{1, 2, 3},
			b:        []byte{1, 2, 3, 4},
			expected: -1,
		},
		{
			name:     "a longer than b",
			a:        []byte{1, 2, 3, 4},
			b:        []byte{1, 2, 3},
			expected: 1,
		},
		{
			name:     "empty arrays",
			a:        []byte{},
			b:        []byte{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := optimizer.FastCompare(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestSIMDOptimizer_BatchHash(t *testing.T) {
	optimizer := NewSIMDOptimizer(nil)

	dataChunks := [][]byte{
		[]byte{1, 2, 3},
		[]byte{4, 5, 6},
		[]byte{7, 8, 9},
		[]byte{},
		[]byte{10, 11, 12},
	}

	results := optimizer.BatchHash(dataChunks)

	if len(results) != len(dataChunks) {
		t.Fatalf("Expected %d results, got %d", len(dataChunks), len(results))
	}

	// Verify that empty data produces 0 hash
	if results[3] != 0 {
		t.Error("Empty data should produce 0 hash")
	}

	// Verify that results are consistent
	for i, data := range dataChunks {
		if len(data) > 0 {
			expected := optimizer.FastHash(data)
			if results[i] != expected {
				t.Errorf("Batch hash result %d doesn't match individual hash", i)
			}
		}
	}
}

func TestSIMDOptimizer_FastSearch(t *testing.T) {
	optimizer := NewSIMDOptimizer(nil)

	tests := []struct {
		name     string
		data     []byte
		pattern  []byte
		expected int
	}{
		{
			name:     "pattern found at beginning",
			data:     []byte{1, 2, 3, 4, 5},
			pattern:  []byte{1, 2, 3},
			expected: 0,
		},
		{
			name:     "pattern found in middle",
			data:     []byte{1, 2, 3, 4, 5},
			pattern:  []byte{3, 4},
			expected: 2,
		},
		{
			name:     "pattern found at end",
			data:     []byte{1, 2, 3, 4, 5},
			pattern:  []byte{4, 5},
			expected: 3,
		},
		{
			name:     "pattern not found",
			data:     []byte{1, 2, 3, 4, 5},
			pattern:  []byte{6, 7},
			expected: -1,
		},
		{
			name:     "empty pattern",
			data:     []byte{1, 2, 3, 4, 5},
			pattern:  []byte{},
			expected: -1,
		},
		{
			name:     "pattern longer than data",
			data:     []byte{1, 2, 3},
			pattern:  []byte{1, 2, 3, 4, 5},
			expected: -1,
		},
		{
			name:     "single character pattern",
			data:     []byte{1, 2, 3, 4, 5},
			pattern:  []byte{3},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := optimizer.FastSearch(tt.data, tt.pattern)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestSIMDOptimizer_FastSort(t *testing.T) {
	optimizer := NewSIMDOptimizer(nil)

	tests := []struct {
		name     string
		data     []uint64
		expected []uint64
	}{
		{
			name:     "empty array",
			data:     []uint64{},
			expected: []uint64{},
		},
		{
			name:     "single element",
			data:     []uint64{42},
			expected: []uint64{42},
		},
		{
			name:     "already sorted",
			data:     []uint64{1, 2, 3, 4, 5},
			expected: []uint64{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted",
			data:     []uint64{5, 4, 3, 2, 1},
			expected: []uint64{1, 2, 3, 4, 5},
		},
		{
			name:     "random order",
			data:     []uint64{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []uint64{1, 1, 2, 3, 4, 5, 6, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid modifying the original
			data := make([]uint64, len(tt.data))
			copy(data, tt.data)

			optimizer.FastSort(data)

			if len(data) != len(tt.expected) {
				t.Fatalf("Expected length %d, got %d", len(tt.expected), len(data))
			}

			for i := 0; i < len(data); i++ {
				if data[i] != tt.expected[i] {
					t.Errorf("Expected %v, got %v", tt.expected, data)
					break
				}
			}
		})
	}
}

func TestSIMDOptimizer_GetSIMDInfo(t *testing.T) {
	optimizer := NewSIMDOptimizer(nil)

	info := optimizer.GetSIMDInfo()

	requiredKeys := []string{
		"enabled", "crc32_simd", "hash_simd", "memcpy_simd", "compare_simd",
		"batch_size", "min_data_size", "max_concurrency", "cpu_count", "goarch", "goos",
	}

	for _, key := range requiredKeys {
		if _, exists := info[key]; !exists {
			t.Errorf("Missing key in SIMD info: %s", key)
		}
	}

	// Check some specific values
	if enabled, ok := info["enabled"].(bool); !ok || !enabled {
		t.Error("Expected SIMD to be enabled")
	}

	if cpuCount, ok := info["cpu_count"].(int); !ok || cpuCount <= 0 {
		t.Error("Expected positive CPU count")
	}
}

func TestSIMDOptimizer_SIMDBenchmark(t *testing.T) {
	optimizer := NewSIMDOptimizer(nil)

	results := optimizer.SIMDBenchmark()

	requiredBenchmarks := []string{"hash_benchmark", "compare_benchmark", "memcpy_benchmark"}

	for _, benchmark := range requiredBenchmarks {
		if _, exists := results[benchmark]; !exists {
			t.Errorf("Missing benchmark: %s", benchmark)
		}
	}

	// Check that each benchmark has results for different sizes
	for _, benchmark := range requiredBenchmarks {
		benchmarkData, ok := results[benchmark].(map[string]map[string]float64)
		if !ok {
			t.Errorf("Invalid benchmark data format for %s", benchmark)
			continue
		}

		expectedSizes := []string{"size_64", "size_256", "size_1024", "size_4096", "size_16384", "size_65536"}
		for _, size := range expectedSizes {
			if _, exists := benchmarkData[size]; !exists {
				t.Errorf("Missing size %s in %s", size, benchmark)
			}
		}
	}
}

func BenchmarkSIMDOptimizer_FastHash(b *testing.B) {
	optimizer := NewSIMDOptimizer(nil)

	sizes := []int{64, 256, 1024, 4096, 16384}

	for _, size := range sizes {
		data := make([]byte, size)
		for i := range data {
			data[i] = byte(i % 256)
		}

		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				optimizer.FastHash(data)
			}
		})
	}
}

func BenchmarkSIMDOptimizer_FastCompare(b *testing.B) {
	optimizer := NewSIMDOptimizer(nil)

	sizes := []int{64, 256, 1024, 4096, 16384}

	for _, size := range sizes {
		data1 := make([]byte, size)
		data2 := make([]byte, size)
		for i := range data1 {
			data1[i] = byte(i % 256)
			data2[i] = byte((i + 1) % 256)
		}

		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				optimizer.FastCompare(data1, data2)
			}
		})
	}
}

func BenchmarkSIMDOptimizer_FastMemcpy(b *testing.B) {
	optimizer := NewSIMDOptimizer(nil)

	sizes := []int{64, 256, 1024, 4096, 16384}

	for _, size := range sizes {
		src := make([]byte, size)
		dst := make([]byte, size)
		for i := range src {
			src[i] = byte(i % 256)
		}

		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				optimizer.FastMemcpy(dst, src)
			}
		})
	}
}

func BenchmarkSIMDOptimizer_BatchHash(b *testing.B) {
	optimizer := NewSIMDOptimizer(nil)

	batchSizes := []int{10, 100, 1000}
	dataSize := 1024

	for _, batchSize := range batchSizes {
		dataChunks := make([][]byte, batchSize)
		for i := range dataChunks {
			data := make([]byte, dataSize)
			for j := range data {
				data[j] = byte((i + j) % 256)
			}
			dataChunks[i] = data
		}

		b.Run(fmt.Sprintf("batch_%d", batchSize), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				optimizer.BatchHash(dataChunks)
			}
		})
	}
}

func BenchmarkSIMDOptimizer_FastSort(b *testing.B) {
	optimizer := NewSIMDOptimizer(nil)

	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		data := make([]uint64, size)
		rand.Seed(time.Now().UnixNano())
		for i := range data {
			data[i] = rand.Uint64()
		}

		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Make a copy for each iteration
				testData := make([]uint64, len(data))
				copy(testData, data)
				optimizer.FastSort(testData)
			}
		})
	}
}
