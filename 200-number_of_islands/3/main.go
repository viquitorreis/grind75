package main

func main() {

}

func numIslands(grid [][]byte) int {
	totalIslands := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '1' {
				continue
			}

			totalIslands++

			dfs(grid, i, j)
		}
	}

	return totalIslands
}

func dfs(grid [][]byte, x, y int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != '1' {
		return
	}

	grid[x][y]++

	dfs(grid, x-1, y)
	dfs(grid, x+1, y)
	dfs(grid, x, y-1)
	dfs(grid, x, y+1)
}
