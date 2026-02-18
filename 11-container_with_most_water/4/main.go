package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	trapped := 0
	if len(height) == 0 {
		return trapped
	}

	left := 0
	right := len(height) - 1
	for left < right {
		h := min(height[right], height[left])

		w := right - left
		area := w * h

		if area > trapped {
			trapped = area
		} else {
			if height[left] < height[right] {
				left++
			} else {
				right--
			}
		}
	}

	return trapped
}
