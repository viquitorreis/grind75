package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println()
}

func topKFrequent(nums []int, k int) []int {
	res := []int{}
	if len(nums) == 0 {
		return res
	}

	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	mh := &maxHeap{}
	heap.Init(mh)
	for num, f := range freq {
		heap.Push(mh, Element{
			Val:  num,
			Freq: f,
		})
	}

	for k > 0 {
		res = append(res, heap.Pop(mh).(Element).Val)
		k--
	}

	return res
}

type Element struct {
	Val  int
	Freq int
}

type maxHeap []Element

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(Element))
}

func (m *maxHeap) Pop() any {
	old := *m
	size := len(old)
	x := old[size-1]
	*m = old[:size-1]
	return x
}

func (m *maxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeap) Less(i, j int) bool {
	return (*m)[i].Freq > (*m)[j].Freq
}

func (m *maxHeap) Len() int {
	return len(*m)
}
