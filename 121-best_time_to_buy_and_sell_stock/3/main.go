package main

func main() {
	println("max profit is: ", maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

// Big O:
// 		O(n) - n being the number of prices. The rest are only basic operations O(1) such as variables and conditionals

// Space Complexity
// 		O(1) - we are using only basic memory space in memory, variables

func maxProfit(prices []int) int {
	buy, sell := 0, 1
	profit := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	for buy < len(prices) && sell < len(prices) {
		if prices[sell] > prices[buy] {
			profit = max(prices[sell]-prices[buy], profit)
		} else {
			buy = sell
		}
		sell++
	}

	return profit
}
