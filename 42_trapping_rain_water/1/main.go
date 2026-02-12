package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2}))
}

func trap(height []int) int {
	// nada antes é usado para represar agua.
	// nada depois é usado
	for len(height) <= 1 {
		return 0
	}

	l, r := 0, len(height)-1
	maxL, maxR := 0, 0
	trapped := 0

	for l < r {
		if height[l] <= height[r] {
			// 1. quando max left é fator limitante, aumentamos o muro à esquerda
			// 		pois não tem como represar a agua nesse intervalo
			// 2. caso contrário aumentamos a quantidade de água represada
			if height[l] >= maxL {
				maxL = height[l]
			} else {
				trapped += maxL - height[l]
			}

			l++
		} else {
			// 1. max right é fator limitante, aumentamos o muro a direita.
			// 2. caso contrario somamos na agua represada
			if height[r] >= maxR {
				maxR = height[r]
			} else {
				trapped += maxR - height[r]
			}

			r--
		}
	}

	return trapped
}
