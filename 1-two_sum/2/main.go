package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Printf("%v", twoSum(nums, 9))
}

// Big O:
//		O(n) - we will traverse the numbers at most once, so its the length of n. The rest are basic comparisons and allocations

// Space Complexity
//		O(n) - the hash map created will store at most n numbers, being the amount of elements on nums.
//			The rest are only basic allocations - variables O(1)

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if idx, ok := seen[diff]; ok {
			return []int{idx, i}
		}

		seen[num] = i
	}

	return []int{}
}
