package main

func main() {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 223}}}}
	reorderList(l)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	lists := make([]*ListNode, 0)

	curr := head

	for curr != nil {
		lists = append(lists, curr)
		curr = curr.Next
	}

	left, right := 0, len(lists)-1

	for left < right {
		lists[left].Next = lists[right] // l0 -> Ln
		left++

		if left == right {
			break
		}

		lists[right].Next = lists[left] // ln -> l1
		right--
	}
	lists[left].Next = nil // fecha a lista
}
