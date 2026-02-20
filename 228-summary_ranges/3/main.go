package main

import (
	"strconv"
)

func main() {

}

func summaryRanges(nums []int) []string {
	// precisamos checar se os elementos são contíguos
	res := []string{}
	if len(nums) == 0 {
		return res
	}

	start := 0
	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i] != nums[i-1]+1 {
			if nums[i-1] == nums[start] {
				res = append(res, strconv.Itoa(nums[start]))
			} else {
				res = append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[i-1]))
			}

			start = i
		}

	}

	return res
}
