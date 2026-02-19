package main

import "sort"

func main() {

}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := []job{}
	for i, profit := range profit {
		jobs = append(jobs, job{
			Profit:    profit,
			StartTime: startTime[i],
			EndTime:   endTime[i],
		})
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].EndTime < jobs[j].EndTime
	})

	dp := make([]int, len(jobs))
	dp[0] = jobs[0].Profit

	for i := 1; i < len(jobs); i++ {
		lastCompatible := findLastCompatibleJob(jobs, i)

		curr := jobs[i].Profit
		if lastCompatible != -1 {
			curr += dp[lastCompatible]
		}

		dp[i] = max(curr, dp[i-1])
	}

	return dp[len(dp)-1]
}

type job struct {
	Profit    int
	StartTime int
	EndTime   int
}

func findLastCompatibleJob(jobs []job, i int) int {
	left, right := 0, i-1

	candidate := -1
	for left <= right {
		mid := (left + right) >> 1
		if jobs[mid].EndTime <= jobs[i].StartTime {
			candidate = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return candidate
}
