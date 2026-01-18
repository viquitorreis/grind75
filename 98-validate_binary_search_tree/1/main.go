package main

import (
	"fmt"
	"math"
)

func main() {
	tree1 := &TreeNode{Val: 2}
	tree2 := &TreeNode{Val: 2}
	tree3 := &TreeNode{Val: 2}

	tree1.Left = tree2
	tree1.Right = tree3

	fmt.Println(isValidBST(tree1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validateRec(root, math.MinInt, math.MaxInt)
}

func validateRec(curr *TreeNode, min, max int) bool {
	if curr == nil {
		return true
	}

	if curr.Val <= min || curr.Val >= max {
		return false
	}

	return validateRec(curr.Left, min, curr.Val) && validateRec(curr.Right, curr.Val, max)
}
