package ch4

// Brute Force Maximum subarray - 4.1 in-chapter algo & Ex 4.1-2
//
// This is the first solution proposed for the volatile chemical
// company stock price problem before transforming the input to allow
// for a proper maximum subarray algorithm.
//
// The goal here is to simply find every combination of buy and sell prices
// and return the slice of the original array that leads to the highest
// return possible. This approach yields an Omega(n^2) algorithm.
//
// If we assume that the input array is a list of stock prices,
// Ex: BruteMaximumSubarray([1, 2, 0, 5, 10]) -> [0, 5, 10]
//
// So we buy at 0, hold until 10, and then sell.
func BruteMaximumSubarray(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	valid := false
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-prev >= 0 {
			valid = true
		}
		prev = arr[i]
	}
	if !valid {
		return []int{}
	}

	globalMaxIdx := 0
	globalMinIdx := 0
	localMaxIdx := 0
	localMinIdx := 0
	for i := 0; i < len(arr); i++ {
		localMinIdx = i
		for j := i; j < len(arr); j++ {
			localMaxIdx = j
			diff := arr[localMaxIdx] - arr[localMinIdx]
			if diff > arr[globalMaxIdx]-arr[globalMinIdx] {
				globalMinIdx = localMinIdx
				globalMaxIdx = localMaxIdx
			}
		}
	}
	return arr[globalMinIdx : globalMaxIdx+1]
}

// Divide-and-conquer Maximum Subarray - 4.1 in-chapter algorithm
//
// This algorithm approaches the previous problem, the volatile company
// stock price problem, using divide-and-conquer to allow us to get a better
// running time: o(n^2). Note the difference between the two: the first had a
// lower bound of n^2, this one should have a loose upper bound of n^2 and
// should therefore be strictly faster.
//
// Note: You can assume that there is one and only one maximum subarray in
// each test case.
//
// Interface is identical to brute force algorithm.
func MaximumSubarray(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	valid := false
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-prev >= 0 {
			valid = true
		}
		prev = arr[i]
	}
	if !valid {
		return []int{}
	}

	daily_changes := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		daily_changes = append(daily_changes, arr[i]-arr[i-1])
	}

	result := findMaxSubarray(daily_changes, 0, len(daily_changes)-1)
	low := result[0]
	high := result[1]
	return arr[low : high+2]
}

func findMaxSubarray(arr []int, low int, high int) [3]int {
	if low == high {
		return [3]int{low, high, arr[low]}
	} else {
		mid := int((low + high) / 2)

		leftResult := findMaxSubarray(arr, low, mid)
		leftLow := leftResult[0]
		leftHigh := leftResult[1]
		leftSum := leftResult[2]

		rightResult := findMaxSubarray(arr, mid+1, high)
		rightLow := rightResult[0]
		rightHigh := rightResult[1]
		rightSum := rightResult[2]

		crossResult := findMaxCrossingSubarray(arr, low, mid, high)
		crossLow := crossResult[0]
		crossHigh := crossResult[1]
		crossSum := crossResult[2]

		if leftSum >= rightSum && leftSum >= crossSum {
			return [3]int{leftLow, leftHigh, leftSum}
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return [3]int{rightLow, rightHigh, rightSum}
		} else {
			return [3]int{crossLow, crossHigh, crossSum}
		}
	}
}

