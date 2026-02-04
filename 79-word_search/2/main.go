package main

func main() {

}

func exist(board [][]byte, word string) bool {
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	var backtrack func(k, i, j int) bool
	backtrack = func(k, i, j int) bool {
		if k == len(word) {
			return true
		}

		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[k] {
			return false
		}

		temp := board[i][j]
		board[i][j] = '#'

		for _, dir := range directions {
			if backtrack(k+1, i+dir[0], j+dir[1]) {
				return true
			}
		}

		board[i][j] = temp
		return false
	}

	for row := range board {
		for col := range board[0] {
			if backtrack(0, row, col) {
				return true
			}
		}
	}

	return false
}
