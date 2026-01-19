package main

import "fmt"

func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 1},
		{0, 1, 2},
		// {1, 1, 0},
		// {0, 1, 1},
	}
	fmt.Println(orangesRotting(grid))
}

type Coordinate struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}

	rottens := 0
	normal := 0

	res := make([][]int, len(grid))
	q := []Coordinate{}
	for i := range grid {
		res[i] = make([]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				res[i][j] = 2
				q = append(q, Coordinate{
					x: i,
					y: j,
				})
				rottens++
			} else if grid[i][j] == 1 {
				res[i][j] = 1
				normal++
			}
		}
	}

	if normal == 0 {
		return 0
	}

	if rottens == 0 {
		return -1
	}

	directions := [][]int{
		{0, 1},  // direita
		{0, -1}, // esquerda
		{1, 0},  // baixo
		{-1, 0}, // cima
	}

	minutes := 0

	for len(q) > 0 {
		levelSize := len(q)

		// PRECISAMOS PROCESSAR POR NIVEL VISITADO (BFS - tamanho da fila - cada elemento da fila, deve somar apenas 1)
		// a cada nivel processado, somamos, começando por 0

		for i := 0; i < levelSize; i++ {
			curr := q[0]
			q = q[1:]

			for _, dir := range directions {
				newX := curr.x + dir[0]
				newY := curr.y + dir[1]

				if newX >= 0 && newY >= 0 && newX < len(grid) && newY < len(grid[0]) {
					if res[newX][newY] == 1 {
						res[newX][newY] = 2
						q = append(q, Coordinate{
							x: newX,
							y: newY,
						})
					}
				}
			}
		}

		if len(q) > 0 {
			minutes++
		}
	}

	for i := range res {
		for j := 0; j < len(res[0]); j++ {
			if res[i][j] == 1 {
				return -1
			}
		}
	}

	return minutes
}
