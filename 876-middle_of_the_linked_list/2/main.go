package main

import "fmt"

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	printList(middleNode(node1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// slow -> dá um passo por vez
// fast -> dá dois passos por vez
// quando o fast estiver no final da lista, slow estará na metade, logicamente
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	println()
}
