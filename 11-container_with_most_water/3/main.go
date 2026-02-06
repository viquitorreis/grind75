package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right, maxWater := 0, len(height)-1, 0

	for left < right {
		width := right - left
		high := 0

		// fmt.Println("width left right", width, right, left)

		if height[left] < height[right] {
			high = height[left]
			left++
		} else {
			high = height[right]
			right--
		}

		maxWater = max(maxWater, width*high)
	}

	return maxWater
}
