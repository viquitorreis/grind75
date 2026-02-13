package main

import "fmt"

func main() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}

func trap(height []int) int {
	left, right := 0, len(height)-1
	trapped := 0
	maxL, maxR := 0, 0

	for left < right {
		if height[left] <= height[right] {
			if height[left] >= maxL {
				maxL = height[left]
			} else {
				trapped += maxL - height[left]
			}

			left++
		} else {
			if height[right] >= maxR {
				maxR = height[right]
			} else {
				trapped += maxR - height[right]
			}

			right--
		}
	}

	return trapped
}
