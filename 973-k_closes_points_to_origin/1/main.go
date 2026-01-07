package main

import (
	"container/heap"
	"sort"
)

func main() {

}

// https://github.com/leetcode-golang-classroom/golang_k_closest_point_to_origin

// esse problema podemos implementar uma max heap e retornar os k elementos da ponta da árvore
func kClosest(points [][]int, k int) [][]int {
	if k >= len(points) {
		return points
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	distances := []int{}
	for i := 0; i < len(points); i++ {

	}

	sort.Ints(distances)

	return [][]int{}
}

type Coordinate struct {
	x, y int
}

type MaxHeap []Coordinate

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Coordinate))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := (*h)[n-1] // ultimo elemento
	*h = (*h)[:n-1]
	return x
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].EuclideanDist() > h[j].EuclideanDist()
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Coordinate) EuclideanDist() int {
	return ((h).x * (h).x) + ((h).y * (h).y)
}
