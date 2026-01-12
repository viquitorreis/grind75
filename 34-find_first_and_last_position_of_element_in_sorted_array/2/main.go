package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{1, 1, 2}, 1))
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
			// fmt.Printf("mid: %d - nums[mid]: %d - start: %d - finish: %d\n", mid, nums[mid], start, finish)
			// precisamos encontrar a janela correta
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
