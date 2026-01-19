package main

func main() {

}

type Coordinate struct {
	x, y int
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	islands := 0

	for i := range grid {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				islands++
				traverseIsland(grid, i, j)
			}
		}
	}

	return islands
}

func traverseIsland(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	grid[i][j]++

	traverseIsland(grid, i-1, j)
	traverseIsland(grid, i+1, j)
	traverseIsland(grid, i, j-1)
	traverseIsland(grid, i, j+1)
}
