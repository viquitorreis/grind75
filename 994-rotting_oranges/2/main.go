package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		// {2, 1, 1},
		// {1, 1, 0},
		// {0, 1, 1},
		{1},
	}
	fmt.Println(orangesRotting(grid))
}

type Coordinate struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	rottensCount := 0
	totalOranges := 0
	q := []Coordinate{}
	rottens := make([][]int, len(grid))
	for i := range grid {
		rottens[i] = make([]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				rottens[i][j] = 2
				rottensCount++
				q = append(q, Coordinate{
					x: i,
					y: j,
				})
				totalOranges++
			} else if grid[i][j] == 1 {
				rottens[i][j] = 1
				totalOranges++
			}
		}
	}

	if totalOranges == 0 {
		return 0 // sem laranjar é 0 minutos
	}

	if rottensCount == 0 {
		return -1
	}

	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	minutes := 0

	for len(q) > 0 {
		// salvamos o tamanhoa da q AGORA
		levelSize := len(q)

		for i := 0; i < levelSize; i++ { // processamos TODO O NÍVEL ATUAL (bfs)
			curr := q[0]
			q = q[1:]

			// percorrer nos 4 cantos antes de somar
			for _, dir := range directions {
				newX := curr.x + dir[0]
				newY := curr.y + dir[1]

				fmt.Println(newX, newY)
				// fmt.Println("new x >= 0?", newX >= 0)
				// fmt.Println("visiting: ", rottens[newX][newY])
				fmt.Println("rottens now:", rottens)

				if newX >= 0 && newY >= 0 && newX < len(grid) && newY < len(grid[0]) && rottens[newX][newY] == 1 {
					rottens[newX][newY]++
					rottensCount++
					q = append(q, Coordinate{
						x: newX,
						y: newY,
					})
				}
			}

		}

		// se adicionou mais elementos na fila soma + 1 minuto
		if len(q) > 0 {
			minutes++
		}
	}

	fmt.Println("rottens and total:", rottensCount, totalOranges)

	if rottensCount != totalOranges {
		return -1
	}

	fmt.Println("minutes: ", minutes)

	return minutes
}
