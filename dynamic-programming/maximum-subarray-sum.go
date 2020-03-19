package main

import (
	"errors"
	"fmt"
)

func main() {
	var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(MaximumSubarraySum(nums))

	var nums2 = []int{-2, -3, 4, -1, -2, 1, 5, -3}
	fmt.Println(MaximumSubarraySum(nums2))

	var nums3 = []int{-1, -2, -3, -4, -5}
	fmt.Println(MaximumSubarraySum(nums3))

	// Using dynamic programming
	fmt.Println(MaxSubarraySumWithNegativeNumbers(nums))
	fmt.Println(MaxSubarraySumWithNegativeNumbers(nums2))
	fmt.Println(MaxSubarraySumWithNegativeNumbers(nums3))
}

// Kadane's algorithm
//https://www.geeksforgeeks.org/largest-sum-contiguous-subarray/
func MaximumSubarraySum(numbers []int) int {
	var globalMax = 0
	var localMax = 0
	for i := 0; i < len(numbers); i++ {
		localMax += numbers[i]
		if localMax < 0 {
			localMax = 0
		} else if localMax > globalMax {
			globalMax = localMax
		}
	}
	return globalMax
}

// Using dynamic programming
func MaxSubarraySumWithNegativeNumbers(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("cannot find maximum in an empty array")
	}
	var globalMax = numbers[0]
	var localMax = numbers[0]
	for i := 1; i < len(numbers); i++ {
		localMax = max(numbers[i], localMax + numbers[i])
		globalMax = max(localMax, globalMax)
	}
	return globalMax, nil
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
