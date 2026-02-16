package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func findKthLargest(nums []int, k int) int {
	mh := minHeap(nums)
	heap.Init(&mh)
	for mh.Len() > len(nums)-k+1 {
		heap.Pop(&mh)
		// fmt.Println("popped", popped)
	}
	return heap.Pop(&mh).(int)
}

type minHeap []int

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() any {
	old := *m
	size := len(old)
	x := old[size-1]
	*m = old[:size-1]
	return x
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}
