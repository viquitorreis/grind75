package main

import (
	"fmt"
)

func main() {
	fmt.Println(orangesRotting([][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
		// {1, 0, 1},
		// {1, 2},
	}))
}

type Coordinate struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	q := []*Coordinate{}
	oranges := 0
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == 2 {
				oranges++
				q = append(q, &Coordinate{
					x: i,
					y: j,
				})
			} else if grid[i][j] == 1 {
				oranges++
			}
		}
	}

	if oranges == len(q) {
		return 0
	}

	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	minutes := 0
	totalProcessed := len(q)

	for len(q) > 0 {
		// old queue?
		size := len(q)
		flag := false

		for range size {
			curr := q[0]
			q = q[1:]

			for _, dir := range dirs {
				newX := curr.x + dir[0]
				newY := curr.y + dir[1]

				if newX >= 0 && newY >= 0 && newX < len(grid) && newY < len(grid[0]) {
					fmt.Println("newX, newY", newX, newY)
					if grid[newX][newY] == 1 {
						totalProcessed++
						flag = true
						grid[newX][newY] = 2
						q = append(q, &Coordinate{
							x: newX,
							y: newY,
						})
					}
				}
			}
		}

		if flag {
			minutes++
		}
	}

	if totalProcessed != oranges {
		return -1
	}

	return minutes
}
