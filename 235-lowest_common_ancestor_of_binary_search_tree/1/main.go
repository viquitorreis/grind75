package main

import "fmt"

func main() {
	tree2 := &TreeNode{Val: 5}
	tree1 := &TreeNode{Val: 3}
	tree := &TreeNode{
		Val:   4,
		Left:  tree1,
		Right: tree2,
	}

	fmt.Println(lowestCommonAncestor(tree, tree1, tree2))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// lca != p && lca != q
	lca := root

	for lca != nil {
		// 1.p e q > que current node
		if p.Val > lca.Val && q.Val > lca.Val {
			lca = lca.Right
			// 2. p e q < que current node
		} else if p.Val < lca.Val && q.Val < lca.Val {
			lca = lca.Left
			// 3. Caso o split da árvore aconteça, ou um dos valores é o lca, nesse caso encontramos o valor desejado
		} else {
			return lca
		}
	}

	return lca
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
