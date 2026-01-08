package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-1, 2, 3}
	fmt.Println(maxSubArray(arr))
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// guardamos a soma em um array
	// essa soma[i] será a soma total até agora
	sums := make([]int, len(nums))
	sums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		sums[i] = max(sums[i-1]+nums[i], nums[i])
	}

	maxSum := math.MinInt
	for i := range len(nums) {
		maxSum = max(maxSum, sums[i])
	}

	return maxSum
}
