package main

import (
	"fmt"
	"math"
)

func main() {
	// nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	nums := []int{-2, 1, -3}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sums := make([]int, len(nums))
	sums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		sums[i] = max(nums[i]+sums[i-1], nums[i])
	}

	maxSum := math.MaxInt16

	for i := range sums {
		maxSum = max(maxSum, sums[i])
	}

	return maxSum
}
