package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 6))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)-1
	start, end := 0, 0

	for left <= right {
		mid := (left + right) >> 1

		if nums[mid] == target {
			start = mid
			end = mid

			for start > 0 && nums[start-1] == target {
				start--
			}

			for end < len(nums)-1 && nums[end+1] == target {
				end++
			}
		}

		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if start == 0 && end == 0 && nums[0] != target {
		return []int{-1, -1}
	}

	return []int{start, end}
}
