package main

func main() {

}

func permute(nums []int) [][]int {
	perm := [][]int{}
	current := []int{}
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			combo := make([]int, len(nums))
			copy(combo, current)
			perm = append(perm, combo)
			return
		}

		for i := range nums {
			if used[i] {
				continue
			}

			used[i] = true
			current = append(current, nums[i])
			backtrack()

			current = current[:len(current)-1]
			used[i] = false
		}
	}

	backtrack()

	return perm
}
