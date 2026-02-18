package main

import (
	"sort"
)

func main() {
	jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70})
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := []job{}
	for i := range startTime {
		jobs = append(jobs, job{
			profit:    profit[i],
			startTime: startTime[i],
			endTime:   endTime[i],
		})
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].endTime < jobs[j].endTime
	})

	dp := make([]int, len(profit))
	dp[0] = jobs[0].profit

	for i := 1; i < len(jobs); i++ {
		// achamos o indice do ultimo job compatível com o i
		j := findLastCompatible(jobs, i)

		take := jobs[i].profit
		if j != -1 {
			// pegamos o maximo de lucro que o ultimo job compatível teve, o array dp vai nos dizer isso
			take += dp[j]
		}

		// quando decidimos pular o job i, estamo pensando "não vou fazer esse job
		// vou manter exatamente o resultado que tinha antes de olhar para ele, que é o indice i-1 no dp
		skip := dp[i-1]

		// max entre:
		// 		- soma do profit do job atual + profit maximo do ultimo job compativel
		// 		- soma do profit no indice anterior do dp - sem contar esse job
		dp[i] = max(take, skip)
	}

	return dp[len(dp)-1]
}

type job struct {
	profit    int
	startTime int
	endTime   int
}

func findLastCompatible(jobs []job, i int) int {
	left, right := 0, i-1
	candidate := -1

	for left <= right {
		mid := (left + right) >> 1

		if jobs[mid].endTime <= jobs[i].startTime {
			candidate = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return candidate
}
