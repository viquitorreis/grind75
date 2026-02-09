package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}

func leastInterval(tasks []byte, n int) int {
	tasksMap := make(map[byte]Task)
	for _, b := range tasks {
		val, ok := tasksMap[b]
		if ok {
			val.Freq++
			tasksMap[b] = val
		} else {
			tasksMap[b] = Task{
				Val:  b,
				Freq: 1,
			}
		}
	}

	thp := &taskHeap{}
	heap.Init(thp)

	for _, v := range tasksMap {
		heap.Push(thp, v)
	}

	minCycles := 0
	for thp.Len() > 0 {
		tasks := []Task{}
		levelCycles := 0

		// precisamos processar no maximo n+1 tasks, e nao todas. POis é a janela de cooldown
		// for range levelSize {
		for i := 0; i <= n && thp.Len() > 0; i++ {
			el := heap.Pop(thp).(Task)
			el.Freq--
			levelCycles++

			if el.Freq > 0 {
				tasks = append(tasks, el)
			}
		}

		// manda de volta pra heap
		for _, task := range tasks {
			heap.Push(thp, task)
		}

		// quando tem mais elementos na heap, precisamos contar os idles
		// ou seja, se ainda tem tasks para processar, contamos a janela completa (com os intervalos idle)
		// se nao tem mais elemento na heap, contamos apenas o cycle computados
		if thp.Len() > 0 {
			minCycles += n + 1
		} else {
			minCycles += levelCycles
		}
	}

	return minCycles
}

type taskHeap []Task

type Task struct {
	Val  byte
	Freq uint16
}

func (t *taskHeap) Push(x any) {
	*t = append(*t, x.(Task))
}

func (t *taskHeap) Pop() any {
	old := *t
	size := len(old)
	x := old[size-1]
	*t = old[:size-1]
	return x
}

func (t *taskHeap) Len() int {
	return len(*t)
}

func (t *taskHeap) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

// max heap
func (t *taskHeap) Less(i, j int) bool {
	return (*t)[i].Freq > (*t)[j].Freq
}
