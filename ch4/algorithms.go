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
