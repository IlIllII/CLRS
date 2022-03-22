package ch2

import (
	"testing"
)

// * Tests

func TestInsertionSort(t *testing.T) {
	for _, test := range integerSortTests {
		cp := make([]int, len(test.unsorted))
		copy(cp, test.unsorted)
		result := InsertionSort(cp)
		if !areSlicesEqual(result, test.sorted) {
			t.Errorf("Expected: %v; Got: %v", test.sorted, result)
		}
	}
}

func TestNonincreasingInsertionSort(t *testing.T) {
	for _, test := range integerSortTests {
		cp := make([]int, len(test.unsorted))
		copy(cp, test.unsorted)
		cpSorted := make([]int, len(test.unsorted))
		copy(cpSorted, test.sorted)
		cpSorted = reverseSlice(cpSorted[:])
		result := NonincreasingInsertionSort(cp)
		if !areSlicesEqual(result, cpSorted) {
			t.Errorf("Expected: %v; Got: %v", cpSorted, result)
		}
	}
}

func TestLinearSearch(t *testing.T) {
	for _, test := range integerSearchTests {
		index := LinearSearch(test.value, test.array)
		if index != test.index {
			t.Errorf("Searching %v for %v - Expected: %v; Got: %v", test.array, test.value, test.index, index)
		}
	}
}

