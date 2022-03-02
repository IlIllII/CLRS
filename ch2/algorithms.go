package ch2

// Ch 2.1 in-chapter algo
// InsertionSort: return an array sorted in non-decreasing order.
//
// Ex: InsertionSort([3, 2, 1]) -> [1, 2, 3]
//
// Complexity?
//
// Pros, cons?
func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		elementToInsert := array[i]
		insertionIndex := i

		for j := i - 1; j >= 0; j-- {
			if elementToInsert > array[j] {
				break
			}

			array[j+1] = array[j]
			insertionIndex--
		}
		array[insertionIndex] = elementToInsert
	}
	return array
}

// Exercise 2.1-2
//
// Ex: NonincreasingInsertionSort([1, 2, 3]) -> [3, 2, 1]
func NonincreasingInsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		elementToInsert := array[i]
		insertionIndex := i

		for j := i - 1; j >= 0; j-- {
			if elementToInsert < array[j] {
				break
			}

			array[j+1] = array[j]
			insertionIndex--
		}
		array[insertionIndex] = elementToInsert
	}
	return array
}

// Exercise 2.1-3
//
// Linear Search: search an array for a value and return the index
// at which its first found. If the value isn't found in the array, return -1
// for the index.
//
// Ex: LinearSearch(1, [1, 2, 3]) -> 0
//
// Complexity?
//
// Pros, cons?
func LinearSearch(value int, array []int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return i
		}
	}

	return -1
}

// Exercise 2.1-4
//
// Binary Addition: return an n+1-element array containing the
// addition of two n-bit binary integers stored in n-element arrays.
//
// You can assume that only 0s and 1s will be passed as elements in the input arrays.
//
// Ex: BinaryAddition([0, 1], [1, 0]) -> [0, 1, 1]
func BinaryAddition(A []int, B []int) []int {
	C := make([]int, len(A)+1)
	carry := 0

	for i := len(A) - 1; i >= 0; i-- {
		digit := A[i] + B[i] + carry
		carry = int(digit / 2)
		digit = digit % 2
		C[i+1] = digit
	}

	C[0] = carry
	return C
}

// Exercise 2.2-2
//
// Selection Sort: Return a sorted array.
//
// Ex: SelectionSort([3, 2, 1], 2) -> [1, 2, 3]
//
// Complexity?
//
// Pros, cons?
func SelectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		element := array[i]
		min := element
		minIndex := i
		for j := i; j < len(array); j++ {
			if array[j] < min {
				min = array[j]
				minIndex = j
			}
		}
		array[minIndex] = element
		array[i] = min
	}
	return array
}

// Ch 2.3 in-chapter algo/Exercise 2.3-2
// MergeSort: return an array sorted in non-decreasing order.
//
// Ex: MergeSort([3, 2, 1]) -> [1, 2, 3]
//
// Complexity?
//
// Pros, cons?
//
// Hint: you can write a helper function.
func MergeSort(array []int) []int {
	if len(array) == 1 || len(array) == 0 {
		return array
	} else if len(array) == 2 {
		if array[1] < array[0] {
			tmp := array[0]
			array[0] = array[1]
			array[1] = tmp
		}
		return array
	}

	left := array[:len(array)/2]
	right := array[len(array)/2:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left []int, right []int) []int {
	result := make([]int, 1)[:0]
	i, j := 0, 0
	for {
		if i >= len(left) || j >= len(right) {
			break
		}
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	if i < len(left) {
		result = append(result, left[i:]...)
	} else if j < len(right) {
		result = append(result, right[j:]...)
	}
	return result
}

// Exercise 2.3-5
// BinarySearch: search an array for a value and return the index
// at which its first found. If the value isn't found in the array, return -1
// for the index.
//
// You can assume that the input array is already sorted in non-decreasing
// order.
//
// Ex: BinarySearch([1, 2, 3], 1) -> 0
//
// Complexity?
//
// Pros, cons?
func BinarySearch(array []int, val int) int {
	if len(array) == 0 {
		return -1
	}
	return binaryRecur(array, val, 0, len(array)-1)
}

func binaryRecur(array []int, val int, start int, end int) int {
	if start == end {
		if array[start] == val {
			return start
		} else {
			return -1
		}
	}
	left_start := start
	left_end := start + (end-start)/2
	right_start := start + (end-start)/2 + 1
	right_end := end

	left := binaryRecur(array, val, left_start, left_end)
	right := binaryRecur(array, val, right_start, right_end)

	if left != -1 {
		return left
	} else if right != -1 {
		return right
	} else {
		return -1
	}
}

// Exercise 2.3-7
// Given an array S of n integers and an integer x, is there a combination of
// two integers in S that sum to exactly x?
//
// Ex: SumInSet([1, 2, 3], 3) -> true
//
// Complexity?
//
// Pros, cons?
//
// Bonus: CLRS wants an O(nlgn) solution, can you find an O(n) solution?
func SumInSet(array []int, x int) bool {
	complements := map[int]int{}
	for _, num := range array {
		_, ok := complements[num]
		if ok {
			return true
		}
		complement := x - num
		complements[complement] = num
	}
	return false
}

// Problem 2-1
// CoarsenedMergeSort: return an array sorted in non-decreasing order.
//
// Ex: CoarsenedMergeSort([3, 2, 1]) -> [1, 2, 3]
func CoarseMergeSort(array []int) []int {
	if len(array) < 4 {
		return InsertionSort(array)
	}

	left := array[:len(array)/2]
	right := array[len(array)/2:]
	return merge(MergeSort(left), MergeSort(right))
}

// Problem 2-2
// BubbleSort: return an array sorted in non-decreasing order.
//
// Complexity?
//
// Pros? Cons?
func BubbleSort(array []int) []int {
	for i := range array {
		for j := len(array) - 1; j > i; j-- {
			if array[j] < array[j-1] {
				tmp := array[j]
				array[j] = array[j-1]
				array[j-1] = tmp
			}
		}
	}
	return array
}

// Problem 2-3
// PolynomialEvaluation: given an array of coefficients, evaluate
// the polynomial at x.
//
// Ex: PolynomialEval([1,1,1], 2) -> 7
func PolynomialEval(coeffs []int, x int) int {
	y := 0
	for i := len(coeffs) - 1; i >= 0; i-- {
		y = coeffs[i] + x*y
	}
	return y
}

// Problem 2-4
// NumInversions: given an array of integers, how many inversions does
// contain?
//
// Ex: NumInversions([2, 3, 8, 6, 1]) -> 5
func NumInversions(array []int) int {
	inversions := 0
	for i := range array {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				inversions++
			}
		}
	}
	return inversions
}
