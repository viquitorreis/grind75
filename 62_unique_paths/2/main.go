package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	// preenchemos tudo da linha de cima com 1
	grid := make([][]int, m)
	for i := range len(grid) {
		grid[i] = make([]int, n)
		grid[i][0] = 1
	}

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
