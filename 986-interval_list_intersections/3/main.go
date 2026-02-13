package main

import "fmt"

func main() {
	first := [][]int{
		{0, 2}, {5, 10}, {13, 23}, {24, 25},
	}
	second := [][]int{
		{1, 5}, {8, 12}, {15, 24}, {25, 26},
	}
	fmt.Println(intervalIntersection(first, second))
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	i, j := 0, 0

	for i < len(firstList) && j < len(secondList) {
		a1, a2 := firstList[i][0], firstList[i][1]
		b1, b2 := secondList[j][0], secondList[j][1]

		if a1 <= b2 && b1 <= a2 {
			res = append(res, []int{max(a1, b1), min(a2, b2)})
		}

		// quem termina primeiro deve avançár...
		if a2 < b2 {
			i++
		} else {
			j++
		}
	}

	return res
}
