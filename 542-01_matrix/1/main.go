package main

import (
	"fmt"
	"math"
)

func main() {
	m := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	fmt.Println(updateMatrix(m))
}

type Coordinate struct {
	x, y int
}

func updateMatrix(mat [][]int) [][]int {
	res := make([][]int, len(mat))
	q := []Coordinate{}

	for i := range len(mat) {
		res[i] = make([]int, len(mat[0]))

		for j := range len(mat[0]) {
			res[i][j] = math.MaxInt64
			if mat[i][j] == 0 {
				q = append(q, Coordinate{
					x: i, // posicação da linha
					y: j, // posição da coluna
				})
				res[i][j] = 0
			}
		}
	}

	// direcoes possiveis
	directions := [][]int{
		{-1, 0}, // cima
		{1, 0},  // baixo
		{0, 1},  // direita,
		{0, -1}, // esquerda
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		// precisamos percorrer o elemento atual em todas as direcoes possiveis
		// ou seja, para cada um dos vizinhos
		for _, dir := range directions {
			newX := curr.x + dir[0]
			newY := curr.y + dir[1]

			// checa se o vizinho está dentro dos limites
			if newX >= 0 && newX < len(mat) && newY >= 0 && newY < len(mat[0]) {
				// se nao foi visitada ainda, calculamos a distancia
				if res[newX][newY] == math.MaxInt64 {
					res[newX][newY] = res[curr.x][curr.y] + 1 // distancia do vizinho
					q = append(q, Coordinate{
						x: newX,
						y: newY,
					}) // adicionamos o vizinho na queue
				}
			}
		}
	}

	return res
}
