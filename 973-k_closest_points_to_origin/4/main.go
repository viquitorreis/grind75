package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(kClosest([][]int{
		{1, 3},
		{-2, 2},
	}, 1))
}

func kClosest(points [][]int, k int) [][]int {
	res := [][]int{}
	if len(points) == 0 {
		return res
	}

	mh := &maxHeap{}
	heap.Init(mh)

	for _, point := range points {
		heap.Push(mh, Coordinate{
			x: point[0],
			y: point[1],
		})
	}

	for k > 0 {
		coord := heap.Pop(mh).(Coordinate)
		res = append(res, []int{coord.x, coord.y})
		k--
	}

	return res
}

type Coordinate struct {
	x, y int
}

type maxHeap []Coordinate

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(Coordinate))
	// fmt.Println("after push:", *m)
}

func (m *maxHeap) Print() {
	for _, num := range *m {
		fmt.Printf("%d:%d ", num.x, num.y)
	}
	fmt.Printf("\n")
}

func (c *Coordinate) EuclideanDist() float64 {
	return math.Sqrt(math.Pow(float64(c.x), 2) + math.Pow(float64(c.y), 2))
}

func (m *maxHeap) Less(i, j int) bool {
	return (*m)[i].EuclideanDist() < (*m)[j].EuclideanDist()
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

func (m *maxHeap) Len() int {
	return len(*m)
}
