package main

import (
	"container/heap"
	"fmt"
)

func main() {
	c := Constructor()
	c.AddNum(1)
	c.AddNum(2)
	fmt.Println(c.FindMedian())
	c.AddNum(3)

	fmt.Println(c.FindMedian())
}

type MedianFinder struct {
	lo *maxHeap
	hi *minHeap
}

func Constructor() MedianFinder {
	mxH := &maxHeap{}
	minH := &minHeap{}
	heap.Init(mxH)
	heap.Init(minH)

	return MedianFinder{
		lo: mxH,
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
	// se tamanho é impar, lo tem 1 a mais que hi
	if this.lo.Len() > this.hi.Len() {
		return float64((*this.lo)[0])
	}
	// fmt.Println("/")
	return float64((*this.hi)[0]+(*this.lo)[0]) / 2.0
}

type maxHeap []int
type minHeap []int

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() (x any) {
	x, *m = (*m)[len(*m)-1], (*m)[:len(*m)-1]
	return x
}

func (m *maxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *maxHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() (x any) {
	x, *m = (*m)[len(*m)-1], (*m)[:len(*m)-1]
	return x
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeap) Len() int {
	return len(*m)
}
