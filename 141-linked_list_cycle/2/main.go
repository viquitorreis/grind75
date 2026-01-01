package main

import "fmt"

func main() {
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}

	node1.Next = node2
	node2.Next = node3
	node3.Next = nil

	println("has cycle?", hasCycle(node1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	curr := head
	seen := make(map[string]int)

	for curr != nil {
		nodeAddr := fmt.Sprintf("%s", curr)
		if val, ok := seen[nodeAddr]; ok && val >= 2 {
			return true
		}

		seen[nodeAddr]++
		curr = curr.Next
	}

	return false
}
