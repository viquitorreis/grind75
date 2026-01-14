package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return prices[0]
	}

	buy, sell := 0, 1
	profit := 0

	for buy < len(prices) && sell < len(prices) {
		if prices[sell] > prices[buy] {
			profit = max(profit, prices[sell]-prices[buy])
		} else {
			buy = sell
		}

		sell++
	}

	return profit
}
