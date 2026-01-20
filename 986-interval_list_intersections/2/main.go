package main

func main() {

}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ins := [][]int{}
	f, s := 0, 0

	for f < len(firstList) && s < len(secondList) {
		start := max(firstList[f][0], secondList[s][0])
		finish := min(firstList[f][1], secondList[s][1])

		if start <= finish {
			ins = append(ins, []int{start, finish})
		}

		// AVANÇA O QUE TERMINA PRIMEIRO!
		if firstList[f][1] < secondList[s][1] {
			f++
		} else {
			s++
		}
	}

	return ins
}
