package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	curr := []int{}
	res := [][]int{}
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(curr) == len(nums) {
			combo := make([]int, len(curr))
			copy(combo, curr)
			res = append(res, combo)
			return
		}

		for i := range nums {
			if used[i] {
				continue // usou nessa iteração
			}

			// escolhe
			curr = append(curr, nums[i])
			used[i] = true

			// explora
			backtrack()

			// desfaz
			curr = curr[:len(curr)-1]
			used[i] = false
		}
	}

	backtrack()

	return res
}
