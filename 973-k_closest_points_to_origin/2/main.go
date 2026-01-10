package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	// points := [][]int{
	// 	{1, 3},
	// 	{-2, 2},
	// }
	points := [][]int{
		{3, 3},
		{5, -1},
		{-2, 4},
	}
	// points := [][]int{
	// 	{1, 3},
	// 	{-2, 2},
	// 	{2, -2},
	// }
	fmt.Println(kClosest(points, 2))
}

// esse problema é uma min heap
// precisamos dos k elementos na melhor distância
// adicionamos à heap com base na distância euclidiana
// depois disso retornamos os k primeiros elementos da heap
func kClosest(points [][]int, k int) [][]int {
	if len(points) <= k {
		return points
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for _, point := range points {
		// heap.Push(maxHeap, points[i])
		x := point[0]
		y := point[1]

		newPoint := Coordinate{
			x: x,
			y: y,
		}

		if maxHeap.Len() < k {
			heap.Push(maxHeap, newPoint)
		} else {
			if maxHeap.Len() > 0 {
				// top := heap.Pop(maxHeap).(Coordinate)
				top := maxHeap.Peek()

				if k < maxHeap.Len() && top.EuclideanDist() > newPoint.EuclideanDist() {
					heap.Push(maxHeap, newPoint)
				}
			}
		}
	}

	res := [][]int{}
	for range k {
		el := heap.Pop(maxHeap).(Coordinate)
		res = append(res, []int{el.x, el.y})
	}

	return res
}

type Coordinate struct {
	x, y int
}

type MaxHeap []Coordinate

// heap
func (m *MaxHeap) Pop() any {
	size := len(*m)
	x := (*m)[size-1]
	// size := len(*m) - 1
	// newHeap := []Coordinate{}
	// newHeap = append(newHeap, (*m)[1:]...)
	*m = (*m)[:size-1]
	return x
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(Coordinate))
}

// sorting
func (m *MaxHeap) Less(i, j int) bool {
	return (*m)[i].EuclideanDist() > (*m)[j].EuclideanDist()
}

func (m *MaxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MaxHeap) Len() int {
	return len(*m)
}

func (m *MaxHeap) Peek() *Coordinate {
	return &(*m)[len(*m)-1]
}

func (c *Coordinate) EuclideanDist() int {
	// sum := float64(c.x + c.y)
	x := float64(c.x * c.x)
	y := float64(c.y * c.y)
	sum := x + y
	sqrt := int(math.Sqrt(sum))
	return sqrt
}
