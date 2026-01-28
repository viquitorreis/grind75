package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi(" -42"))
}

func myAtoi(s string) int {
	// ignorar ' ' ou +
	// ler o inteiro ao pular zeros a esquerda até que encontre um caractere que não seja digito ou o fim da string chegar
	// se nao achar nenhum digito, o resultado é 0
	// arredondar pra max inf 32 bits
	if len(s) == 0 {
		return 0
	}

	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	s = s[i:]

	fmt.Println("s: ", s)

	isNegative := false
	if len(s) > 0 {
		switch s[0] {
		case '-':
			isNegative = true
			fallthrough
		case '+':
			s = s[1:]
		}
	}

	if len(s) == 0 {
		return 0
	}

	i = 0
	for i < len(s) && s[i] == '0' {
		i++
	}

	// so tem zero
	if i >= len(s) {
		return 0
	}
	s = s[i:]

	fmt.Println("s 1: ", s)

	num := 0
	for _, char := range s {
		fmt.Println("char is: ", string(char))
		ascii := char - '0'
		if ascii < 0 || ascii > 9 {
			break
		}

		num = num*10 + int(ascii)
		if num > math.MaxInt32 {
			if isNegative {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
	}

	if isNegative {
		return -num - 1
	}

	return num
}
