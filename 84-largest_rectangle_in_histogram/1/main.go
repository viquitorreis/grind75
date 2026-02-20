package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

func largestRectangleArea(heights []int) int {
	var square int
	if len(heights) == 0 {
		return square
	}

	heights = append(heights, 0)

	// stack armazena os índices do menor pro maior
	for i, stack, top := 0, make([]int, len(heights)), -1; i < len(heights); i++ {
		for top != -1 && heights[i] <= heights[stack[top]] {
			limitHeight := heights[stack[top]]
			// como salvamos a altura do topo, vamos descartar o índice do topo
			// nao apagamos o valor, apenas movendo o ponteiro para baixo
			// dessa forma no proximo passo vamos pegar o novo topo da stack
			top--

			width := 0
			// não tem trasação mais barata aberta ainda. Altura maxima é a do elemento atual
			if top == -1 {
				width = i
			} else {
				width = i - 1 - stack[top]
			}

			square = max(square, limitHeight*width)
		}
		// depois do loop interno, a barra atual heights[i] é MENOR OU IGUAL A NENHUM TOPO RESTANTE
		// ou seja, é a menor barra vista até agora que ainda não foi fechada
		// então abrimos uma nova transação com ela, empilhamos o indice i
		top++
		stack[top] = i
	}

	return square
}
