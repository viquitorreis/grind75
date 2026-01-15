package main

import (
	"fmt"
	"math"
)

func main() {
	tree1 := &TreeNode{Val: 3}
	tree2 := &TreeNode{Val: 9}
	tree3 := &TreeNode{Val: 20}
	tree4 := &TreeNode{Val: 15}
	tree5 := &TreeNode{Val: 7}

	tree1.Left = tree2
	tree1.Right = tree3
	tree3.Left = tree4
	tree3.Right = tree5

	fmt.Println(isBalanced(tree1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) (res int) {
		if root == nil {
			return 0
		}

		return max(dfs(root.Left), dfs(root.Right)) + 1
	}

	l := dfs(root.Left)
	r := dfs(root.Right)

	return math.Abs(float64(l-r)) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}
