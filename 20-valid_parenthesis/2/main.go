package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
}

func isValid(s string) bool {
	pair := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// elemenetos de inicio simplesmente mandamos para a stack
			stack = append(stack, char)
		case ')', '}', ']':
			// 1. a stack NAO PODE comecar com um elemento de fechamento
			//		pois nao vai ter como jamais fechar ele -> achar um par
			if len(stack) == 0 {
				return false
			}

			// 2. como achamos um elemento de fechamento, o ultimo elemento precisa
			// 		ser prioritamente um de ABERTURA que dê match ao atual
			if stack[len(stack)-1] != pair[char] {
				return false
			}

			// 3. deu match, portanto tiramos da stack o ultimo elemento (o de abertura)
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
