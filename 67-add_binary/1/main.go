package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(addBinary("11", "1"))
}

// big O: O( max(len (a), len(b)))
// Space Complexity:  O( max(len (a), len(b)))

func addBinary(a string, b string) string {
	var (
		lenA  = len(a)
		lenB  = len(b)
		carry int
		res   strings.Builder
	)

	for lenA > 0 && lenB > 0 {
		// primeiro somamos os dois números a direita
		// aqui precisamos somar o carry, pois é o resultado da soma das unidades
		tmp := int(a[lenA-1]-'0') + int(b[lenB-1]-'0') + carry

		// res nessa casa será o resto da divisão por 2
		res.WriteByte(byte(tmp%2 + '0')) // -> convertendo para caractere
		// 0 % 2 -> 0
		// 1 % 2 -> 1
		// 2 % 2 -> 0
		// 3 % 2 -> 1

		// se for par (0, 2 - ou seja, 00, 01), não tem carry
		// se for imapar (1, 3 - ou seja 01 e 11), tem carry
		carry = tmp / 2
		lenA--
		lenB--
	}

	if lenA == 0 {
		for lenB > 0 {
			tmp := int(b[lenB-1]-'0') + carry

			res.WriteByte(byte(tmp%2 + '0')) // -> convertendo para caractere
			carry = tmp / 2
			lenB--
		}
	}

	if lenB == 0 {
		for lenA > 0 {
			tmp := int(a[lenA-1]-'0') + carry
			res.WriteByte(byte(tmp%2 + '0')) // -> convertendo para caractere
			carry = tmp / 2
			lenA--
		}
	}

	// pode sobrar carry
	if carry == 1 {
		res.WriteByte(byte(carry) + '0')
	}

	result := res.String()
	var reversed strings.Builder

	for i := len(result) - 1; i >= 0; i-- {
		reversed.WriteByte(result[i])
	}

	return reversed.String()
}
