package main

import "fmt"

func main() {
	t := &TreeNode{Val: 3}
	t1 := &TreeNode{Val: 9}
	t2 := &TreeNode{Val: 20}
	t3 := &TreeNode{Val: 15}
	t4 := &TreeNode{Val: 7}

	t.Left = t1
	t.Right = t2
	t2.Left = t3
	t2.Right = t4

	fmt.Println(maxDepth(t))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(depthRec(root.Left), depthRec(root.Right)) + 1
}

func depthRec(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := depthRec(node.Left)
	right := depthRec(node.Right)

	return max(left, right) + 1
}
