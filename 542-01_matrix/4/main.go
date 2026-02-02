package main

import "math"

func main() {

}

type Coordinate struct {
	x, y int
}

func updateMatrix(mat [][]int) [][]int {
	graph := make([][]int, len(mat))
	q := []*Coordinate{}
	for i := range mat {
		graph[i] = make([]int, len(mat[0]))
		for j := range mat[0] {
			if mat[i][j] == 0 {
				graph[i][j] = 0
				q = append(q, &Coordinate{
					x: i,
					y: j,
				})
			} else {
				graph[i][j] = math.MaxInt
			}
		}
	}

	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, dir := range dirs {
			newX := curr.x + dir[0]
			newY := curr.y + dir[1]

			if newX >= 0 && newX < len(graph) && newY >= 0 && newY < len(graph[0]) {
				if graph[newX][newY] == math.MaxInt {
					graph[newX][newY] = graph[curr.x][curr.y] + 1
					q = append(q, &Coordinate{
						x: newX,
						y: newY,
					})
				}
			}
		}
	}

	return graph
}

func dfs(graph [][]int, curr Coordinate) {
	if curr.x < 0 || curr.y < 0 || curr.x > len(graph) || curr.y > len(graph[0]) {
		return
	}

	if graph[curr.x][curr.y] != 0 {

	}
}
