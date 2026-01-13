package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}

// essential:
// 		char have even occurances: palindrome
//		char have odd occurances: can be palindrome with a different char at the middle
//
//	.:. a palindrome cant have more than one odd occur character

// for this exercise we could have it in any order

func longestPalindrome(s string) int {
	occurances := Occurances(s)
	longest := 0
	for _, occurance := range occurances {
		longest += occurance / 2 * 2 // pegamos a quantidade par de ocorrencias
	}

	if longest < len(s) {
		longest++
	}

	return longest
}

func Occurances(s string) map[byte]int {
	occur := make(map[byte]int)
	for i := range s {
		occur[s[i]]++
	}
	return occur
}
