package main

import "fmt"

func main() {
	// intervals := [][]int{
	// 	{1, 3},
	// 	{6, 9},
	// }
	// newInterval := []int{2, 5}
	// intervals := [][]int{
	// 	{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	// }
	// newInterval := []int{4, 8}
	intervals := [][]int{{1, 5}}
	newInterval := []int{6, 8}
	fmt.Println(insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	newIntervals := [][]int{}
	n := 0

	// 1. Achar todos intervalos menores sem overlap
	for n <= len(intervals)-1 && newInterval[0] > intervals[n][1] {
		newIntervals = append(newIntervals, intervals[n])
		n++
	}

	if len(intervals) < n && newInterval[1] < intervals[n][0] {
		newIntervals = append(newIntervals, newInterval)
		n++
	}

	// 2. Achar todos intervalos com overlap -> faz o merge
	for n <= len(intervals)-1 && newInterval[1] >= intervals[n][0] {
		// for i := n;
		newInterval[0] = min(intervals[n][0], newInterval[0])
		newInterval[1] = max(intervals[n][1], newInterval[1])
		n++
	}
	newIntervals = append(newIntervals, newInterval)

	// 3. Append do resto
	newIntervals = append(newIntervals, intervals[n:]...)

	return newIntervals
}
