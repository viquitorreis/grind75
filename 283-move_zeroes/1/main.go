package main

import (
	"fmt"
)

func main() {
	// nums := []int{0, 1, 0, 3, 12}
	nums := []int{0, 1}
	printArr(nums)
	moveZeroes(nums)
	printArr(nums)
}

// move zeros to the end of the array

// Big O:
// O(n) - we are traversing the slice exactly once. The rest are only basic operations

// Space Complexity:
// O(1) - since only basic variables

func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}

	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] == 0 {
			if nums[fast] != 0 {
				swap(slow, fast, nums)
			} else {
				fast++
				continue
			}
		}

		slow++
		fast++
	}
}

func swap(i, j int, nums []int) []int {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
	fmt.Printf("temp: %d - nums[i]: %d - nums[j]: %d\n", temp, nums[i], nums[j])
	printArr(nums)

	return nums
}

func printArr(arr []int) {
	for _, el := range arr {
		fmt.Printf("%d ", el)
	}
	println()
}
