package main

func main() {

}

func longestPalindrome(s string) string {
	// palindromo é formado por string contígua
	if len(s) <= 1 {
		return s
	}

	palindrome := func(l, r int) string {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}

		return s[l+1 : r]
	}

	longest := ""

	for i := 0; i < len(s); i++ {
		first := palindrome(i, i)
		second := palindrome(i, i+1)

		if len(first) > len(longest) {
			longest = first
		}

		if len(second) > len(longest) {
			longest = second
		}
	}

	return longest
}
