package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	current := []int{}

	var backtrack func(remaining, start int)
	backtrack = func(remaining, start int) {
		if remaining == 0 {
			// copiamos para não perder a referencia
			curr := [][]int{}
			copy(curr, res)
			curr = append(curr, current)
			copy(res, curr)
			return
		}

		if remaining < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			current = append(current, candidates[i]) // escolhemos o current
			backtrack(remaining-current[i], i)
			current = current[:len(current)-1] // removemos o current
		}
	}

	backtrack(target, 0)

	return res
}
