package main

import "fmt"

func main() {
	tree1 := &TreeNode{Val: 3}
	tree2 := &TreeNode{Val: 9}
	tree3 := &TreeNode{Val: 20}
	tree4 := &TreeNode{Val: 15}
	tree5 := &TreeNode{Val: 7}

	tree1.Left = tree2
	tree1.Right = tree3
	tree3.Left = tree4
	tree4.Right = tree5

	fmt.Println(maxDepth(tree1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var dfs func(r *TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}

		leftHeight := dfs(r.Left) + 1
		rightHeight := dfs(r.Right) + 1

		return max(leftHeight, rightHeight)
	}

	return dfs(root)
}
