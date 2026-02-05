package main

import (
	"fmt"
	"math"
)

func main() {
	t := &TreeNode{Val: 2}
	t1 := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 3}

	t.Left = t1
	t.Right = t2

	fmt.Println(isValidBST(t))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return validRec(root, math.MinInt, math.MaxInt)
}

func validRec(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val < min || node.Val > max {
		return false
	}

	return validRec(node.Left, min, node.Val) && validRec(node.Right, node.Val, max)
}
