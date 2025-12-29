package main

import "log"

func main() {
	nums := []int{2, 3, 6}
	target := 9
	twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	// idx -> num
	seen := make(map[int]int, 0)
	for i := 0; i <= len(nums)-1; i++ {
		curr := nums[i]
		diff := target - curr
		if k, ok := seen[diff]; ok {
			log.Printf("k: %d - ok: %t", k, ok)
			return []int{k, i}
		}

		log.Printf("seen[curr] %d - i %d", curr, i)

		seen[curr] = i
	}

	return []int{}
}
