package main

import (
	"fmt"
)

func main() {
	// abccccdd
	fmt.Println(longestPalindrome("aaabbb"))
}

// case sensitive
func longestPalindrome(s string) int {
	longest := 0
	counter := Counter(s)

	for _, count := range counter {
		longest += (count / 2) * 2 // quantidades pares
	}

	// se nao usamos todas as letras disponiveis
	// PODEMOS COLOCAR UMA LETRA NO MEIO! Pois a soma do longest até esse ponto SEMPRE será par
	// portanto cabe letra no meio....
	if longest < len(s) {
		longest++
	}

	return longest
}

func Counter(s string) map[byte]int {
	counter := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}
	return counter
}
