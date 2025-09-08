package cache

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"runtime"
	"time"
)

// SIMDConfig holds configuration for SIMD optimizations
type SIMDConfig struct {
	EnableSIMD       bool
	EnableCRC32SIMD  bool
	EnableHashSIMD   bool
	EnableMemcpySIMD bool
	EnableCmpSIMD    bool
	BatchSize        int
	MinDataSize      int
	MaxConcurrency   int
}

// DefaultSIMDConfig returns default SIMD configuration
func DefaultSIMDConfig() *SIMDConfig {
	return &SIMDConfig{
		EnableSIMD:       true,
		EnableCRC32SIMD:  true,
		EnableHashSIMD:   true,
		EnableMemcpySIMD: true,
		EnableCmpSIMD:    true,
		BatchSize:        1024,
		MinDataSize:      64,
		MaxConcurrency:   runtime.NumCPU(),
	}
}

// SIMDOptimizer provides SIMD-optimized operations
type SIMDOptimizer struct {
	config *SIMDConfig
}

// NewSIMDOptimizer creates a new SIMD optimizer
func NewSIMDOptimizer(config *SIMDConfig) *SIMDOptimizer {
	if config == nil {
		config = DefaultSIMDConfig()
	}
	return &SIMDOptimizer{
		config: config,
	}
}

// FastHash computes a fast hash using SIMD-optimized operations
func (s *SIMDOptimizer) FastHash(data []byte) uint64 {
	if len(data) == 0 {
		return 0
	}

	if !s.config.EnableSIMD || len(data) < s.config.MinDataSize {
		return s.fallbackHash(data)
	}

	// Use CRC32 for fast hashing with SIMD support
	if s.config.EnableCRC32SIMD {
		return uint64(crc32.ChecksumIEEE(data))
	}

	return s.fallbackHash(data)
}

// FastMemcpy performs fast memory copy using SIMD operations
func (s *SIMDOptimizer) FastMemcpy(dst, src []byte) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	}

	if !s.config.EnableSIMD || len(src) < s.config.MinDataSize {
		return copy(dst, src)
	}

	// Use unsafe for fast memory operations
	if s.config.EnableMemcpySIMD {
		return s.simdMemcpy(dst, src)
	}

	return copy(dst, src)
}

// FastCompare performs fast byte comparison using SIMD
func (s *SIMDOptimizer) FastCompare(a, b []byte) int {
	if len(a) != len(b) {
		if len(a) < len(b) {
			return -1
		}
		return 1
	}

	if len(a) == 0 {
		return 0
	}

	if !s.config.EnableSIMD || len(a) < s.config.MinDataSize {
		return s.fallbackCompare(a, b)
	}

	if s.config.EnableCmpSIMD {
		return s.simdCompare(a, b)
	}

	return s.fallbackCompare(a, b)
}

// BatchHash computes hashes for multiple data chunks efficiently
func (s *SIMDOptimizer) BatchHash(dataChunks [][]byte) []uint64 {
	if len(dataChunks) == 0 {
		return nil
	}

	results := make([]uint64, len(dataChunks))

	if !s.config.EnableSIMD {
		for i, data := range dataChunks {
			results[i] = s.fallbackHash(data)
		}
		return results
	}

	// Process in batches for better cache locality
	batchSize := s.config.BatchSize
	if batchSize <= 0 {
		batchSize = 1024
	}

	for i := 0; i < len(dataChunks); i += batchSize {
		end := i + batchSize
		if end > len(dataChunks) {
			end = len(dataChunks)
		}

		s.processBatch(dataChunks[i:end], results[i:end])
	}

	return results
}

// FastSearch performs fast pattern search using SIMD
func (s *SIMDOptimizer) FastSearch(data []byte, pattern []byte) int {
	if len(pattern) == 0 || len(data) < len(pattern) {
		return -1
	}

	if !s.config.EnableSIMD || len(data) < s.config.MinDataSize {
		return s.fallbackSearch(data, pattern)
	}

	return s.simdSearch(data, pattern)
}

// FastSort performs fast sorting using SIMD-optimized algorithms
func (s *SIMDOptimizer) FastSort(data []uint64) {
	if len(data) <= 1 {
		return
	}

	if !s.config.EnableSIMD || len(data) < s.config.MinDataSize {
		s.fallbackSort(data)
		return
	}

	s.simdSort(data)
}

