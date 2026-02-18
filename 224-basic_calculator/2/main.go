package main

func main() {

}

func calculate(s string) int {
	sign := 1
	res := 0
	num := 0
	stack := []int{}

	for _, char := range s {
		if char-'0' >= 0 && char-'0' <= 9 {
			num = num*10 + int(char-'0')
			// res = res + num
		} else if char == '(' {
			stack = append(stack, sign, res)
			sign = 1
			// num = 0
			res = 0
		} else if char == ')' {
			// precisamos calcular o numero pendente
			res += num * sign

			stackSign := stack[len(stack)-2]
			resSign := stack[len(stack)-1]

			stack = stack[:len(stack)-2]

			res += resSign + (res * stackSign)

			sign = 1
			num = 0
		} else if char == '+' {
			res += num * sign
			sign = 1
			num = 0
		} else if char == '-' {
			res += num * sign
			sign = -1
			num = 0
		}
	}

	res += sign * num
	return res
}
