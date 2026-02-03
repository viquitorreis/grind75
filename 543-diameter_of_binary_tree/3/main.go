package main

import "fmt"

func main() {
	t1 := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 3}
	t4 := &TreeNode{Val: 4}

	t1.Left = t2
	t1.Right = t3
	t3.Right = t4

	fmt.Println(diameterOfBinaryTree(t1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var (
		dfs      func(node *TreeNode) int
		diamater int
	)

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		// diametro é sempre left + right
		diamater = max(diamater, left+right)

		fmt.Println(left, right)

		return max(left, right) + 1
	}

	dfs(root)

	return diamater
}
