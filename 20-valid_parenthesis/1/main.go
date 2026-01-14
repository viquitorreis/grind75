package main

import "log"

func main() {
	pars1 := "()"
	pars2 := "()[]{}"
	pars3 := "(]"
	println(isValid(pars1)) // true
	println(isValid(pars2)) // true
	println(isValid(pars3)) // false
}

func isValid(s string) bool {
	pair := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	var stack []rune
	for _, char := range s {
		switch char {
		// inicio do match de parenteses
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			// não tem matching - elemento de fechamento sem inicio na stack
			if len(stack) == 0 {
				return false
			}

			topIdx := len(stack) - 1
			topElement := stack[topIdx]
			log.Printf("pair[char]: %s - char: %s\n", string(pair[char]), string(char))
			// checa se o ultimo elemento da stack é um elemento de inicio - {, [, (
			// se não for retorna false
			if topElement != pair[char] {
				return false
			}

			// pop - elemento tem pair matching
			stack = stack[:topIdx]
		}
	}

	return len(stack) == 0
}
