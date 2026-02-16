package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func summaryRanges(nums []int) []string {
	res := []string{}
	if len(nums) == 0 {
		return res
	}

	start := 0
	for i := 1; i <= len(nums); i++ {
		// só entramos para adicionar quando detectamos o fim de uma sequencia
		// seja por final do array, seja pelo numero atual não ser subsequente ao anterior
		if i == len(nums) || nums[i] != nums[i-1]+1 {
			// sequencia de um número só, o número anterior é igual a esse
			if nums[start] == nums[i-1] {
				res = append(res, strconv.Itoa(nums[start]))
			} else {
				res = append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[i-1]))
			}

			// i no fechamento de uma sequencia é exatamente o indice do primeiro elemento da proxima sequencia
			start = i
		}
	}

	return res
}
