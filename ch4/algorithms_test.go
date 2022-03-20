package ch2

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

// func TestMaximumSubarray(t *testing.T) {
// 	for _, test := range maxSubarrayTests {
// 		result := MaximumSubarray(test.arr)
// 		if !areSlicesEqual(result, test.result) {
// 			t.Errorf("Expected: %v; Got: %v", test.result, result)
// 		}
// 	}
// }

type maxSubarray struct {
	arr    []int
	result []int
}

var maxSubarrayTests []maxSubarray = []maxSubarray{
	{[]int{}, []int{}},
	{[]int{0, 1, 0}, []int{0, 1}},
	{[]int{1, 2, 0, 3}, []int{0, 3}},
	{[]int{5, 2, 1, 3, 5}, []int{1, 3, 5}},
	{[]int{1, 2, 5, 2, 4}, []int{1, 2, 5}},
	{[]int{1, 2, 5, 0, 10}, []int{0, 10}},
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
