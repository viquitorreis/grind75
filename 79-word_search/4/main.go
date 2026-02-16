package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED"))
}

func exist(board [][]byte, word string) bool {
	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var backtrack func(idx, row, col int) bool
	backtrack = func(idx, row, col int) bool {
		if idx == len(word) {
			return true
		}

		if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || board[row][col] != word[idx] {
			return false
		}

		// marcamos a celula atual
		temp := board[row][col]
		board[row][col] = '#'

		//DFS
		// nos 4 vizinhos
		for _, dir := range dirs {
			if backtrack(idx+1, row+dir[0], col+dir[1]) {
				// só retorna true se chegou no final da palavra bem sucedido (base case)
				return true
			}
		}

		// se todos os 4 retornarem false na condição de cima
		// aqui vai retornar false
		board[row][col] = temp
		return false
	}

	for row := range len(board) {
		for col := range len(board[0]) {
			if backtrack(0, row, col) {
				return true
			}
		}
	}

	return false
}
