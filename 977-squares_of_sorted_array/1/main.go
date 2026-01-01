package main

import "fmt"

func main() {
	// nums := []int{-4, -1, 0, 3, 10}
	// nums := []int{0, 1, 9, 16, 100}
	nums := []int{-1}
	sortedSquares(nums)
	printArr(nums)
}

// BIG O
//		- O(n²)
//			Squaring all elements: O(n)
//			Bubble Sort: O(n²) - not great, since for each elements can check each other one
//
// Space complexity
//		- O(1)
//			All operations are in place

func sortedSquares(nums []int) []int {
	for i, num := range nums {
		nums[i] = num * num
	}

	return sort(nums)
}

// bubble sort
func sort(nums []int) []int {
	slow, fast := 0, 1
	for slow < len(nums) {
		if fast < len(nums) && slow < len(nums) && nums[slow] > nums[fast] {
			temp := nums[fast]
			nums[fast] = nums[slow]
			nums[slow] = temp
		}

		fast++
		if fast >= len(nums) {
			slow++
			fast = slow + 1
		}
	}

	return nums
}

func printArr(arr []int) {
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}

	println()
}
