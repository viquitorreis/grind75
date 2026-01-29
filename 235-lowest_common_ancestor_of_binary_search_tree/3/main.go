package main

import "fmt"

func main() {
	t := &TreeNode{Val: 6}
	t1 := &TreeNode{Val: 2}
	t2 := &TreeNode{Val: 0}
	t3 := &TreeNode{Val: 4}

	t.Left = t1
	t1.Left = t2
	t1.Right = t3

	fmt.Println(lowestCommonAncestor(t, t2, t3))
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

	curr := root

	for curr != nil {
		if p.Val < curr.Val && q.Val < curr.Val {
			curr = curr.Left
		} else if p.Val > curr.Val && q.Val > curr.Val {
			curr = curr.Right
		} else {
			// (q.Val < curr.Val && p.Val > curr.Val) || curr == q || curr == p {
			return curr
		}

	}

	return nil
}
