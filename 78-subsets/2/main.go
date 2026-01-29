package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	var backtrack func(i int)
	current := []int{}

	backtrack = func(i int) {
		if i >= len(nums) {
			combo := make([]int, len(current))
			copy(combo, current)
			res = append(res, combo)
			return
		}

		current = append(current, nums[i])
		backtrack(i + 1)

		current = current[:len(current)-1]
		backtrack(i + 1)
	}

	backtrack(0)

	return res
}
