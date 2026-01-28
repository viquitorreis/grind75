package main

import "fmt"

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}

func canPartition(nums []int) bool {
	// nao pode ser odd logicamente
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	targetSum := sum / 2
	dp := make([]bool, targetSum+1)
	dp[0] = true

	// precisamos explorar possibilidades...
	// começamos do ultimo número
	for _, num := range nums {
		for j := targetSum; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[targetSum]
}
