package main

func main() {
	res := [][]int{
		{
			3, 1,
		},
		{
			8, 1,
		},
	}
	diff := []int{2, 2, 3, 1, 4, 5, 1, 1, 2}
	println(findMaxVal(10, res, diff))
}

func findMaxVal(n int, restrictions [][]int, diff []int) int {
	// Passo 1: Setup restrictions
	restrictionMap := make(map[int]int)
	for _, r := range restrictions {
		restrictionMap[r[0]] = r[1]
	}

	// arrays para armazenar max alcançável em cada direction
	forwardMax := make([]int, n)
	backwardMax := make([]int, n)

	// inicializar
	forwardMax[0] = 0 // a[0] = 0

	// para backward, inicialmente todos podem ser infinito
	for i := range n {
		backwardMax[i] = 1000000 // valor bem grande
	}

	for i := 1; i < n; i++ {
		// maximo que podemos ter aqui vindo da position anterior
		forwardMax[i] = forwardMax[i-1] + diff[i-1]

		// aplicar restriction se existir
		if maxAllowed, exists := restrictionMap[i]; exists {
			if forwardMax[i] > maxAllowed {
				forwardMax[i] = maxAllowed
			}
		}
	}

	// BACKWARD PASS: propagar restrictions de volta
	for i := n - 1; i >= 0; i-- {
		// aplicar restriction local se existir
		if maxAllowed, exists := restrictionMap[i]; exists {
			if backwardMax[i] > maxAllowed {
				backwardMax[i] = maxAllowed
			}
		}

		// propagar para position anterior
		if i > 0 {
			// se em position i o máximo é X, em i-1 o máximo é X + diff[i-1]
			maxFromNext := backwardMax[i] + diff[i-1]
			if maxFromNext < backwardMax[i-1] {
				backwardMax[i-1] = maxFromNext
			}
		}
	}

	// combinar e encontrar o máximo global
	globalMax := 0
	for i := range n {
		// o valor real alcançável é o mínimo entre forward e backward
		actualMax := forwardMax[i]
		if backwardMax[i] < actualMax {
			actualMax = backwardMax[i]
		}

		if actualMax > globalMax {
			globalMax = actualMax
		}
	}

	return globalMax
}
