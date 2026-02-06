package main

import (
	"container/heap"
)

func main() {

}

func leastInterval(tasks []byte, n int) int {
	hp := &TaskHeap{}
	heap.Init(hp)

	freq := make(map[byte]int)
	for _, task := range tasks {
		freq[task]++
	}

	for task, count := range freq {
		heap.Push(hp, Task{
			Task:      task,
			Frequency: count,
		})
	}

	minimum := 0
	for hp.Len() > 0 {
		tasksLeft := []Task{} // precisamos retornar as tasks pro heap, se ainda tiver frequencias positivas
		cyclesProcessed := 0  // quantas tasks foram executadas nessa rodada

		for i := 0; i <= n; i++ {
			if hp.Len() > 0 {
				task := heap.Pop(hp).(Task)
				task.Frequency--
				cyclesProcessed++

				if task.Frequency > 0 {
					tasksLeft = append(tasksLeft, task)
				}
			}
		}

		for _, task := range tasksLeft {
			heap.Push(hp, task)
		}

		// se ainda tem tasks no heap, precisamos de n + 1 intervalos (incluindo as idles)
		// se nao tem mais nada, só conta quantas executamos
		if hp.Len() > 0 {
			minimum += n + 1
		} else {
			minimum += cyclesProcessed
		}
	}

	return minimum
}

type TaskHeap []Task

type Task struct {
	Task      byte
	Frequency int
}

func (h *TaskHeap) Push(x any) {
	*h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() any {
	old := *h
	n := len(*h)
	x := (old)[n-1]
	*h = (*h)[:n-1]
	return x
}

func (h *TaskHeap) Len() int {
	return len(*h)
}

func (h *TaskHeap) Less(i, j int) bool {
	return (*h)[i].Frequency > (*h)[j].Frequency
}

func (h *TaskHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
