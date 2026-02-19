package main

import (
	"container/heap"
	"fmt"
)

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	fmt.Println(leastInterval(tasks, 2))
}

func leastInterval(tasks []byte, n int) int {
	freq := make(map[byte]int)
	for _, t := range tasks {
		freq[t]++
	}

	ts := []Task{}
	for key, freq := range freq {
		ts = append(ts, Task{
			Val:  key,
			Freq: freq,
		})
	}

	hp := cpuHeap(ts)
	heap.Init(&hp)

	fmt.Println("heap", hp)

	cycles := 0
	for hp.Len() > 0 {
		levelsTasks := []Task{}

		for i := 0; i < n+1; i++ {
			if hp.Len() > 0 {
				task := heap.Pop(&hp).(Task)
				task.Freq--
				if task.Freq > 0 {
					levelsTasks = append(levelsTasks, task)
				}
				cycles++
			} else if len(levelsTasks) == 0 {
				break // nada pra fazer
			} else {
				cycles++ // idle slot
			}
		}

		if len(levelsTasks) > 0 {
			for _, task := range levelsTasks {
				heap.Push(&hp, task)
			}
		}
	}

	return cycles
}

type Task struct {
	Val  byte
	Freq int
}

type cpuHeap []Task

func (c *cpuHeap) Push(x any) {
	*c = append(*c, x.(Task))
}

func (c *cpuHeap) Pop() (m any) {
	*c, m = (*c)[:len(*c)-1], (*c)[len(*c)-1]
	return m
}

func (c *cpuHeap) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *cpuHeap) Less(i, j int) bool {
	return (*c)[i].Freq > (*c)[j].Freq
}

func (c *cpuHeap) Len() int {
	return len(*c)
}
