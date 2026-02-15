package main

import (
	"container/heap"
	"fmt"
)

func main() {

}

func findKthLargest(nums []int, k int) int {
	hp := minHeap(nums)
	fmt.Println("hp: ", hp)
	heap.Init(&hp)
	for hp.Len() > k {
		heap.Pop(&hp)
	}
	return heap.Pop(&hp).(int)
}

type minHeap []int

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() any {
	old := *m
	size := len(*m)
	x := old[size-1]
	*m = old[:size-1]
	return x
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}
