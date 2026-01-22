package main

func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	current := []int{}
	var backtrack func(remaining, start int)
	backtrack = func(remaining, start int) {
		if remaining == 0 {
			combo := make([]int, len(current))
			copy(combo, current)
			res = append(res, combo)
			return
		}

		if remaining < 0 {
			return
		}

		// start para nao repetir numeros no array
		for i := start; i < len(current); i++ {
			current = append(current, candidates[i])
			backtrack(remaining-candidates[i], i)
			current = current[:len(current)-1]
		}
	}
	backtrack(target, 0)
	return res
}
