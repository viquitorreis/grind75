package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-2, -1}
	println(maxSubArray(nums))
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

	maxSum := math.MinInt

	for i := range len(sums) {
		maxSum = max(sums[i], maxSum)
	}

	return maxSum
}

func printArr(nums []int) {
	for i := range nums {
		fmt.Printf("%d ", nums[i])
	}
	println()
}