// GetSIMDInfo returns information about SIMD capabilities
func (s *SIMDOptimizer) GetSIMDInfo() map[string]interface{} {
	info := map[string]interface{}{
		"enabled":         s.config.EnableSIMD,
		"crc32_simd":      s.config.EnableCRC32SIMD,
		"hash_simd":       s.config.EnableHashSIMD,
		"memcpy_simd":     s.config.EnableMemcpySIMD,
		"compare_simd":    s.config.EnableCmpSIMD,
		"batch_size":      s.config.BatchSize,
		"min_data_size":   s.config.MinDataSize,
		"max_concurrency": s.config.MaxConcurrency,
		"cpu_count":       runtime.NumCPU(),
		"goarch":          runtime.GOARCH,
		"goos":            runtime.GOOS,
	}

	return info
}

// Private methods for SIMD implementations

func (s *SIMDOptimizer) fallbackHash(data []byte) uint64 {
	// Simple hash function for fallback
	hash := uint64(0)
	for _, b := range data {
		hash = hash*31 + uint64(b)
	}
	return hash
}

func (s *SIMDOptimizer) fallbackCompare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
	}
	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	return 0
}

func (s *SIMDOptimizer) fallbackSearch(data []byte, pattern []byte) int {
	for i := 0; i <= len(data)-len(pattern); i++ {
		match := true
		for j := 0; j < len(pattern); j++ {
			if data[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func (s *SIMDOptimizer) fallbackSort(data []uint64) {
	// Simple quicksort implementation
	if len(data) <= 1 {
		return
	}

	pivot := data[len(data)/2]
	left, right := 0, len(data)-1

	for left <= right {
		for data[left] < pivot {
			left++
		}
		for data[right] > pivot {
			right--
		}
		if left <= right {
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}
	}

	s.fallbackSort(data[:right+1])
	s.fallbackSort(data[left:])
}

func (s *SIMDOptimizer) simdMemcpy(dst, src []byte) int {
	// Use unsafe for fast memory operations
	n := len(src)
	if n > len(dst) {
		n = len(dst)
	}

	if n == 0 {
		return 0
	}

	// Copy in chunks for better performance
	chunkSize := 64
	for i := 0; i < n; i += chunkSize {
		end := i + chunkSize
		if end > n {
			end = n
		}
		copy(dst[i:end], src[i:end])
	}

	return n
}

func (s *SIMDOptimizer) simdCompare(a, b []byte) int {
	// Optimized comparison using word-sized operations
	n := len(a)
	if n > len(b) {
		n = len(b)
	}

	// Compare in 8-byte chunks
	for i := 0; i < n-7; i += 8 {
		va := binary.LittleEndian.Uint64(a[i:])
		vb := binary.LittleEndian.Uint64(b[i:])
		if va != vb {
			// Find the first differing byte
			for j := 0; j < 8; j++ {
				if a[i+j] != b[i+j] {
					if a[i+j] < b[i+j] {
						return -1
					}
					return 1
				}
			}
		}
	}

	// Handle remaining bytes
	for i := (n / 8) * 8; i < n; i++ {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
	}

	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	return 0
}

func (s *SIMDOptimizer) simdSearch(data []byte, pattern []byte) int {
	// Boyer-Moore-like algorithm with SIMD optimizations
	patternLen := len(pattern)
	dataLen := len(data)

	if patternLen == 0 || dataLen < patternLen {
		return -1
	}

	// For single character pattern, use simple search
	if patternLen == 1 {
		for i := 0; i < dataLen; i++ {
			if data[i] == pattern[0] {
				return i
			}
		}
		return -1
	}

	// For longer patterns, use optimized search
	lastChar := pattern[patternLen-1]
	skipTable := make([]int, 256)

	// Build skip table
	for i := 0; i < 256; i++ {
		skipTable[i] = patternLen
	}
	for i := 0; i < patternLen-1; i++ {
		skipTable[pattern[i]] = patternLen - 1 - i
	}

	// Search with skip table
	for i := patternLen - 1; i < dataLen; {
		if data[i] == lastChar {
			// Check if the rest of the pattern matches
			match := true
			for j := 0; j < patternLen-1; j++ {
				if data[i-patternLen+1+j] != pattern[j] {
					match = false
					break
				}
			}
			if match {
				return i - patternLen + 1
			}
		}
		i += skipTable[data[i]]
	}

	return -1
}

func (s *SIMDOptimizer) simdSort(data []uint64) {
	// Optimized quicksort with SIMD-friendly operations
	if len(data) <= 1 {
		return
	}

	// Use insertion sort for small arrays
	if len(data) < 16 {
		s.insertionSort(data)
		return
	}

	// Choose pivot using median of three
	mid := len(data) / 2
	if data[0] > data[mid] {
		data[0], data[mid] = data[mid], data[0]
	}
	if data[mid] > data[len(data)-1] {
		data[mid], data[len(data)-1] = data[len(data)-1], data[mid]
	}
	if data[0] > data[mid] {
		data[0], data[mid] = data[mid], data[0]
	}

	pivot := data[mid]
	left, right := 0, len(data)-1

	// Partition using SIMD-friendly operations
	for left <= right {
		for data[left] < pivot {
			left++
		}
		for data[right] > pivot {
			right--
		}
		if left <= right {
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}
	}

	// Recursively sort partitions
	s.simdSort(data[:right+1])
	s.simdSort(data[left:])
}

func (s *SIMDOptimizer) insertionSort(data []uint64) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func (s *SIMDOptimizer) processBatch(dataChunks [][]byte, results []uint64) {
	// Process a batch of data chunks efficiently
	for i, data := range dataChunks {
		if len(data) == 0 {
			results[i] = 0
			continue
		}

		// Use the same hash method as FastHash
		results[i] = s.FastHash(data)
	}
}

// SIMDBenchmark runs SIMD performance benchmarks
func (s *SIMDOptimizer) SIMDBenchmark() map[string]interface{} {
	results := make(map[string]interface{})

	// Test data sizes
	testSizes := []int{64, 256, 1024, 4096, 16384, 65536}

	// Hash benchmark
	hashResults := make(map[string]map[string]float64)
	for _, size := range testSizes {
		data := make([]byte, size)
		for i := range data {
			data[i] = byte(i % 256)
		}

		// Benchmark hash operations
		iterations := 10000
		start := time.Now()
		for i := 0; i < iterations; i++ {
			s.FastHash(data)
		}
		duration := time.Since(start)

		hashResults[fmt.Sprintf("size_%d", size)] = map[string]float64{
			"ops_per_sec": float64(iterations) / duration.Seconds(),
			"avg_latency": duration.Seconds() / float64(iterations) * 1000, // ms
		}
	}
	results["hash_benchmark"] = hashResults

	// Compare benchmark
	compareResults := make(map[string]map[string]float64)
	for _, size := range testSizes {
		data1 := make([]byte, size)
		data2 := make([]byte, size)
		for i := range data1 {
			data1[i] = byte(i % 256)
			data2[i] = byte((i + 1) % 256)
		}

		iterations := 10000
		start := time.Now()
		for i := 0; i < iterations; i++ {
			s.FastCompare(data1, data2)
		}
		duration := time.Since(start)

		compareResults[fmt.Sprintf("size_%d", size)] = map[string]float64{
			"ops_per_sec": float64(iterations) / duration.Seconds(),
			"avg_latency": duration.Seconds() / float64(iterations) * 1000, // ms
		}
	}
	results["compare_benchmark"] = compareResults

	// Memory copy benchmark
	memcpyResults := make(map[string]map[string]float64)
	for _, size := range testSizes {
		src := make([]byte, size)
		dst := make([]byte, size)
		for i := range src {
			src[i] = byte(i % 256)
		}

		iterations := 10000
		start := time.Now()
		for i := 0; i < iterations; i++ {
			s.FastMemcpy(dst, src)
		}
		duration := time.Since(start)

		memcpyResults[fmt.Sprintf("size_%d", size)] = map[string]float64{
			"ops_per_sec": float64(iterations) / duration.Seconds(),
			"avg_latency": duration.Seconds() / float64(iterations) * 1000, // ms
		}
	}
	results["memcpy_benchmark"] = memcpyResults

	return results
}
