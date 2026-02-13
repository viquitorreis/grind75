package main

import "fmt"

func main() {
	fmt.Println()
}

func uniquePaths(m int, n int) int {
	grid := make([][]int, m)
	for i := range m {
		grid[i] = make([]int, n)
	}

	// primeira linhas é tudo 1
	for i := range len(grid) {
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

	return grid[len(grid)-1][len(grid[0])-1]
}
