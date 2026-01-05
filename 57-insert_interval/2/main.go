package main

import "fmt"

func main() {
	intervals := [][]int{
		{
			1, 2,
		},
		{
			2, 3,
		},
		{
			4, 7,
		},
	}

	new := []int{3, 8}
	printIntervals(insert(intervals, new))
}

// Complexity analysis
// Big O
//		- O(n)
// 		Em cada loop, o PIOR CASO, será O(n)
//		Porém, os loops são independentes, se um executa no array todo, os outros não.
//

// Space Complexity
//		- O(n)
//		Estamos criando um novo array, de NO MAXIMO tamanho do outro array
//		O restante são variáveis simples O(1)

// 1. intervalos atuais que vem antes do novo (sem overlap) - adiciona ao resultado
// 2. intervalos com overlap - faz o merge deles
// 3. intervalos depois - sem overlap
func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i := 0
	n := len(intervals)

	// 1. adicionar todos intervalos que terminam antes do novo começar
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// 2. merge / fundir, todos intervalos que sobrepoe o novo

	//	ex: [a, b] e [c, d], overlap se:
	// 		1. a <= d -> o primeiro começa antes ou quando o segundo termina
	//		2. c <= b -> o segundo começa antes ou quando o primeiro termina
	//		Se as duas forem verdadeiras, tem overlap

	//		OBS: a fase 1, ja garante a SEGUNDA condição
	// 			quando saimos da fase 1, temos a garantia de que intervals[i][1] >= newInterval[0]
	//			pois remove todos os elementos intervals[i][1] < newInterval[0]

	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	// adiciona o novo intervalo completo - mergeado com os que sobrepoe
	result = append(result, newInterval)

	// 3. adiciona todos os outros intervalos restantes
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func printIntervals(intervals [][]int) {
	for i := range intervals {
		fmt.Printf("%d -  %d ", intervals[i][0], intervals[i][1])
	}
	println()
}
