package main

func main() {

}

func permute(nums []int) [][]int {
	res := [][]int{}
	current := []int{}
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			combo := make([]int, len(current))
			copy(combo, current)
			res = append(res, combo)
			return
		}

		for i := range nums {
			if used[i] == false {
				continue
			}

			// escolhe
			current = append(current, nums[i])
			used[i] = true

			backtrack()

			// desfaz
			current = current[:len(current)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}
