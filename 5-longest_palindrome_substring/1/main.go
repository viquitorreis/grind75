package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("babee"))
}

func longestPalindrome(s string) string {
	// sliding window
	// longest palindrome -> posso ter inumeras letras com quantidade par de vezes
	//	apenas uma com quantidade ímpar de vezes (a do meio)
	if len(s) <= 1 {
		return s
	}

	var res string

	palindrome := func(s string, l, r int) string {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}

		//
		return s[l+1 : r]
	}

	for i := 0; i < len(s); i++ {
		first := palindrome(s, i, i)    // checando se valor ímpar é palindromo
		second := palindrome(s, i, i+1) // checando se valor par é palindromo

		if len(res) < len(first) {
			res = first
		}

		if len(res) < len(second) {
			res = second
		}
	}

	return res
}
