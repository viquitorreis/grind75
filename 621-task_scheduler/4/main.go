package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}

func leastInterval(tasks []byte, n int) int {
	res := 0
	if len(tasks) == 0 {
		return res
	}

	freq := make(map[byte]int)
	for _, task := range tasks {
		freq[task]++
	}

	mh := &cpuHeap{}
	heap.Init(mh)
	for b, freq := range freq {
		heap.Push(mh, Task{
			Val:  b,
			Freq: freq,
		})
	}

	for mh.Len() > 0 {
		cycles := n + 1
		updatedTasks := []Task{}

		for range cycles {
			if mh.Len() > 0 {
				curr := heap.Pop(mh).(Task)
				curr.Freq--

				if curr.Freq > 0 {
					updatedTasks = append(updatedTasks, curr)
				}

				// quando executar tarefa sempre conta
				res++
			} else if len(updatedTasks) > 0 {
				// slot vazio, MAS ainda tem trabalho para voltar pra heap
				res++
			}
		}

		for _, task := range updatedTasks {
			heap.Push(mh, task)
		}
	}

	return res
}

type Task struct {
	Val  byte
	Freq int
}

type cpuHeap []Task

func (m *cpuHeap) Push(x any) {
	*m = append(*m, x.(Task))
}

func (m *cpuHeap) Pop() (x any) {
	x, *m = (*m)[len(*m)-1], (*m)[:len(*m)-1]
	return x
}

func (m *cpuHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *cpuHeap) Less(i, j int) bool {
	return (*m)[i].Freq > (*m)[j].Freq
}

func (m *cpuHeap) Len() int {
	return len(*m)
}
