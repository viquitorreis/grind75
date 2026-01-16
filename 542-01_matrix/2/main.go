package main

import "math"

func main() {

}

type Coordinate struct {
	x, y int
}

func updateMatrix(mat [][]int) [][]int {
	// 1. Preenche uma nova matriz inteira com os valores 0
	res := make([][]int, len(mat))
	q := []*Coordinate{}
	for i := 0; i < len(mat); i++ {
		res[i] = make([]int, len(mat[0]))

		for j := 0; j < len(mat[0]); j++ {

			if mat[i][j] == 0 {
				res[i][j] = 0
				q = append(q, &Coordinate{
					x: i,
					y: j,
				})
			} else {
				res[i][j] = math.MaxInt
			}
		}
	}

	directions := [][]int{
		{0, 1},  // cima
		{0, -1}, // baixo
		{1, 0},  // direita
		{-1, 0}, // esquerda
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		// precisamos achar o primeiro elemento mais próximo do zero
		for _, dir := range directions {
			newX := curr.x + dir[0] // nova direcao no eixo x
			newY := curr.y + dir[1] // nova direcao no eixo y

			if newX >= 0 && newX < len(mat) && newY >= 0 && newY < len(res[0]) {
				if res[newX][newY] == math.MaxInt {
					res[newX][newY] = res[curr.x][curr.y] + 1
					q = append(q, &Coordinate{
						x: newX,
						y: newY,
					})
				}
			}
		}
	}

	return res
}
