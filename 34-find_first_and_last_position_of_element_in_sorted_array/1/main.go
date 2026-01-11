package main

import "fmt"

func main() {
	nums := []int{2, 2}
	fmt.Println(searchRange(nums, 2))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) >> 1

		if nums[mid] == target {
			start, finish := mid, mid
			// fmt.Printf("achamos target: %d - mid: %d - start: %d - finish: %d\n", target, mid, start, finish)
			// precisamos achar todos os números iguais ao target
			for start > 0 && nums[start-1] == target {
				start--
			}

			for finish < len(nums)-1 && nums[finish+1] == target {
				finish++
			}

			return []int{start, finish}
		} else if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return []int{-1, -1}
}
