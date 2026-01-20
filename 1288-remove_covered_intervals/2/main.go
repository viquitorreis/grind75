package main

import "sort"

func main() {

}

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	count := 0 // intervalo não cobertos
	maxEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] > maxEnd {
			count++
			maxEnd = intervals[i][1]
		}
		// se interval[1] <= maxEnd, está coberto, não precisa incrementar
	}

	return count
}
