package main

import "fmt"

func main() {
	board := [][]byte{
		{
			'A', 'B', 'C', 'E',
		},
		{
			'S', 'F', 'C', 'S',
		},
		{
			'A', 'D', 'E', 'E',
		},
	}
	fmt.Println(exist(board, "ABCCED"))
	// fmt.Println(exist(board, "ABCB"))
}

func exist(board [][]byte, word string) bool {
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var backtrackDFS func(startIdx, row, col int) bool
	backtrackDFS = func(startIdx, row, col int) bool {
		if startIdx == len(word) {
			return true
		}

		// out of bound or doesnt exist on word / already visited (string changed on graph)
		if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || board[row][col] != word[startIdx] {
			return false
		}

		temp := board[row][col]
		board[row][col] = '#' // block from revisiting vertex

		for _, dir := range directions {
			if backtrackDFS(startIdx+1, row+dir[0], col+dir[1]) {
				return true
			}
		}

		// backtrack
		board[row][col] = temp
		return false
	}

	for i := range len(board) {
		for j := range len(board[0]) {
			if backtrackDFS(0, i, j) {
				return true
			}
		}
	}

	return false
}