func findMaxCrossingSubarray(arr []int, low int, mid int, high int) [3]int {
	leftSum := -100000
	maxLeft := 0
	maxRight := 0
	sum := 0
	for i := mid; i >= low; i-- {
		sum += arr[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	rightSum := -100000
	sum = 0

	for i := mid + 1; i <= high; i++ {
		sum += arr[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}
	result := [3]int{maxLeft, maxRight, leftSum + rightSum}
	return result
}

// 4.1-3 Coarsening the max subarray algorithm
//
// Use a brute force implementation to coarsen the leaves of the
// divide-and-conquer approach and speed up the algorithm.
//
// Interface is identical to previous divide and conquer algorithm.
func CoarsenedMaximumSubarray(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	valid := false
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-prev >= 0 {
			valid = true
		}
		prev = arr[i]
	}
	if !valid {
		return []int{}
	}

	daily_changes := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		daily_changes = append(daily_changes, arr[i]-arr[i-1])
	}

	result := coarseMaxSubarray(daily_changes, 0, len(daily_changes)-1)
	low := result[0]
	high := result[1]
	return arr[low : high+2]
}

func coarseMaxSubarray(arr []int, low int, high int) [3]int {
	if low == high {
		return [3]int{low, high, arr[low]}
	} else if high-low < 2 {
		array := BruteMaximumSubarray(arr[low:high])
		if len(array) == 2 {
			return [3]int{low, high, arr[low] + arr[high]}
		} else if array[0] == arr[low] {
			return [3]int{low, low, arr[low]}
		} else {
			return [3]int{high, high, arr[high]}
		}
	} else {
		mid := int((low + high) / 2)

		leftResult := findMaxSubarray(arr, low, mid)
		leftLow := leftResult[0]
		leftHigh := leftResult[1]
		leftSum := leftResult[2]

		rightResult := findMaxSubarray(arr, mid+1, high)
		rightLow := rightResult[0]
		rightHigh := rightResult[1]
		rightSum := rightResult[2]

		crossResult := findMaxCrossingSubarray(arr, low, mid, high)
		crossLow := crossResult[0]
		crossHigh := crossResult[1]
		crossSum := crossResult[2]

		if leftSum >= rightSum && leftSum >= crossSum {
			return [3]int{leftLow, leftHigh, leftSum}
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return [3]int{rightLow, rightHigh, rightSum}
		} else {
			return [3]int{crossLow, crossHigh, crossSum}
		}
	}
}

// Ex 4.1-5 Linear time maximum subarrays
//
// Interface is identical to previous maximum subarray algorithms.
func LinearTimeMaximumSubarray(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	valid := false
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-prev >= 0 {
			valid = true
		}
		prev = arr[i]
	}
	if !valid {
		return []int{}
	}

	daily_changes := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		daily_changes = append(daily_changes, arr[i]-arr[i-1])
	}

	maxSum := 0
	maxStart := 0
	maxEnd := 0
	currSum := 0
	currStart := 0

	for i := 0; i < len(daily_changes); i++ {
		currSum += daily_changes[i]
		if currSum > maxSum {
			maxSum = currSum
			maxEnd = i
			maxStart = currStart
		}

		if currSum < 0 {
			currSum = 0
			currStart = i + 1
		}
	}

	return arr[maxStart : maxEnd+2]
}

// Theta(n^3) approach to matrix multiplication 4.2 in chapter algorithm
//
// This is the first approach to matrix multiplation and uses a simple
// triply nested loop.
//
// You can assume that all inputs are square matrices.
//
// Ex: [[1, 0], [0, 1]]
func BasicMatrixMultiply(A, B [][]int) [][]int {
	n := len(A)
	C := [][]int{}
	for i := 0; i < n; i++ {
		CRow := []int{}
		for j := 0; j < n; j++ {
			c := 0
			for k := 0; k < n; k++ {
				c += A[i][k] * B[k][j]
			}
			CRow = append(CRow, c)
		}
		C = append(C, CRow)
	}

	return C
}

// 4.2 in-chapter algorithm : Recursive divide and conquer approach to
// matrix multiplication. Theta(n^3).
//
// You can assume that input matrices are square matrices with
// dimensions that are powers of two.
func DivideAndConquerMatMul(A, B [][]int) [][]int {
	n := len(A)
	C := makeMat(2 * n)

	// Base case
	if n == 1 {
		C[0][0] = A[0][0] * B[0][0]
		return C
	}

	// Partition Matrices
	A11 := makeMat(n)
	A12 := makeMat(n)
	A21 := makeMat(n)
	A22 := makeMat(n)
	B11 := makeMat(n)
	B12 := makeMat(n)
	B21 := makeMat(n)
	B22 := makeMat(n)
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			A11[i][j] = A[i][j]
			A12[i][j] = A[i][j+n/2]
			A21[i][j] = A[i+n/2][j]
			A22[i][j] = A[i+n/2][j+n/2]
			B11[i][j] = B[i][j]
			B12[i][j] = B[i][j+n/2]
			B21[i][j] = B[i+n/2][j]
			B22[i][j] = B[i+n/2][j+n/2]
		}
	}

	// Compute submatrices
	C11 := matAdd(DivideAndConquerMatMul(A11, B11), DivideAndConquerMatMul(A12, B21))
	C12 := matAdd(DivideAndConquerMatMul(A11, B12), DivideAndConquerMatMul(A12, B22))
	C21 := matAdd(DivideAndConquerMatMul(A21, B11), DivideAndConquerMatMul(A22, B21))
	C22 := matAdd(DivideAndConquerMatMul(A21, B12), DivideAndConquerMatMul(A22, B22))

	combineSubMatrices(C, C11, C12, C21, C22, n)
	return C
}

