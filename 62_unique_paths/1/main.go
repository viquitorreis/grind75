package main

func main() {

}

func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for y := range matrix {
		matrix[y] = make([]int, n)
	}

	for row := range m {
		matrix[row][0] = 1
	}

	for col := range n {
		matrix[0][col] = 1
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			matrix[row][col] = matrix[row-1][col] + matrix[row][col-1]
		}
	}

	return matrix[m-1][n-1]
}
