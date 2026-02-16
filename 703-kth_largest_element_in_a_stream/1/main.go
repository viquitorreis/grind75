package main

import (
	"container/heap"
	"fmt"
)

func main() {
	kth := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kth.Add(3))
}

type KthLargest struct {
	hi *minHeap
	k  int
}

func Constructor(k int, nums []int) KthLargest {
	mh := &minHeap{}
	heap.Init(mh)

	for _, num := range nums {
		heap.Push(mh, num)
	}

	return KthLargest{
		hi: mh,
		k:  k,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.hi, val)
	for this.hi.Len() > this.k {
		heap.Pop(this.hi)
	}
	return (*this.hi)[0]
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
	return (*m)[i] < (*m)[j]
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Len() int {
	return len(*m)
}
