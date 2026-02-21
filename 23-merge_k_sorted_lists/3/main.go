package main

import (
	"container/heap"
	"fmt"
)

func main() {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}

	fmt.Println(mergeKLists([]*ListNode{l, l1, l2}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	mh := &minHeap{}
	heap.Init(mh)

	for _, l := range lists {
		if l != nil {
			heap.Push(mh, l)
		}
	}

	list := &ListNode{}
	curr := list

	for mh.Len() > 0 {
		node := heap.Pop(mh).(*ListNode)
		curr.Next = node
		curr = node
		if node.Next != nil {
			heap.Push(mh, node.Next)
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
