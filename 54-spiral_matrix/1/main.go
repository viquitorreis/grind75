package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	// precisamos de 4 indices, left, right, top, bottom
	lines, cols := len(matrix)-1, len(matrix[0])-1
	// bounds
	top, bottom := 0, lines
	left, right := 0, cols

	res := []int{}

	for top <= bottom && left <= right {
		// primeiro percorremos da esquerda para a direita
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++ // ja percorremos essa linha

		// baixo
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right-- // ja percorremos essa coluna

		// percorremos apenas se tem como ir para a esquerda...
		// se top é menor que bottom, ainda tem espaço
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom-- // ja percorremos essa linha de baixo
		}

		// percorrer a de cima, só podemos quando left <= right, que é quando
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++ // ja percorremos essa coluna
		}
	}

	return res
}
