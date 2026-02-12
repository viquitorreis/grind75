package main

import (
	"math"
)

func main() {

}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	window := make(map[byte]int)

	needCount := len(need)

	// caracteres unicos que satisfazemos na janela atual
	formed := 0

	minLen := math.MaxInt
	left := 0
	minL, minR := 0, 0

	for right := range len(s) {
		char := s[right]
		window[char]++

		// elemento é necessário e a contagem é a correta que precisamos
		if count, exists := need[char]; exists && count == window[char] {
			formed++

			// vamos tentar encolher a janela, enquanto for válida (deve existir em need)
			for formed == needCount && left <= right {
				currentLen := right - left + 1

				// achamos uma janela menor. Atualizamos os minimos
				if currentLen < minLen {
					minLen = currentLen
					minL = left
					minR = right
				}

				// vam0os tentar encolher a janela
				window[s[left]]--

				if count, exists := need[s[left]]; exists && count != window[s[left]] {
					formed--
				}

				left++
			}
		}
	}

	if minLen > len(s) {
		return ""
	}

	return s[minL : minR+1]
}
