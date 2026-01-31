package main

func main() {

}

func uniquePaths(m int, n int) int {
	if m == 0 && n == 0 {
		return 0
	}

	grid := make([][]int, m)
	for i := range len(grid) {
		grid[i] = make([]int, n)
	}

	// first row all 1
	for i := range len(grid) {
		grid[i][0] = 1
	}

	// firs column all 1
	for i := range len(grid[0]) {
		grid[0][i] = 1
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[m-1][n-1]
}
