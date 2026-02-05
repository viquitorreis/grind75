package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		last := res[len(res)-1]

		if curr[0] <= last[1] {
			last[1] = max(curr[1], last[1])
		} else {
			res = append(res, curr)
		}
	}

	return res
}
