package main

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	println(maxProfit(prices))
}

// prices[i] -> price of a stock on a single day
// return the maximum profit
func maxProfit(prices []int) int {
	minPrice := prices[0]
	prof := 0

	max := func(i, j int) int {
		if i > j {
			return i
		}

		return j
	}

	for i := 1; i <= len(prices)-1; i++ {
		if prices[i] > minPrice {
			prof = max(prof, prices[i]-minPrice)
		} else {
			minPrice = prices[i]
		}
	}

	return prof
}
