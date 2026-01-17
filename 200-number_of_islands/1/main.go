package main

func main() {

}

func numIslands(grid [][]byte) int {
	// 1 -> LAND
	// 0 -> WATER
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	count := 0

	for row := range grid {
		for col := range len(grid[0]) {
			if grid[row][col] != '1' {
				continue
			}

			count++

			findIslands(grid, row, col)
		}
	}

	return count
}

func findIslands(grid [][]byte, row, col int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] != '1' {
		return
	}

	// somamos o elemento atual, pois só queremos visitar as ilhas
	grid[row][col]++

	findIslands(grid, row+1, col) // baixo
	findIslands(grid, row, col+1) // direita
	findIslands(grid, row-1, col) // cima
	findIslands(grid, row, col-1) // esquerda
}
