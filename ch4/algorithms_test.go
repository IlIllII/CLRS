package ch4

import (
	"testing"
)

// * Tests

func TestBruteMaximumSubarray(t *testing.T) {
	for _, test := range maxSubarrayTests {
		result := BruteMaximumSubarray(test.arr)
		if !areSlicesEqual(result, test.result) {
			t.Errorf("Expected: %v; Got: %v", test.result, result)
		}
	}
}

func TestMaximumSubarray(t *testing.T) {
	for _, test := range maxSubarrayTests {
		result := MaximumSubarray(test.arr)
		if !areSlicesEqual(result, test.result) {
			t.Errorf("Expected: %v; Got: %v", test.result, result)
		}
	}
}

func TestCoarsenedMaximumSubarray(t *testing.T) {
	for _, test := range maxSubarrayTests {
		result := CoarsenedMaximumSubarray(test.arr)
		if !areSlicesEqual(result, test.result) {
			t.Errorf("Expected: %v; Got: %v", test.result, result)
		}
	}
}

func TestLinearTimeMaximumSubarray(t *testing.T) {
	for _, test := range maxSubarrayTests {
		result := LinearTimeMaximumSubarray(test.arr)
		if !areSlicesEqual(result, test.result) {
			t.Errorf("Expected: %v; Got: %v", test.result, result)
		}
	}
}

func TestEmptySubarrayMaximumSubarrays(t *testing.T) {
	result := BruteMaximumSubarray([]int{3, 2, 1})
	if !areSlicesEqual(result, []int{}) {
		t.Errorf("BruteMaximumSubarray - Expected: %v; Got: %v", []int{}, result)
	}
	result = MaximumSubarray([]int{3, 2, 1})
	if !areSlicesEqual(result, []int{}) {
		t.Errorf("MaximumSubarray - Expected: %v; Got: %v", []int{}, result)
	}
	result = CoarsenedMaximumSubarray([]int{3, 2, 1})
	if !areSlicesEqual(result, []int{}) {
		t.Errorf("CoarsenedMaximumSubarray - Expected: %v; Got: %v", []int{}, result)
	}
	result = LinearTimeMaximumSubarray([]int{3, 2, 1})
	if !areSlicesEqual(result, []int{}) {
		t.Errorf("LinearTimeMaximumSubarray - Expected: %v; Got: %v", []int{}, result)
	}
}

// * Helpers

type maxSubarray struct {
	arr    []int
	result []int
}

var maxSubarrayTests []maxSubarray = []maxSubarray{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{0, 1, 0}, []int{0, 1}},
	{[]int{1, 2, 0, 3}, []int{0, 3}},
	{[]int{5, 2, 1, 3, 5}, []int{1, 3, 5}},
	{[]int{1, 2, 5, 2, 4}, []int{1, 2, 5}},
	{[]int{1, 2, 5, 0, 10}, []int{0, 10}},
	{[]int{1, 2, 5, 7, 10}, []int{1, 2, 5, 7, 10}},
}

func areSlicesEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, val := range a {
		if val != b[i] {
			return false
		}
	}

	return true
}
