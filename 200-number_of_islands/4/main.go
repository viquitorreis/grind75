package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '1', '0', '1'},
		{'1', '1', '0', '0', '1'},
		{'1', '0', '0', '0', '1'},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	numIslands := 0

	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] != '1' {
				continue
			}

			numIslands++

			dfs(grid, i, j)
		}
	}

	return numIslands
}

func dfs(grid [][]byte, i, j int) {
	// fmt.Println("grid:", grid)
	// fmt.Println("i and j", i, j)
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	grid[i][j]++

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}
