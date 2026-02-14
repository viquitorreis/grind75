package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	trapped := 0
	if len(height) == 0 {
		return trapped
	}

	left, right := 0, len(height)-1
	minL, minR := 0, 0
	for left < right {
		if height[left] <= height[right] {
			if minL < height[left] {
				minL = height[left]
			} else {
				trapped += minL - height[left]
			}

			left++
		} else {
			if minR < height[right] {
				minR = height[right]
			} else {
				trapped += minR - height[right]
			}

			right--
		}
	}

	return trapped
}
