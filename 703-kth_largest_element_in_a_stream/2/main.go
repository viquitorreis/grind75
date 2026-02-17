package main

import (
	"container/heap"
	"fmt"
)

func main() {
	c := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(c.Add(3))
}

type maxHeap []int

type KthLargest struct {
	nums *maxHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	mh := maxHeap(nums)
	heap.Init(&mh)

	return KthLargest{
		nums: &mh,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.nums, val)
	for this.nums.Len() > this.k {
		heap.Pop(this.nums)
	}
	return (*this.nums)[0]
}

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(int))
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
	return (*m)[i] < (*m)[j]
}

func (m *maxHeap) Len() int {
	return len(*m)
}
