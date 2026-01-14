package main

import "fmt"

func main() {
	// tree1 := &TreeNode{Val: 1}
	// tree2 := &TreeNode{Val: 2}
	// tree3 := &TreeNode{Val: 3}
	// tree4 := &TreeNode{Val: 4}
	// tree5 := &TreeNode{Val: 5}

	// tree1.Left = tree2
	// tree2.Right = tree3
	// tree2.Left = tree4
	// tree2.Right = tree5
	// tree1.Right = tree3

	tree1 := &TreeNode{Val: 1}
	tree2 := &TreeNode{Val: 2}

	// tree1.Left = tree2
	// tree1 := &TreeNode{Val: 3}
	// tree2 := &TreeNode{Val: 1}
	// tree3 := &TreeNode{Val: 2}

	tree1.Left = tree2
	// tree2.Right = tree3

	fmt.Println(diameterOfBinaryTree(tree1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) (diamater int) {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftHeight := dfs(root.Left)
		rightheight := dfs(root.Right)

		diamater = max(diamater, leftHeight+rightheight)

		return max(leftHeight, rightheight) + 1
	}

	dfs(root)
	return diamater
}

func diameter(root *TreeNode) (diamater int) {

	return
}
