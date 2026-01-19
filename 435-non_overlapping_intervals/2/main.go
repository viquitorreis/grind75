package main

import "sort"

func main() {

}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
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
