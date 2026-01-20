package main

func main() {

}

func permute(nums []int) [][]int {
	res := [][]int{}
	current := []int{}
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		// base case - na permutação, para cada elemento, temos n combinações, sendo n = len(nums)
		if len(current) == len(nums) {
			combo := make([]int, len(current))
			copy(combo, current)
			res = append(res, combo)
			return
		}

		// loop em todos os números
		for i := range nums {
			// base case, ja usamos esse número...
			if used[i] {
				continue
			}

			// caso nao usou o número, adicionamos a current (escolha), e chamamos recorrente
			used[i] = true
			current = append(current, nums[i])
			backtrack()

			// ao fim da recursão, se não chegou no base case, removemos o número dos usados e current, pois tivemos um backtrack
			current = current[:len(current)-1]
			used[i] = false
		}
	}

	backtrack()

	return res
}
