package main

import "container/heap"

func main() {

}

type MedianFinder struct {
	lo *maxHeap
	hi *minHeap
}

func Constructor() MedianFinder {
	minH := &minHeap{}
	maxH := &maxHeap{}
	heap.Init(minH)
	heap.Init(maxH)

	return MedianFinder{
		lo: maxH,
		hi: minH,
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.hi, num)
	heap.Push(this.lo, heap.Pop(this.hi).(int))

	if this.lo.Len() > this.hi.Len()+1 {
		heap.Push(this.hi, heap.Pop(this.lo).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		return float64((*this.lo)[0])
	}
	return float64((*this.lo)[0]+(*this.hi)[0]) / 2.0
}

type minHeap []int
type maxHeap []int

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

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() any {
	old := *m
	size := len(*m)
	x := old[size-1]
	*m = old[:size-1]
	return x
}

func (m *maxHeap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *maxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeap) Len() int {
	return len(*m)
}
