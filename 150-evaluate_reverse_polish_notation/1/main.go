package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	// fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"4", "3", "-"}))
}

func evalRPN(tokens []string) int {
	var stk stack

	for _, s := range tokens {
		switch val := s; s {
		case "+", "-", "*", "/":
			if len(stk) == 0 {
				return 0
			}

			switch val {
			case "+":
				stk.Push(stk.Pop() + stk.Pop())
			case "-":
				el1 := stk.Pop()
				el2 := stk.Pop()
				stk.Push(el2 - el1)
			case "*":
				stk.Push(stk.Pop() * stk.Pop())
			case "/":
				el1 := stk.Pop()
				el2 := stk.Pop()
				stk.Push(el2 / el1)
			}
		default:
			num, _ := strconv.Atoi(s)
			stk.Push(num)
		}
	}

	return stk.Pop()
}

type stack []int

func (s *stack) Push(x int) {
	*s = append(*s, x)
}

func (s *stack) Pop() int {
	size := len(*s)
	elem := (*s)[size-1]
	*s = (*s)[:size-1]
	return elem
}
