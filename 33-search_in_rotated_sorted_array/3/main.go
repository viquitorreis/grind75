package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 0))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) >> 1

		fmt.Println("mid:", mid, nums[mid])

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= target {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[right] && nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
