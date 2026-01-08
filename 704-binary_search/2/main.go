package main

import "fmt"

func main() {
	// nums := []int{-1, 0, 3, 5, 9, 12}
	// fmt.Println("search: ", search(nums, 9))
	nums := []int{-1, 0, 3, 5}
	fmt.Println("search: ", search(nums, 3))
}

// Complexity Analisys

// Big O:
//		O (log n)
//		Divide and Conquer -> Ex: 16 → 8 → 4 → 2 → 1 ->  2^4 = 16, so log₂(16) = 4\

// Space Complexity:
//		O(log n) -> RECURSION STACk!!!!!!

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, min, max, target int) int {
	if min <= max {
		midd := (min + max) / 2

		if nums[midd] == target {
			return midd
		} else if nums[midd] > target {
			return binarySearch(nums, min, midd-1, target)
		} else {
			return binarySearch(nums, midd+1, max, target)
		}
	}

	return -1
}
