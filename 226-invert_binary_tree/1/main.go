package main

import "fmt"

func main() {
	tree2 := TreeNode{Val: 3}
	tree1 := TreeNode{Val: 1}
	tree := TreeNode{Val: 2, Left: &tree1, Right: &tree2}
	printTree(invertTree(&tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//

func invertTree(root *TreeNode) *TreeNode {
	invertRec(root)
	return root
}

func invertRec(root *TreeNode) {
	if root != nil {
		// 1. invertendo os filhos left -> right, right -> left
		temp := root.Left
		root.Left = root.Right
		root.Right = temp

		// fazendo o mesmo processo em todas subtrees, começando da subtree esquerda
		invertRec(root.Left)
		invertRec(root.Right)
	}
}

func printTree(tree *TreeNode) {
	if tree != nil {
		fmt.Printf("Val: %d\n", tree.Val)
		printTree(tree.Left)
		printTree(tree.Right)
	}
}
