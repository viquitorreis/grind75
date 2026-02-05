package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	// precisamos de 4 ponteiros
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	res := []int{}

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		// precisamos checar se top é menor que bottom, se Não nao tem como percorrer
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}

	return nil
}
