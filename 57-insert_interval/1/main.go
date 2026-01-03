package main

import (
	"fmt"
	"log"
)

func main() {
	intervals := [][]int{
		{
			1, 2,
		},
		{
			2, 3,
		},
		{
			4, 7,
		},
	}

	new := []int{3, 8}

	newIntervals := insert(intervals, new)

	fmt.Printf("new interval: %+v\n", newIntervals)
}

// 1. checar se cada intervalo é menor que o novo intervalo. Adicionar ao resultado
// 2. se qualquer um tiver overlapping, merge do intervalo e adiciona o resultado
// 3. adiciona o resto dos intervalos que são maiores que o newInterval ao resultado

func insert(intervals [][]int, newInterval []int) [][]int {
	for i := range intervals {
		curr := intervals[i]
		log.Println("curr: ", curr, newInterval)
		// is overlapping ?
		if newInterval[1] == curr[0] {
			newInterval[1] = curr[1]
			log.Printf("newInterval[1] == curr[0] - newInterval: %+v", newInterval)
		}

		if newInterval[0] == curr[0] {
			// interval = max()
			log.Println("min is equal")
			newInterval[1] = max(curr[1], newInterval[1])
			log.Println("newInterval now is: ", newInterval)
		}

		if newInterval[0] >= curr[0] && newInterval[0] < curr[1] {
			fmt.Printf("its overlapping curr[0] - %+v - %+v\n", newInterval, curr)
			newInterval[0] = min(curr[0], newInterval[0])
		}

		if newInterval[1] >= curr[0] && newInterval[1] <= curr[1] {
			fmt.Printf("its overlapping curr[0] - %+v - %+v\n", newInterval, curr)
			newInterval[1] = max(newInterval[1], curr[1])
		}

		intervals[i] = newInterval
	}

	return intervals
}
