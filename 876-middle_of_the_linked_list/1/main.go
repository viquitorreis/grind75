package main

import (
	"fmt"
)

func main() {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 4}
	node5 := ListNode{Val: 5}
	node6 := ListNode{Val: 6}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6

	printList(middleNode(&node1))
}

// return the middle node of the linked list
// if there is two middle nodes, return the second middle (when len(list) % 2 == 0)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(node *ListNode) *ListNode {
	seen := make(map[int]*ListNode, 0)
	len := 1
	curr := node
	head := node

	for curr != nil {
		seen[len] = curr
		len++
		curr = curr.Next
	}

	if len%2 != 0 {
		len++
	}

	var arr []*ListNode
	for i := len / 2; i < len; i++ {
		nodeSeen, ok := seen[i]
		if ok {
			arr = append(arr, nodeSeen)
		}
	}

	for head != arr[0] {
		head = head.Next
	}

	return head
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("node val: %d\n", node.Val)
		node = node.Next
	}
	println()
}
