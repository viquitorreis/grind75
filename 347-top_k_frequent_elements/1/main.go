package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	res := []int{}
	if len(nums) == 0 {
		return res
	}

	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	mxh := NewMaxHeap()
	heap.Init(mxh)

	for num, freq := range freq {
		heap.Push(mxh, Element{Val: num, Freq: freq})
	}

	for range k {
		res = append(res, heap.Pop(mxh).(Element).Val)
	}

	return res
}

type Element struct {
	Val  int
	Freq int
}

type maxHeap []Element

func NewMaxHeap() *maxHeap {
	return &maxHeap{}
}

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

func (m *maxHeap) Less(i, j int) bool {
	return (*m)[i].Freq > (*m)[j].Freq
}

func (m *maxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeap) Len() int {
	return len(*m)
}
