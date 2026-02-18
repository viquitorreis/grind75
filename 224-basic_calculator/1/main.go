package main

func main() {

}

func calculate(s string) int {
	num := 0
	sign := 1
	res := 0
	// na stack, vamos SEMPRE manter o operador + resultado até então
	stack := make([]int, 0)

	for _, char := range s {
		if char-'0' >= 0 && char-'0' <= 9 {
			num = num*10 + int(char-'0')
		} else if char == '(' {
			// empurramos o resultado atual e o singal atual para a stack
			stack = append(stack, res, sign)
			sign = 1
			res = 0
		} else if char == ')' {
			res += num * sign
			// sempre pegamos em pares da stack -> sinal anterior e resultado anterior
			res *= stack[len(stack)-1]
			res += stack[len(stack)-2]

			// removemos o par consumido
			stack = stack[:len(stack)-2]

			// resetamos o num e o operando
			num = 0
			sign = 1
		} else if char == '+' {
			res += sign * num
			sign = 1
			num = 0
		} else if char == '-' {
			res += sign * num
			sign = -1
			num = 0
		}
	}

	res += sign * num

	return res
}
