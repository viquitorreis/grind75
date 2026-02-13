package main

import (
	"container/heap"
	"fmt"
)

func main() {
	elements := []int{-1, -2, -3, -4, -5}
	finder := Constructor()
	for _, num := range elements {
		finder.AddNum(num)
		fmt.Println("middle: ", finder.FindMedian())
	}
	fmt.Println(finder.FindMedian())
}

type MedianFinder struct {
	lo *maxHeap
	hi *minHeap
}

type maxHeap []int
type minHeap []int

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

	// precisamos garantir a ordem, topo esquerdo <= tiopo direito
	// caso constrário enviamos o elemento para a min heap (hi)
	if this.hi.Len() > 0 && this.lo.Peek() > this.hi.Peek() {
		heap.Push(this.hi, heap.Pop(this.lo))
	}

	// se hi fica maior que lo, precisamos mover de volta
	if this.hi.Len() > this.lo.Len() {
		heap.Push(this.lo, heap.Pop(this.hi))
	}

	// precisamos garantir o balanço nas heaps, lo pode ter no máximo 1 elemento a mais.
	if this.lo.Len() > this.hi.Len()+1 {
		heap.Push(this.hi, heap.Pop(this.lo))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		return float64(this.lo.Peek())
	}
	return float64(this.lo.Peek()+this.hi.Peek()) / 2.0
}

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeap) Pop() any {
	old := *m
	size := len(old)
	x := old[size-1]
	*m = old[:size-1]
	return x
}

func (m *maxHeap) Peek() int {
	return (*m)[0]
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

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Pop() any {
	old := *m
	size := len(old)
	x := old[size-1]
	*m = old[:size-1]
	return x
}

func (m *minHeap) Peek() int {
	return (*m)[0]
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeap) Len() int {
	return len(*m)
}
