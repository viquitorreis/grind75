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

	var backtrack func(k, row, col int) bool
	backtrack = func(k, row, col int) bool {
		if k == len(word) {
			return true
		}

		if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || board[row][col] != word[k] {
			return false
		}

		temp := board[row][col]
		board[row][col] = '#'

		for _, dir := range directions {
			// escolhe
			if backtrack(k+1, row+dir[0], col+dir[1]) {
				return true
			}
		}

		board[row][col] = temp
		return false
	}

	for i := range len(board) {
		for j := range len(board[0]) {
			if backtrack(0, i, j) {
				return true
			}
		}
	}

	return false
}