func matAdd(A, B [][]int) [][]int {
	n := len(A)
	C := makeMat(2 * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

func matSub(A, B [][]int) [][]int {
	n := len(A)
	C := makeMat(2 * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] - B[i][j]
		}
	}
	return C
}

func makeMat(n int) [][]int {
	C := make([][]int, n/2)
	for i := range C {
		C[i] = make([]int, n/2)
	}
	return C
}

func combineSubMatrices(C, C11, C12, C21, C22 [][]int, n int) {
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			C[i][j] += C11[i][j]
			C[i][j+n/2] += C12[i][j]
			C[i+n/2][j] += C21[i][j]
			C[i+n/2][j+n/2] += C22[i][j]
		}
	}
}

// 4.2 in-chapter algo & ex 4.2-2
//
// Strassen's method for matrix multiplication.
//
// As above, assume square matrices with dimensions that are
// power's of two.
func StrassenMatMul(A, B [][]int) [][]int {
	n := len(A)
	C := makeMat(2 * n)

	// Base case
	if n == 1 {
		C[0][0] = A[0][0] * B[0][0]
		return C
	}

	// Partition Matrices
	A11 := makeMat(n)
	A12 := makeMat(n)
	A21 := makeMat(n)
	A22 := makeMat(n)
	B11 := makeMat(n)
	B12 := makeMat(n)
	B21 := makeMat(n)
	B22 := makeMat(n)
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			A11[i][j] = A[i][j]
			A12[i][j] = A[i][j+n/2]
			A21[i][j] = A[i+n/2][j]
			A22[i][j] = A[i+n/2][j+n/2]
			B11[i][j] = B[i][j]
			B12[i][j] = B[i][j+n/2]
			B21[i][j] = B[i+n/2][j]
			B22[i][j] = B[i+n/2][j+n/2]
		}
	}

	S1 := matSub(B12, B22)
	S2 := matAdd(A11, A12)
	S3 := matAdd(A21, A22)
	S4 := matSub(B21, B11)
	S5 := matAdd(A11, A22)
	S6 := matAdd(B11, B22)
	S7 := matSub(A12, A22)
	S8 := matAdd(B21, B22)
	S9 := matSub(A11, A21)
	S10 := matAdd(B11, B12)

	P1 := StrassenMatMul(A11, S1)
	P2 := StrassenMatMul(S2, B22)
	P3 := StrassenMatMul(S3, B11)
	P4 := StrassenMatMul(A22, S4)
	P5 := StrassenMatMul(S5, S6)
	P6 := StrassenMatMul(S7, S8)
	P7 := StrassenMatMul(S9, S10)

	// Compute submatrices
	C11 := matSub(matAdd(P5, matAdd(P4, P6)), P2)
	C12 := matAdd(P1, P2)
	C21 := matAdd(P3, P4)
	C22 := matSub(matSub(matAdd(P5, P1), P3), P7)

	combineSubMatrices(C, C11, C12, C21, C22, n)
	return C
}

// Ex 4.2-3 Strassen's with any n
//
// You should not assume that input matrices have dimensions that are
// powers of 2, but you can assume they are still square.
func GeneralStrassenMatMul(A, B [][]int) [][]int {
	n := len(A)
	wasAdjusted := false
	if !isPowerOfTwo(n) {
		wasAdjusted = true
		m := nextPowerOfTwo(len(A))
		for i := 0; i < n; i++ {
			for j := n; j < m; j++ {
				A[i] = append(A[i], 0)
				B[i] = append(B[i], 0)
			}
		}
		for i := n; i < m; i++ {
			newRow1 := []int{}
			newRow2 := []int{}
			for j := 0; j < m; j++ {
				newRow1 = append(newRow1, 0)
				newRow2 = append(newRow2, 0)
			}
			A = append(A, newRow1)
			B = append(B, newRow2)
		}
	}

	C := StrassenMatMul(A, B)

	if wasAdjusted {
		for i := 0; i < n; i++ {
			C[i] = C[i][:n]
		}
		C = C[:n]
	}

	return C
}

func isPowerOfTwo(x int) bool {
	return x&(x-1) == 0 && x != 0
}

func nextPowerOfTwo(x int) int {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return (x - (x >> 1)) * 2
}
