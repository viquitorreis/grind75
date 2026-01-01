package main

func main() {
	println("max profit is: ", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	// println("max profit is: ", maxProfit([]int{7, 6, 4, 3, 1}))
	// println("max profit is: ", maxProfit([]int{2, 1, 4}))
	// println("max profit is: ", maxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
	// println("max profit is: ", maxProfit([]int{1, 2}))
}

// PROBLEMA ==========================
// temos um array prices onde o prices[i] é o preço de uma ação no dia i^th
// queremos maximizar o lucro ao escolher um único dia para comprar uma ação e vender em outro dia futuro para vender
// precisamos retornar o o lucro máximo que podemos obter. Caso nao dê para lucrar retornar 0

func maxProfit(prices []int) int {
	buy, maxProfit := 0, 0

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	for sell := 1; sell <= len(prices)-1; sell++ {
		// se o preco de venda > preco de compra, temos um POTENCIAL lucro maximo
		if prices[sell] > prices[buy] {
			maxProfit = max(prices[sell]-prices[buy], maxProfit)
		} else {
			buy = sell
		}
	}

	return maxProfit
}
