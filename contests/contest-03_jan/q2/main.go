package main

func main() {
	// arr := []int{4}
	arr := []int{6, 5, 4, 5}
	// arr := []int{1, 6}
	// k := 6
	// k := 9
	k := 7
	println(minLength(arr, k))
}

// Problema: quer o tamanho mínimo de elementos que conseguem
// se somados um resultado maior que a soma de todos elementos do array

// Ex:
// nums = [6,5], k = 9
//	 	sum = 11
// 		11 >= 9
//		Answer: minLength = 2, we only need the two elements

// Pattern: Sliding Window
//		Why?
//		sliding -> which unique numbers exist in this array? Add its value it on total sum
//			For each pass, we check if we have found the total length

func minLength(nums []int, k int) int {
	// num -> exists
	// unique := make(map[int]bool)
	windowFreq := make(map[int]int)
	distinctSum := 0
	minWinLen := len(nums) + 1 // começamos em um valor impossível

	left := 0

	for right := 0; right < len(nums); right++ {
		// adicionamos nums[right] na janela
		if windowFreq[nums[right]] == 0 {
			distinctSum += nums[right]
		}
		windowFreq[nums[right]]++

		// diminuimos a janela enquanto é válida
		for distinctSum >= k {
			// se acharmos uma janela válida, traqueamos seu tamanho
			// janela válida -> sempre que > 1
			currentWinLen := right - left + 1
			if currentWinLen < minWinLen {
				minWinLen = currentWinLen
			}

			// quando formos remover da janela, diminuimos a frequencia do numero...
			windowFreq[nums[left]]--
			if windowFreq[nums[left]] == 0 { // se for a ultima ocorrencia, removemos da soma
				distinctSum -= nums[left]
			}
			left++
		}

	}

	if minWinLen == len(nums)+1 {
		return -1
	}

	return minWinLen
}
