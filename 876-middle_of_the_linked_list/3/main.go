package main

import "fmt"

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	printList(middleList(node1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleList(head *ListNode) *ListNode {
	slow, fast := head, head
	// curr := head

	for fast != nil {
		slow = slow.Next
		fast = slow.Next.Next
	}

	return slow
}

func printList(l *ListNode) {
	if l != nil {
		fmt.Printf("curr value is: %d\n", l.Val)
		list := l.Next
		printList(list)
	}
}
