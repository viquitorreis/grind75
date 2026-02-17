package main

func main() {

}

func longestPalindrome(s string) int {
	// palindromo
	// funciona igual um espelho, a letra da ponta é igual da outra ponta
	// só pode ter uma letra não igual, quando for de tamanho impar e for a letra do meio
	if len(s) == 0 {
		return 0
	}

	count := make(map[rune]int)
	for _, char := range s {
		count[char]++
	}

	num := 0
	for _, occurrance := range count {
		num += occurrance / 2 * 2
	}

	if num < len(s) {
		return num + 1
	}

	return num
}
