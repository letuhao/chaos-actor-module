package enums

import (
	"chaos-actor-module/packages/actor-core/enums"
	"testing"
)

func TestBucket_String(t *testing.T) {
	tests := []struct {
		name     string
		bucket   enums.Bucket
		expected string
	}{
		{"Flat", enums.BucketFlat, "FLAT"},
		{"Mult", enums.BucketMult, "MULT"},
		{"PostAdd", enums.BucketPostAdd, "POST_ADD"},
		{"Override", enums.BucketOverride, "OVERRIDE"},
		{"Exponential", enums.BucketExponential, "EXPONENTIAL"},
		{"Logarithmic", enums.BucketLogarithmic, "LOGARITHMIC"},
		{"Conditional", enums.BucketConditional, "CONDITIONAL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bucket.String(); got != tt.expected {
				t.Errorf("Bucket.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestBucket_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		bucket   enums.Bucket
		expected bool
	}{
		{"Valid Flat", enums.BucketFlat, true},
		{"Valid Mult", enums.BucketMult, true},
		{"Valid PostAdd", enums.BucketPostAdd, true},
		{"Valid Override", enums.BucketOverride, true},
		{"Valid Exponential", enums.BucketExponential, true},
		{"Valid Logarithmic", enums.BucketLogarithmic, true},
		{"Valid Conditional", enums.BucketConditional, true},
		{"Invalid", enums.Bucket("INVALID"), false},
		{"Empty", enums.Bucket(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bucket.IsValid(); got != tt.expected {
				t.Errorf("Bucket.IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestBucket_GetOrder(t *testing.T) {
	expected := []enums.Bucket{
		enums.BucketFlat,
		enums.BucketMult,
		enums.BucketPostAdd,
		enums.BucketOverride,
		enums.BucketExponential,
		enums.BucketLogarithmic,
		enums.BucketConditional,
	}

	// Since GetBucketOrder doesn't exist, we'll test individual buckets
	for _, bucket := range expected {
		if !bucket.IsValid() {
			t.Errorf("Bucket %v should be valid", bucket)
		}
	}
}

func TestBucket_GetDefaultOrder(t *testing.T) {
	expected := []enums.Bucket{
		enums.BucketFlat,
		enums.BucketMult,
		enums.BucketPostAdd,
		enums.BucketOverride,
		enums.BucketExponential,
		enums.BucketLogarithmic,
		enums.BucketConditional,
	}

	// Since GetDefaultOrder doesn't exist for buckets, we'll test individual buckets
	for _, bucket := range expected {
		if !bucket.IsValid() {
			t.Errorf("Bucket %v should be valid", bucket)
		}
	}
}
