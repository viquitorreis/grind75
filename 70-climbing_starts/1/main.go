package main

import "fmt"

func main() {
	fmt.Println("missing steps: ", climbStairs(4))
}

func climbStairs(n int) int {
	// base case
	if n <= 2 {
		return n
	}

	// dynamic programming
	prev1, prev2 := 1, 2
	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		// previous mudaram agora
		prev1 = prev2
		prev2 = current
	}

	return prev2
}
