package main

import (
	"fmt"
	"math"
)

func main() {
	t := &TreeNode{Val: 1}
	t1 := &TreeNode{Val: 2}
	t2 := &TreeNode{Val: 3}
	t3 := &TreeNode{Val: 4}
	t4 := &TreeNode{Val: 5}

	t.Left = t1
	t.Right = t2
	t2.Right = t3
	t3.Right = t4

	fmt.Println(isBalanced(t))
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

	return isBalancedRec(root) && isBalancedRec(root.Left) && isBalancedRec(root.Right)
}

func isBalancedRec(curr *TreeNode) bool {
	if curr == nil {
		return true
	}

	left := getHeight(curr.Left)
	right := getHeight(curr.Right)

	// left e right atual devem ter diferença menor que 2
	// depois disso chama para os dois filhos left e right recusivamente, para conferir todas subarvores
	return math.Abs(float64(left)-float64(right)) < 2 &&
		isBalancedRec(curr.Left) &&
		isBalancedRec(curr.Right)
}

func getHeight(curr *TreeNode) int {
	if curr == nil {
		return 0
	}

	left := getHeight(curr.Left)
	right := getHeight(curr.Right)

	return max(left, right) + 1
}
