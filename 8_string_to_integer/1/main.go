package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d", myAtoi("   +0 123"))
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 1. primeiro remove qualquer espaço antes de chegar no número em si, atualizando a string
	i := 0
	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}
	s = s[i:]

	// 2. Encontrando se é soma ou subtração
	flag := 1
	if len(s) > 0 {
		switch s[0] {
		case '-':
			flag = -1
			fallthrough
		case '+':
			s = s[1:] // atualiza string
		}
	}

	if len(s) == 0 {
		return 0
	}

	// 3. remove leading zeros
	i = 0
	for i < len(s) && s[i] == '0' {
		i++
	}

	// só tem zeros
	if i >= len(s) {
		return 0
	}

	// 3. convertemos os numeros
	num := 0
	for _, char := range s {
		asciiVal := int(char) - '0'
		if asciiVal < 0 || asciiVal > 9 {
			break
		}

		num = num*10 + asciiVal
		if num > math.MaxInt32 {
			if flag == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
	}

	return num * flag
}
