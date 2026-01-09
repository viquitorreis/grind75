package main

import (
	"fmt"
)

func main() {
	tree1 := &TreeNode{Val: 1}
	tree2 := &TreeNode{Val: 2}
	tree3 := &TreeNode{Val: 3}

	tree1.Left = tree2
	tree1.Right = tree3

	printTree(invertTree(tree1))
	println()
	printTree(invertTree(tree1))

	t1 := &TreeNode{Val: 4}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 7}
	t4 := &TreeNode{Val: 1}
	t5 := &TreeNode{Val: 3}
	t6 := &TreeNode{Val: 6}
	t7 := &TreeNode{Val: 9}

	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5
	t3.Left = t6
	t3.Right = t7

	println()

	printTree(t1)
	println()
	printTree(invertTree(t1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	head := root

	if root != nil {
		invertRec(root)
	}

	return head
}

func invertRec(root *TreeNode) {
	if root != nil {
		tmp := root.Right
		root.Right = root.Left
		root.Left = tmp
		invertRec(root.Right)
		invertRec(root.Left)
	}
}

func printTree(root *TreeNode) {
	if root != nil {
		fmt.Printf("%d ", root.Val)
		printTree(root.Left)
		printTree(root.Right)
	}
}
