package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{
		{1, 100},
		{11, 22},
		{1, 11},
		{2, 12},
	}))
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	toRemove := 0
	lastEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < lastEnd {
			toRemove++
		} else {
			lastEnd = intervals[i][1]
		}
	}

	return toRemove
}
