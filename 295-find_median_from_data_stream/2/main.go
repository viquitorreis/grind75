package main

import (
	"container/heap"
)

func main() {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	mf.AddNum(3)
}

type MedianFinder struct {
	lo *maxHeap
	hi *minHeap
}

func Constructor() MedianFinder {
	lo := &maxHeap{}
	hi := &minHeap{}
	heap.Init(lo)
	heap.Init(hi)

	return MedianFinder{
		lo: lo,
		hi: hi,
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.lo, num)

	if this.hi.Len() > 0 && this.hi.Peek() < this.lo.Peek() {
		heap.Push(this.hi, heap.Pop(this.lo))
	}

	if this.hi.Len() > this.lo.Len() {
		heap.Push(this.lo, heap.Pop(this.hi))
	}

	if this.lo.Len() > this.hi.Len()+1 {
		heap.Push(this.hi, heap.Pop(this.lo))
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if (this.lo.Len()+this.hi.Len())%2 == 0 && this.hi.Len() > 0 {
		return float64(this.lo.Peek()+this.hi.Peek()) / 2.0
	}
	return float64(this.lo.Peek())
}

type maxHeap []int
type minHeap []int

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

func (m *maxHeap) Len() int {
	return len(*m)
}

func (m *maxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *maxHeap) Peek() int {
	return (*m)[0]
}

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

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeap) Peek() int {
	return (*m)[0]
}
