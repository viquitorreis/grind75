package main

import "sort"

func main() {

}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	arrows := 1
	lastEnd := points[0][1]

	for i := 1; i < len(points); i++ {
		// overlap, atualizamos a INTERSECÇÃO DOS DOIS
		// a arrow continua acertando os dois, mas precisamos garantir que é sobre a janela correta!
		if points[i][0] <= lastEnd {
			lastEnd = min(lastEnd, points[i][1])
		} else {
			// sem overlap na janela anterior, precisamos atualizar a arrow e o novo lastEnd
			arrows++
			lastEnd = points[i][1]
		}
	}

	return arrows
}
