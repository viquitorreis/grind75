package main

import (
	"fmt"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	window := make(map[byte]int)
	formed, required := 0, len(need)

	// minLen inicializa com um valor bem grande
	left, minLen, minLeft, minRight := 0, len(s)+1, 0, 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		// adicionamos o caractere na janela e aumentamos sua contagem
		window[char]++

		// o caractere é necessário e temos a quantidade certa. Aumentamos a contagem
		if count, exists := need[char]; exists && count == window[char] {
			formed++
		}

		// loop interno para diminuir a janela valida
		for formed == required && left <= right {
			currentLen := right - left + 1

			if currentLen < minLen {
				minLen = currentLen
				minLeft = left
				minRight = right
			}

			leftChar := s[left]
			window[leftChar]--

			// se temos menos que a quantidade necessaria diminuimos formed
			if count, exists := need[leftChar]; exists && window[leftChar] < count {
				formed--
			}

			left++
		}
	}

	if minLen > len(s) {
		return ""
	}

	return s[minLeft : minRight+1]
}
