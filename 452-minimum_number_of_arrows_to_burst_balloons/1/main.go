package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMinArrowShots([][]int{
		{10, 16},
		{2, 8},
		{1, 6},
		{7, 12},
	}))
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	fmt.Println(points)

	arrows := 1
	lastEnd := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] <= lastEnd {
			// overlap: precisamos atualizar a intersecção (menor end)
			lastEnd = min(lastEnd, points[i][1])
		} else {
			arrows++
			lastEnd = points[i][1]
		}
	}

	return arrows
}
