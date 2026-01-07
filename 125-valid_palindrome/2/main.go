package main

import "unicode"

func main() {
	s := "0P"
	println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	forward, backward := []rune{}, []rune{}

	for i, j := 0, len(s)-1; i < len(s) && j >= 0; i, j = i+1, j-1 {
		left := rune(s[i])
		if unicode.IsLetter(left) || unicode.IsNumber(left) {
			forward = append(forward, unicode.ToLower(left))
		}

		right := rune(s[j])
		if unicode.IsLetter(right) || unicode.IsNumber(left) {
			backward = append(backward, unicode.ToLower(right))
		}
	}

	println(string(forward))
	println(string(backward))

	return string(forward) == string(backward)
}
