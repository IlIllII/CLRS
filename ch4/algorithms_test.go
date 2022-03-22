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

func TestBasicMatrixMultiply(t *testing.T) {
	for _, test := range matmulTests {
		result := BasicMatrixMultiply(test.A, test.B)
		if !areMatricesEqual(result, test.C) {
			t.Errorf("Expected: %v; Got: %v", test.C, result)
		}
	}
}

func TestDCMatrixMultiply(t *testing.T) {
	for _, test := range matmulTests {
		result := DivideAndConquerMatMul(test.A, test.B)
		if !areMatricesEqual(result, test.C) {
			t.Errorf("Expected: %v; Got: %v", test.C, result)
		}
	}
}

func TestStrassens(t *testing.T) {
	for _, test := range matmulTests {
		result := StrassenMatMul(test.A, test.B)
		if !areMatricesEqual(result, test.C) {
			t.Errorf("Expected: %v; Got: %v", test.C, result)
		}
	}
}

func TestGeneralStrassens(t *testing.T) {
	test := matmulTest{
		A: [][]int{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
		B: [][]int{
			{1, 2, 2},
			{1, 2, 2},
			{1, 2, 2},
		},
		C: [][]int{
			{1, 2, 2},
			{1, 2, 2},
			{1, 2, 2},
		},
	}
	result := GeneralStrassenMatMul(test.A, test.B)
	if !areMatricesEqual(result, test.C) {
		t.Errorf("Expected: %v; Got: %v", test.C, result)
	}
	test = matmulTest{
		A: [][]int{
			{1, 0},
			{0, 1},
		},
		B: [][]int{
			{1, 2},
			{1, 2},
		},
		C: [][]int{
			{1, 2},
			{1, 2},
		},
	}
	result = GeneralStrassenMatMul(test.A, test.B)
	if !areMatricesEqual(result, test.C) {
		t.Errorf("Expected: %v; Got: %v", test.C, result)
	}
}

// * Helpers

type matmulTest struct {
	A [][]int
	B [][]int
	C [][]int
}

var matmulTests []matmulTest = []matmulTest{
	{
		A: [][]int{
			{1, 0},
			{0, 1},
		},
		B: [][]int{
			{1, 2},
			{1, 2},
		},
		C: [][]int{
			{1, 2},
			{1, 2},
		},
	},
	{
		A: [][]int{
			{0, 0},
			{0, 0},
		},
		B: [][]int{
			{1, 2},
			{1, 2},
		},
		C: [][]int{
			{0, 0},
			{0, 0},
		},
	},
	{
		A: [][]int{
			{1, 2, 3, 1},
			{4, 5, 6, 1},
			{7, 8, 9, 1},
			{1, 1, 1, 1},
		},
		B: [][]int{
			{10, 11, 12, 1},
			{13, 14, 15, 1},
			{16, 17, 18, 1},
			{16, 17, 18, 1},
		},
		C: [][]int{
			{100, 107, 114, 7},
			{217, 233, 249, 16},
			{334, 359, 384, 25},
			{55, 59, 63, 4},
		},
	},
	{
		A: [][]int{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
		B: [][]int{
			{5, 5, 5, 5},
			{0, 4, 4, 4},
			{0, 0, 3, 3},
			{2, 0, 0, 2},
		},
		C: [][]int{
			{5, 5, 5, 5},
			{0, 4, 4, 4},
			{0, 0, 3, 3},
			{2, 0, 0, 2},
		},
	},
}

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

func areMatricesEqual(A [][]int, B [][]int) bool {
	if len(A) != len(B) || len(A[0]) != len(B[0]) {
		return false
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] != B[i][j] {
				return false
			}
		}
	}

	return true
}
