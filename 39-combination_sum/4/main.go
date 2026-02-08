package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	current := []int{}
	var backtrack func(idx, remaining int)
	backtrack = func(start, remaining int) {
		if remaining == 0 {
			combo := make([]int, len(current))
			copy(combo, current)
			res = append(res, combo)
			return
		}

		if remaining < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			// escolhe
			num := candidates[i]
			current = append(current, num)
			backtrack(i, remaining-num)
			current = current[:len(current)-1]
		}
	}

	backtrack(0, target)
	return res
}
