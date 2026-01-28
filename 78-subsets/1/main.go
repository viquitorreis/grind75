package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}

	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(nums) {
			combo := make([]int, len(subset))
			copy(combo, subset)
			res = append(res, combo)
			return
		}

		// decisão para incluir o número em i
		subset = append(subset, nums[i])
		dfs(i + 1)

		// decisão para não incluir o numero na i
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}

	dfs(0)

	return res
}
