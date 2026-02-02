package main

func main() {

}

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	maxArea, start, end := 0, 0, len(height)-1

	for start < end {
		width := end - start
		high := 0

		// sempre pegamos a menor altura para calcular, pois não pode transbordar
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		area := high * width
		maxArea = max(maxArea, area)
	}

	return maxArea
}
