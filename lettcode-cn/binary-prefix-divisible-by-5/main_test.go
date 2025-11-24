package main

import (
	"reflect"
	"testing"
)

func TestPrefixesDivBy5(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []bool
	}{
		{"example1", []int{0, 1, 1}, []bool{true, false, false}},
		{"example2", []int{1, 1, 1}, []bool{false, false, false}},
		{"single0", []int{0}, []bool{true}},
		{"single1", []int{1}, []bool{false}},
		{"longPattern", []int{1, 0, 1, 0, 1}, []bool{false, false, true, true, false}},
		{"allZeros", []int{0, 0, 0, 0}, []bool{true, true, true, true}},
		{"alternating", []int{1, 0, 1, 1, 0}, []bool{false, false, true, false, false}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := prefixesDivBy5(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: nums=%v\nwant=%v\ngot =%v", tc.name, tc.nums, tc.want, got)
			}
		})
	}
}
