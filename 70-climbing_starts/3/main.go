package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev, curr := 1, 2
	for i := 3; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}

	return curr
}
