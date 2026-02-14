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

	curr := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			// escolhe
			curr = append(curr, nums[i])

			backtrack(i + 1)

			// desfaz -> SOBE PARA O PAI.
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0)

	return res
}
