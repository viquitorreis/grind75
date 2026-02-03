package main

func main() {

}

func maxArea(height []int) int {
	left, right, maxHeight := 0, len(height)-1, 0

	for left < right {
		high := min(height[left], height[right])

		width := right - left

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

		maxHeight = max(maxHeight, high*width)
	}

	return maxHeight
}
