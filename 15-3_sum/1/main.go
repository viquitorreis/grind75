package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("%v\n", threeSum(nums))
}

// explicação:

// https://medium.com/codex/golang-with-leetcode-3sum-af8c9d6d27db

// approach:
//		fazemos um slice 2D para armazenar os resultados
//		sort do array de input
//		loop até i == size - 2 -> i vai ser o ponteiro mais a esquierda
//
// 		for:
//			checar se i =0, se i > 0 e se os números não são iguais (nums[i] != nums[i-1])
//
//			pegamos low, high e sum
//				low -> ponteiro mais a esquerda logo depois de i
//				high -> high vai ser o ponteiro mais a direita, localizado no ultimo elemento do array
//
//			subtraimos: nums[i] - 0 = sum
//				pegamos os dois ponteiros low and high e adicionamos nessa soma

//			for:
//				low < high
//				outro loop para encontrarmos combinações de elementos que adicionados, serão a soma que queremos.
//
// 				loops para tirar duplicatas:
//					1: continua até que low < high e o elemento na posição atual == elemento proximo.
//						movemos o ponteiro esquerdo até que um elemento unico seja encontrado

//					2: mesma logica para o ponteiro high, mas movemos para trás

//				depois que os dois loops terminarem e pularmos as duplicatas, movemos os dois ponteiros para frente

//				se a soma em low e high não for a soma correta, checamos se o resultado é maior que a soma que queremo.
// 					Se for maior, movemos o high para trás
//					a razão que fazemos isso é porque sorteamos o array em ASC

//					Se for menor, movemos o ponteiro esquerdo para frente

// return all triplets [nums[i], nums[j], nums[k]
// such that i != j
// i != k
// j != k
// and nums[i] + nums[j] + nums[k] == 0
func threeSum(nums []int) [][]int {
	triplets := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			low := i + 1
			high := len(nums) - 1
			sum := 0 - nums[i]

			for low < high {
				if nums[low]+nums[high] == sum {
					triplets = append(triplets, []int{nums[i], nums[low], nums[high]})

					// achando elementos unicos a partir do low
					for low < high && nums[low] == nums[low+1] {
						low++
					}

					// achando elementos unicos a partir do high
					for high > low && nums[high] == nums[high-1] {
						high--
					}

					low++
					high--
				} else if nums[low]+nums[high] > sum {
					high--
				} else {
					low++
				}
			}
		}
	}

	return triplets
}
