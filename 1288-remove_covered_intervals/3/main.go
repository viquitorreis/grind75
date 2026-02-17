package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(removeCoveredIntervals([][]int{
		{1, 4},
		{3, 6},
		{2, 8},
	}))
}

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	maxEnd := intervals[0][1]
	toRemove := 0

	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]

		if curr[1] <= maxEnd {
			toRemove++
		} else {
			maxEnd = curr[1]
		}
	}

	return len(intervals) - toRemove
}
