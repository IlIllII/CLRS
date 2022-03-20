package ch2

// Brute Force Maximum subarray - 4.1 in-chapter algo
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
	if len(arr) == 0 {
		return arr
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
