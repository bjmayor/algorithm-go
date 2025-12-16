package main

import (
	"reflect"
	"testing"
)

func TestValidateCoupons(t *testing.T) {
	tests := []struct {
		name         string
		codes        []string
		businessLine []string
		isActive     []bool
		want         []string
	}{
		{
			name:         "basic ordering by business line",
			codes:        []string{"e1", "g1", "p1", "r1"},
			businessLine: []string{"electronics", "grocery", "pharmacy", "restaurant"},
			isActive:     []bool{true, true, true, true},
			want:         []string{"e1", "g1", "p1", "r1"},
		},
		{
			name:         "filter inactive and invalid codes",
			codes:        []string{"okA", "bad*code", "okB"},
			businessLine: []string{"electronics", "grocery", "electronics"},
			isActive:     []bool{true, true, false},
			want:         []string{"okA"},
		},
		{
			name:         "sort within same category lexicographically",
			codes:        []string{"z", "a1", "a"},
			businessLine: []string{"grocery", "grocery", "grocery"},
			isActive:     []bool{true, true, true},
			want:         []string{"a", "a1", "z"},
		},
		{
			name:         "unsupported business line is filtered",
			codes:        []string{"x1", "y1"},
			businessLine: []string{"other", "electronics"},
			isActive:     []bool{true, true},
			want:         []string{"y1"},
		},
		{
			name:         "empty code should be filtered",
			codes:        []string{"", "valid"},
			businessLine: []string{"electronics", "electronics"},
			isActive:     []bool{true, true},
			want:         []string{"valid"},
		},
		{
			name:         "underscore allowed and ordering maintained",
			codes:        []string{"_b", "_a"},
			businessLine: []string{"pharmacy", "pharmacy"},
			isActive:     []bool{true, true},
			want:         []string{"_a", "_b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateCoupons(tt.codes, tt.businessLine, tt.isActive)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("validateCoupons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidCode(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "empty", s: "", want: false},
		{name: "letters", s: "AbCz", want: true},
		{name: "digits", s: "12345", want: true},
		{name: "underscore", s: "a_b", want: true},
		{name: "invalid char star", s: "a*b", want: false},
		{name: "invalid hyphen", s: "a-b", want: false},
		{name: "mix valid", s: "A1_b2", want: true},
		{name: "space invalid", s: "A B", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidCode(tt.s)
			if got != tt.want {
				t.Fatalf("isValidCode(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
