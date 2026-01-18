package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(removeCoveredIntervals([][]int{
		{1, 2},
		{1, 4},
		{3, 4},
	}))
}

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1] // quando start é igual o maior end vem primeiro
		}
		return intervals[i][0] < intervals[j][0]
	})

	count := 0 // intervalos cobertos (removidos)
	maxEnd := 0

	for _, interval := range intervals {
		// max end tá coberto pelo intervalo
		if interval[1] <= maxEnd {
			count++
		} else {
			maxEnd = interval[1]
		}
	}

	return len(intervals) - count
}
