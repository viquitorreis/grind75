package main

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 2}
	node4 := &ListNode{Val: 1}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	println("is palindrome: ", isPalindrome(node1))

	node5 := &ListNode{Val: 1}
	node6 := &ListNode{Val: 2}

	node5.Next = node6
	// node6.Next = node3

	println("is palindrome: ", isPalindrome(node5))
}

// return true if its a palindrome
// false otherwise
type ListNode struct {
	Val  int
	Next *ListNode
}

// Complexity Analysis
// Big O:
//		O(n) - n being amount of nodes on the list - assymptotically
//			Since we are traversing the entire list exactly once.
// 			The second loop will do in maximum n/2, and the rest only simple comparisons and variables.

// Space Complexity:
//		O(n) also
//			Since we are storing all the nodes on a slice exactly once

func isPalindrome(head *ListNode) bool {
	var vals []byte
	for head != nil {
		vals = append(vals, byte(head.Val))
		head = head.Next
	}

	for i, j := 0, len(vals)-1; i <= j; i, j = i+1, j-1 {
		if vals[i] != vals[j] {
			return false
		}
	}

	return true
}
