package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}

	// 0 index
	minim := make([]int, amount+1)
	for i := range minim {
		minim[i] = math.MaxInt
	}
	minim[0] = 0

	for _, coin := range coins {
		// para cada um dos coins, vamos sempre iniciar a contagem no indice do coin
		for i := coin; i <= amount; i++ {
			// qual é o número mínimo necessário de moedas para formar o valor i?
			// minim[0] = 0 (zero moedas para formar valor 0)
			// minim[1] = mínimo de moedas para formar valor 1
			// minim[2] = mínimo de moedas para formar valor 2
			// ...
			// minim[6249] = mínimo de moedas para formar valor 6249
			if minim[i-coin] != math.MaxInt {
				minim[i] = min(minim[i], minim[i-coin]+1)
			}
		}
	}

	if minim[amount] == math.MaxInt {
		return -1
	}

	return minim[amount]
}
