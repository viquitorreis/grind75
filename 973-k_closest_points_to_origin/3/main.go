package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {

}

func kClosest(points [][]int, k int) [][]int {
	if len(points) == 0 {
		return nil
	}

	hp := &kHeap{}
	heap.Init(hp)

	for _, point := range points {
		heap.Push(hp, Coordinate{x: point[0], y: point[1]})
	}

	res := [][]int{}
	for range k {
		el := heap.Pop(hp).(Coordinate)
		fmt.Println("el:", el)
		res = append(res, []int{el.x, el.y})
	}

	return res
}

type Coordinate struct {
	x, y int
}

type kHeap []Coordinate

func (k *kHeap) Push(x any) {
	*k = append(*k, x.(Coordinate))
}

func (k *kHeap) Pop() any {
	old := *k
	size := len(old)
	x := old[size-1]
	*k = old[:size-1]
	return x
}

func (k *kHeap) EuclideanDist(i int) float64 {
	return math.Sqrt(float64((*k)[i].x*(*k)[i].x) + float64((*k)[i].y*(*k)[i].y))
}

// func (k *kHeap) Peek()

func (k *kHeap) Less(i, j int) bool {
	return k.EuclideanDist(i) < k.EuclideanDist(j)
}

func (k *kHeap) Len() int {
	return len(*k)
}

func (k *kHeap) Swap(i, j int) {
	(*k)[i], (*k)[j] = (*k)[j], (*k)[i]
}
