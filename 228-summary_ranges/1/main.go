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
		// sequencia continua até chegar ao fim do array ou quando o próximo número quebra a consecutividade
		if i == len(nums) || nums[i] != nums[i-1]+1 {
			// numero unico
			if nums[start] == nums[i-1] {
				res = append(res, strconv.Itoa(nums[start]))
			} else {
				// sequ3encia de numeros
				res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(nums[start]), strconv.Itoa(nums[i-1])))
			}
			start = i
		}
	}

	return res
}
