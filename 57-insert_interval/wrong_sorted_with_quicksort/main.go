package main

import (
	"fmt"
)

func main() {
	intervals := [][]int{
		{
			1, 2,
		},
		{
			2, 1,
		},
		{
			4, 7,
		},
	}
	new := []int{3, 4}
	newIntervals := insert(intervals, new)
	// printIntervals(newIntervals)
	fmt.Printf("new interval: %+v\n", newIntervals)
}

// intervals[i] = [starti, endi]
// intervals are sorted by start ASC
// insert in ascending order
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return insertRec(intervals, 0, len(intervals)-1)
}

func insertRec(intervals [][]int, min, max int) [][]int {
	if min < max {
		pivot := partition(intervals, min, max)
		insertRec(intervals, min, pivot-1)
		insertRec(intervals, pivot+1, max)
	}

	return intervals
}

// partition
//
//	rearrange the array until pivot is in correct order
//	move i, j until j is on the smallest element and swap with pivot, return j as new pivot index
//		i will be placed on the biggest element
func partition(intervals [][]int, min, max int) int {
	pivot := intervals[min]
	i, j := min+1, max

	for {
		// fmt.Printf("pivot: %d - i: %d - j: %d\n", pivot, i, j)
		if i >= j {
			break
		}

		for i <= max && intervals[i][0] <= pivot[0] {
			i++
		}
		for j >= min && intervals[j][0] > pivot[0] {
			j--
		}

		if i < j {
			swap(i, j, intervals)
		}
	}

	intervals[min] = intervals[j]
	intervals[j] = pivot

	return j
}

func swap(i, j int, intervals [][]int) {
	temp := intervals[i]
	intervals[i] = intervals[j]
	intervals[j] = temp
}

func printIntervals(intervals [][]int) {
	println("printing............")
	for i := range intervals {
		fmt.Printf("{%d:%d} ", intervals[i][0], intervals[i][1])
	}
	println()
}
