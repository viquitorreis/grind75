package main

import "container/heap"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	pq := &priorityQ{}
	heap.Init(pq)

	// adiciona o primeiro nó de cada lista. O heap vai sempre manter o menor primeiro
	for _, list := range lists {
		if list != nil {
			heap.Push(pq, list)
		}
	}

	list := &ListNode{}
	curr := list

	for pq.Len() > 0 {
		smallest := heap.Pop(pq).(*ListNode)
		curr.Next = smallest
		curr = curr.Next

		if smallest.Next != nil {
			heap.Push(pq, smallest.Next)
		}
	}

	return list.Next
}

type priorityQ []*ListNode

func (p *priorityQ) Push(x any) {
	*p = append(*p, x.(*ListNode))
}

func (p *priorityQ) Pop() (x any) {
	x, *p = (*p)[len(*p)-1], (*p)[:len(*p)-1]
	return x
}

func (p *priorityQ) Peek(i int) *ListNode {
	return (*p)[i]
}

func (p *priorityQ) Less(i, j int) bool {
	return (*p)[i].Val < (*p)[j].Val
}

func (p *priorityQ) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *priorityQ) Len() int {
	return len(*p)
}
