package main

func main() {

}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	i, j := 0, 0

	for i < len(firstList) && j < len(secondList) {
		start := max(firstList[i][0], secondList[j][0])
		end := min(firstList[i][1], secondList[j][1])

		// se start <= end, existe interseção
		if start <= end {
			res = append(res, []int{start, end})
		}

		// avanca o ponteiro do intervalo que termina primeiro
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}

	return res
}
