package main

import "fmt"

func main() {
	t := &TreeNode{Val: 3}
	t1 := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 4}
	t3 := &TreeNode{Val: 2}

	t.Left = t1
	t.Right = t2
	t2.Right = t3

	fmt.Println(kthSmallest(t, 1))
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	// precisamos fazer inorder traversal
	if root == nil {
		return -1
	}

	curr := root
	count := 0
	stack := []*TreeNode{}

	for curr != nil || len(stack) < k {
		// primeiro vamos no máximo a esquerda possível
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		printStack(stack)

		// processamos os elementos
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// somamos apenas depois de processar o elemento
		count++

		if count == k {
			return curr.Val
		}

		curr = curr.Right
	}

	return -1
}

func printStack(stack []*TreeNode) {
	fmt.Print("stack now ")
	for _, el := range stack {
		fmt.Printf("%d ", el.Val)
	}
	fmt.Printf("\n")
}