func TestBinaryAddition(t *testing.T) {
	for _, test := range binaryAddTests {
		C := BinaryAddition(test.A, test.B)
		if !areSlicesEqual(C, test.C) {
			t.Errorf("Expected: %v + %v = %v; Got: %v", test.A, test.B, test.C, C)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	for _, test := range integerSortTests {
		cp := make([]int, len(test.unsorted))
		copy(cp, test.unsorted)
		res := SelectionSort(cp)
		if !areSlicesEqual(res[:], test.sorted) {
			t.Errorf("Expected: %v; Got: %v", test.sorted, res)
		}
	}
	result := SelectionSort([]int{4, 3, 2, 5, 1})
	if !areSlicesEqual(result, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Expected: %v; Got: %v", []int{}, result)
	}
}

func TestMergeSort(t *testing.T) {
	for _, test := range integerSortTests {
		cp := make([]int, len(test.unsorted))
		copy(cp, test.unsorted)
		res := MergeSort(cp)
		if !areSlicesEqual(res[:], test.sorted) {
			t.Errorf("Expected: %v; Got: %v", test.sorted, res)
		}
	}
	result := MergeSort([]int{4, 3, 2, 5, 1})
	if !areSlicesEqual(result, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Expected: %v; Got: %v", []int{}, result)
	}
}

func TestBinarySearch(t *testing.T) {
	for _, test := range sortedIntegerSearchTests {
		index := BinarySearch(test.array, test.value)
		if index != test.index {
			t.Errorf("Searching %v for %v - Expected: %v; Got: %v", test.array, test.value, test.index, index)
		}
	}
}

func TestSumInSet(t *testing.T) {
	for _, test := range sumInSetTests {
		result := SumInSet(test.array, test.x)
		if result != test.result {
			t.Errorf("Array: %v, x: %v - Expected: %v; Got: %v", test.array, test.x, test.result, result)
		}
	}
}

func TestCoarseMergeSort(t *testing.T) {
	for _, test := range integerSortTests {
		cp := make([]int, len(test.unsorted))
		copy(cp, test.unsorted)
		res := CoarseMergeSort(cp)
		if !areSlicesEqual(res[:], test.sorted) {
			t.Errorf("Expected: %v; Got: %v", test.sorted, res)
		}
	}

	result := CoarseMergeSort([]int{4, 3, 2, 5, 1})
	if !areSlicesEqual(result, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Expected: %v; Got: %v", []int{}, result)
	}
}

func TestBubbleSort(t *testing.T) {
	for _, test := range integerSortTests {
		cp := make([]int, len(test.unsorted))
		copy(cp, test.unsorted)
		res := BubbleSort(cp)
		if !areSlicesEqual(res[:], test.sorted) {
			t.Errorf("Expected: %v; Got: %v", test.sorted, res)
		}
	}
}

func TestPolynomialEval(t *testing.T) {
	for _, test := range polynomialTests {
		cp := make([]int, len(test.coeffs))
		copy(cp, test.coeffs)
		res := PolynomialEval(cp, test.x)
		if res != test.result {
			t.Errorf("Expected: %v; Got: %v", test.result, res)
		}
	}
}

func TestNumInversions(t *testing.T) {
	for _, test := range numInversionTests {
		cp := make([]int, len(test.array))
		copy(cp, test.array)
		res := NumInversions(cp)
		if res != test.result {
			t.Errorf("Array: %v; Expected: %v; Got: %v", test.array, test.result, res)
		}
	}
}

// * Helper definitions

type inversionTest struct {
	array  []int
	result int
}

var numInversionTests []inversionTest = []inversionTest{
	{[]int{2, 3, 8, 6, 1}, 5},
	{[]int{5, 4, 3, 2, 1}, 10},
	{[]int{1, 2, 3, 4, 5}, 0},
	{[]int{1, 1, 1, 1, 1}, 0},
}

type polynomialTest struct {
	coeffs []int
	x      int
	result int
}

var polynomialTests []polynomialTest = []polynomialTest{
	{[]int{0, 1, 2}, 2, 10},
	{[]int{1, 1, 1}, 2, 7},
	{[]int{1, 1, 1, 1, 1}, 2, 31},
}

type sumInSet struct {
	array  []int
	x      int
	result bool
}

var sumInSetTests []sumInSet = []sumInSet{
	{[]int{0, 1, 2, 3, 4}, 1, true},
	{[]int{0, 1, 2, 3, 4}, 3, true},
	{[]int{0, 1, 2, 3, 4}, 5, true},
	{[]int{0, 1, 2, 3, 4}, 20, false},
	{[]int{0, 1, 2, 3, 4}, 9, false},
	{[]int{0, 1, 2, 3, 4}, 7, true},
	{[]int{0, 1, 2, 3, 4}, 0, false},
}

func TestReverseSlice(t *testing.T) {
	odd_forward := []int{1, 2, 3}
	odd_reversed := []int{3, 2, 1}
	even_forward := []int{1, 2, 3, 4}
	even_reversed := []int{4, 3, 2, 1}

	if !areSlicesEqual(odd_reversed, reverseSlice(odd_forward)) || !areSlicesEqual(even_reversed, reverseSlice(even_forward)) {
		t.Errorf("Reversing the slice didn't work...")
	}
}

type binaryAdd struct {
	A []int
	B []int
	C []int
}

var binaryAddTests []binaryAdd = []binaryAdd{
	{[]int{0, 0}, []int{0, 0}, []int{0, 0, 0}},
	{[]int{0, 1}, []int{0, 1}, []int{0, 1, 0}},
	{[]int{1, 1}, []int{1, 1}, []int{1, 1, 0}},
	{[]int{1, 1, 1}, []int{1, 1, 1}, []int{1, 1, 1, 0}},
}

type integerSearch struct {
	array []int
	value int
	index int
}

var integerSearchTests []integerSearch = []integerSearch{
	{[]int{}, 2, -1},
	{[]int{0, 1, 2}, -1, -1},
	{[]int{0, 1, 2}, 1, 1},
	{[]int{0, 1, 2}, 2, 2},
	{[]int{2, 1, 0}, 2, 0},
	{[]int{0, 1, 1, 2}, 1, 1},
}

var sortedIntegerSearchTests []integerSearch = []integerSearch{
	{[]int{}, 2, -1},
	{[]int{0, 1, 2, 2}, -1, -1},
	{[]int{0, 1, 2, 2}, 1, 1},
	{[]int{0, 1, 2, 2}, 2, 2},
	{[]int{0, 0, 0, 1, 2, 3}, 2, 4},
	{[]int{0, 1, 1, 2}, 1, 1},
}

type integerSort struct {
	unsorted []int
	sorted   []int
}

var integerSortTests []integerSort = []integerSort{
	{[]int{}, []int{}},
	{[]int{0}, []int{0}},
	{[]int{1, 0}, []int{0, 1}},
	{[]int{0, 0, 0, 0}, []int{0, 0, 0, 0}},
	{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
	{[]int{1, 2, 3, 1, 2, 3}, []int{1, 1, 2, 2, 3, 3}},
	{[]int{-1, 2, 3, -1, 2, 3}, []int{-1, -1, 2, 2, 3, 3}},
	{[]int{-1, 2, 3, -1, 2, 3, -1, 2, 3, -1, 2, 3}, []int{-1, -1, -1, -1, 2, 2, 2, 2, 3, 3, 3, 3}},
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

func reverseSlice(a []int) []int {
	for i := 0; i < len(a)/2; i++ {
		tmp := a[i]
		a[i] = a[len(a)-1-i]
		a[len(a)-1-i] = tmp
	}
	return a
}
