package main

import "fmt"

func main() {
	tree1 := &TreeNode{Val: 3}
	tree2 := &TreeNode{Val: 9}
	// tree3 := &TreeNode{Val: 20}
	// tree4 := &TreeNode{Val: 15}
	// tree5 := &TreeNode{Val: 7}

	tree1.Left = tree2
	// tree1.Right = tree3
	// tree3.Left = tree4
	// tree3.Right = tree5

	fmt.Println(maxDepth(tree1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) (depth int) {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left) + 1
		right := dfs(root.Right) + 1

		return max(left, right)
	}

	return dfs(root)
}
