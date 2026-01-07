package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := "A man, a plan, a canal: Panama"
	s2 := "race a car"
	println(isPalindrome(s1))
	println(isPalindrome(s2))
}

// palindromo -> se converter para lower case, e tirar todos caracteres não alfanuméricos
// a palavra / frase deve continuar a mesma de trás para frente e vice versa

func isPalindrome(s string) bool {
	var (
		normal, backwards string
	)

	for _, char := range s {
		if unicode.ToLower(char) >= 'a' && unicode.ToLower(char) <= 'z' {
			fmt.Printf(" %s - ", string(unicode.ToLower(char)))
			normal += string(unicode.ToLower(char))
		}
	}

	println()
	println("normal: ", normal)
	println("len normal: ", len(normal))

	for i := len(normal) - 1; i >= 0; i-- {
		char := rune(normal[i])
		fmt.Printf(" %s - ", string(unicode.ToLower(char)))
		backwards += string(char)
	}

	println("backwards: ", backwards)
	println("len backwards: ", len(backwards))

	return strings.EqualFold(normal, backwards)
}
