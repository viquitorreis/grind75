package main

import (
	"log"
	"math"
)

func main() {
	res := [][]int{
		{
			3, 1,
		},
		{
			8, 1,
		},
	}
	diff := []int{2, 2, 3, 1, 4, 5, 1, 1, 2}
	println(findMaxVal(10, res, diff))
}

func findMaxVal(n int, restrictions [][]int, diff []int) int {
	// seq := make([]int, n)
	seq := []int{}
	maxVal := 0
	for i := range n {
		log.Println("oi: ", maxVal, seq)
		seq = append(seq, diff[i])
		if diff[i] > maxVal {
			maxVal = diff[i]
		}

		if i == 0 {
			seq[i] = 0
			continue

		}

		if 0 <= i && i <= n-2 && (int(math.Abs(float64(seq[i]-seq[i+1]))) <= diff[i]) {
			if restrictions[i][0] == i && restrictions[i][1] == maxVal {
				if seq[i] >= maxVal {
					seq[i] = maxVal
				}
			}
		}
	}

	return maxVal
}
