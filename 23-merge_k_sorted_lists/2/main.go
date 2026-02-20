package main

import (
	"container/heap"
	"fmt"
)

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}

	fmt.Println(mergeKLists([]*ListNode{l1, l2, l3}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	vals := make([]*ListNode, len(lists))
	for _, list := range lists {
		// fmt.Println("puxando list:", list)
		vals = append(vals, list)
	}

	mh := &minHeap{}
	heap.Init(mh)

	for _, l := range vals {
		// fmt.Println("l", l)
		// fmt.Println("puishing: ", l.Val)
		if l != nil {
			heap.Push(mh, l)
		}
	}

	// fmt.Println("mh", mh)

	list := &ListNode{}
	curr := list

	// fmt.Println("heap len", mh.Len())

	for mh.Len() > 0 {
		least := heap.Pop(mh).(*ListNode)
		// fmt.Println("least is:", least.Val)
		curr.Next = least
		curr = least
		if least.Next != nil {
			heap.Push(mh, least.Next)
		}
	}

	return list.Next
}

type minHeap []*ListNode

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(*ListNode))
}

func (m *minHeap) Pop() (x any) {
	x, *m = (*m)[len(*m)-1], (*m)[:len(*m)-1]
	return x
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i].Val < (*m)[j].Val
}

func (m *minHeap) Len() int {
	return len(*m)
}
