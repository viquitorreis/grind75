package main

import "fmt"

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}

	current := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		combo := make([]int, len(current))
		copy(combo, current)
		res = append(res, combo)

		for i := start; i < len(nums); i++ {
			// SÓ PULAMOS duplicatas se não for a primeira iteração do loop nessa chamada recursiva especifica
			if i > start && nums[i] == nums[i-1] {
				return
			}

			// escolhe
			current = append(current, nums[i])

			backtrack(i + 1)

			// desfaz
			current = current[:len(current)-1]
		}
	}
	backtrack(0)

	return res
}
