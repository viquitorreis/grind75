package main

import (
	"container/heap"
	"sort"
)

func main() {

}

type Interval struct {
	start int
	end   int
}

type priotityQueue []Interval

func (p *priotityQueue) Push(x any) {
	*p = append(*p, x.(Interval))
}

func (p *priotityQueue) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func (p *priotityQueue) Peek() int {
	return (*p)[0].end
}

func (p *priotityQueue) Len() int {
	return len(*p)
}

// PRECISAMOS ORDENAR POR END!!! (termino da reuniao)
// precisamos ter sempre no topo a reuniao que termina mais cedo, para saber se podemos reusar a sala...
func (p *priotityQueue) Less(i, j int) bool {
	return (*p)[i].end < (*p)[j].end
}

func (p *priotityQueue) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func minMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	if len(intervals) == 1 {
		return 1
	}

	// precisamos contar os que não se sobreboem
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	pq := &priotityQueue{}

	heap.Init(pq)

	for _, interval := range intervals {
		// comparamos o end do topo com o start da proxima reuniao
		if pq.Len() > 0 && (*pq)[0].end <= interval.start {
			heap.Pop(pq) // sala liberou, pode reusar
		}
		heap.Push(pq, interval)
	}

	return pq.Len()
}
