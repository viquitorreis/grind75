package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidRec(root, math.MinInt, math.MaxInt)
}

func isValidRec(curr *TreeNode, min, max int) bool {
	if curr == nil {
		return true
	}

	if curr.Val <= min || curr.Val >= max {
		return false
	}

	return isValidRec(curr.Left, min, curr.Val) && isValidRec(curr.Right, curr.Val, max)
}
