package main

func main() {

}

// COMPLEXITY ANALYSIS

// Time Complexity - Big O
//		O(nxm) - visita cada pixel apenas uma vez

// Space Complexity
//		O(nxm) - RECURSION STACK - recursão acumula chamadas na stack

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	// se [i][j] == image[sr][sc] -> grafos vizinhos
	originalGraphColor := image[sr][sc]

	if originalGraphColor == color {
		return image
	}

	dfs(image, sr, sc, originalGraphColor, color)

	return image
}

func dfs(image [][]int, row, col int, originalColor, newColor int) {
	// checando limites ou se a cor já está certa
	if row < 0 || row >= len(image) ||
		col < 0 || col >= len(image[0]) ||
		image[row][col] != originalColor {
		return
	}

	image[row][col] = newColor

	// vizinhos de cima, baixo, direita, esquerda
	dfs(image, row-1, col, originalColor, newColor) // cima
	dfs(image, row+1, col, originalColor, newColor) // baixo
	dfs(image, row, col+1, originalColor, newColor) // atras
	dfs(image, row, col-1, originalColor, newColor) // frente
}
