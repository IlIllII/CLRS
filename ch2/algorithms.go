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
// Selection Sort: Return the first n sorted elements of an array.
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
