package main

import "fmt"

func main() {
	t := &TreeNode{Val: 1}
	t1 := &TreeNode{Val: 2}
	// t2 := &TreeNode{Val: 6}
	// t3 := &TreeNode{Val: 2}
	// t4 := &TreeNode{Val: 4}
	// t5 := &TreeNode{Val: 1}

	// t.Left = t1
	t.Right = t1
	// t1.Left = t3
	// t1.Right = t4
	// t3.Left = t5

	fmt.Println(kthSmallest(t, 1))
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--

		if k == 0 {
			return curr.Val
		}

		curr = curr.Right
	}

	return -1
}
