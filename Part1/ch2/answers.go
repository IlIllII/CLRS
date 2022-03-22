package ch2

//lint:file-ignore U1000 Ignore all unused code

func _InsertionSort(array []int) []int {
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

func _NonincreasingInsertionSort(array []int) []int {
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

func _LinearSearch(value int, array []int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return i
		}
	}

	return -1
}

func _BinaryAddition(A []int, B []int) []int {
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

func _SelectionSort(array []int) []int {
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

func _MergeSort(array []int) []int {
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
	return _merge(MergeSort(left), MergeSort(right))
}

func _merge(left []int, right []int) []int {
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

func _BinarySearch(array []int, val int) int {
	if len(array) == 0 {
		return -1
	}
	return _binaryRecur(array, val, 0, len(array)-1)
}

func _binaryRecur(array []int, val int, start int, end int) int {
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

	left := _binaryRecur(array, val, left_start, left_end)
	right := _binaryRecur(array, val, right_start, right_end)

	if left != -1 {
		return left
	} else if right != -1 {
		return right
	} else {
		return -1
	}
}

func _SumInSet(array []int, x int) bool {
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

func _CoarseMergeSort(array []int) []int {
	if len(array) < 4 {
		return InsertionSort(array)
	}

	left := array[:len(array)/2]
	right := array[len(array)/2:]
	return _merge(MergeSort(left), MergeSort(right))
}

func _BubbleSort(array []int) []int {
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

func _PolynomialEval(coeffs []int, x int) int {
	y := 0
	for i := len(coeffs) - 1; i >= 0; i-- {
		y = coeffs[i] + x*y
	}
	return y
}

func _NumInversions(array []int) int {
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
