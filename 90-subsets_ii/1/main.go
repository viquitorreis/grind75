package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	// precisa fazer sort para as duplicatas ficarem juntas
	sort.Ints(nums)
	res := [][]int{}
	current := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		combo := make([]int, len(current))
		copy(combo, current)
		res = append(res, combo)

		for i := start; i < len(nums); i++ {
			// pula duplicatas e se for igual ao elemento anteior
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			current = append(current, nums[i])
			backtrack(i + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)

	return res
}
