package main

func main() {

}

type Coordinate struct {
	x, y int
}

func pacificAtlantic(heights [][]int) [][]int {
	res := [][]int{}
	if len(heights) == 0 || len(heights[0]) == 0 {
		return res
	}

	q := []Coordinate{{x: heights[0][0], y: heights[0][1]}}

	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for i := range len(heights) {
		for j := range len(heights[0]) {
			for _, dir := range dirs {
				newX := i + dir[0]
				newY := j + dir[1]

				if newX >= 0 && newY >= 0 && newX <= len(heights) && newY <= len(heights[0]) {
					// get min neighbor

				} else {
					// out of border, its ocean...
					res[i][j] = heights[i][j]
				}
			}
		}
	}
}

func getMinNeighbor(heights [][]int, dirs [][]int, x, y int) Coordinate {

}
