package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validateRec(root, math.MinInt, math.MaxInt)
}

func validateRec(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return validateRec(root.Left, min, root.Val) && validateRec(root.Right, root.Val, max)
}
