package main

import "fmt"

func main() {
	tree1 := &TreeNode{Val: 6}
	tree2 := &TreeNode{Val: 2}
	tree3 := &TreeNode{Val: 8}
	tree4 := &TreeNode{Val: 0}
	tree5 := &TreeNode{Val: 4}
	tree6 := &TreeNode{Val: 7}
	tree7 := &TreeNode{Val: 9}
	tree8 := &TreeNode{Val: 3}
	tree9 := &TreeNode{Val: 5}

	tree1.Left = tree2
	tree1.Right = tree3
	tree2.Left = tree4
	tree2.Right = tree5
	tree5.Left = tree8
	tree5.Right = tree9
	tree3.Left = tree6
	tree3.Right = tree7

	fmt.Println(lowestCommonAncestor(tree1, tree2, tree4).Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	lca := root

	for lca != nil {
		if p.Val < lca.Val && q.Val < lca.Val {
			lca = lca.Left
		} else if p.Val > lca.Val && q.Val > lca.Val {
			lca = lca.Right
		} else {
			// caso de split
			return lca
		}

	}

	return lca
}
