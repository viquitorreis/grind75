package main

import (
	"sort"
)

func main() {

}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	arrows := 1
	lastEnd := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] <= lastEnd {
			lastEnd = min(lastEnd, points[i][1])
		} else {
			lastEnd = points[i][1]
			arrows++
		}
	}

	return arrows
}
