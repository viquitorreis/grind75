package main

import "math"

func main() {

}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	// iteramos em todas as coins
	for _, coin := range coins {
		// iteramos nos valores, começámos de coin apenas para otimizar iterações
		for i := coin; i <= amount; i++ {
			// se o anterior nao for max inf
			if dp[i-coin] != math.MaxInt {
				// elemento atual vai ser o minimo entre o elemento atual e o anterior
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}

	return dp[amount]
}
