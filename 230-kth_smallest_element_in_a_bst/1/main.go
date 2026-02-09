package main

import "fmt"

func main() {
	t := &TreeNode{Val: 5}
	t1 := &TreeNode{Val: 3}
	t2 := &TreeNode{Val: 6}
	t3 := &TreeNode{Val: 2}
	t4 := &TreeNode{Val: 4}
	t5 := &TreeNode{Val: 1}

	t.Left = t1
	t.Right = t2
	t1.Right = t4
	t1.Left = t3
	t3.Left = t5

	// t := &TreeNode{Val: 3}
	// t1 := &TreeNode{Val: 1}
	// t2 := &TreeNode{Val: 4}
	// t3 := &TreeNode{Val: 2}

	// t.Left = t1
	// t.Right = t2
	// t1.Right = t3

	fmt.Println(kthSmallest(t, 1))
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}

	curr := root
	stack := []*TreeNode{}
	count := 0

	// inorder traversal: tudo para esquerda -> raiz -> tudo para direita
	// Precisamos do inorder para buscar de forma crescente...
	for curr != nil || len(stack) > 0 {
		// 1. Primeiro vai tudo para a esquerda
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// 2. processa os nodes e soma count, se chegar em k retorna
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++

		if count == k {
			return curr.Val
		}

		// 3. Vai para a direita
		curr = curr.Right
	}

	return -1
}
