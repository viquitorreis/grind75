package main

import "fmt"

func main() {
	cands := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(cands, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	current := []int{}

	var backtrack func(sum, start int)
	backtrack = func(sum, start int) {
		if sum == target {
			combo := make([]int, len(current))
			copy(combo, current)
			res = append(res, combo)
			return
		}

		if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			current = append(current, candidates[i]) // escolhe o elemento
			backtrack(sum+candidates[i], i)          // desfaz
			current = current[:len(current)-1]
		}
	}

	backtrack(0, 0)

	return res
}
