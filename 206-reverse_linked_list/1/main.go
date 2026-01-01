package main

import (
	"fmt"
)

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node1.Next = node2

	printList(reverseList(node1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return reverseRec(head, nil)
}

// pattern: In-place Reversal of Linked List
// this pattern reverses the direction of the pointers on a linked list without creating any new nodes or using extra space
// Needs to manage the following key points:
// 		a. where i currently am in the list
//		b. where i came from
// 		c. where im going next

// BIG O
// 		Space Complexity: O(n) - linear complexity
// 			we visit each node on the linked list exactly once, and only constant time operations on each node (saving the next pointer, variable...)
//	 	Space Complexity: O(n) - linear complexity
//			the recursive calls the stack, since im doing one recursive call per node before any of them return, i build up n stack frames on the call stack
//			each stack frame stores the local variables for that function call, so im using **space proportinal to the length of the list**
// 			BETTER SOLUTION FOR SPACE: implementing for loop and travess nodes... no building up all those stack frames...

func reverseRec(curr, prev *ListNode) *ListNode {
	if curr == nil {
		return prev
	}

	// 1. armazenamos o proximo node antes de quebrar o link
	next := curr.Next

	// 2. apontamos a lista antiga para o novo valor
	// é aqui que revertemos a arrow (reverse list)
	// pois apontamos o proximo da lista atual para o anterior
	curr.Next = prev

	// 3. atualizamos o curr
	curr.Next = prev

	// 4. current vira o novo prev
	// next vira o novo prev
	return reverseRec(next, curr)
}

func printList(list *ListNode) {
	if list != nil {
		fmt.Printf("node val: %d ", list.Val)
		list = list.Next
		printList(list)
	}

	println()
}
