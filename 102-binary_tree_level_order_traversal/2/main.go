package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: &TreeNode{Val: 5},
		},
	}
	fmt.Println(levelOrder(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	// podemos usar uma fila
	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// sempre pegamos o valor do node atual, e atualizamos a fila

		lenQ := len(queue)
		level := []int{} // elementos nesse nível

		// visitamos todos os nós nesse nível
		for range lenQ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, root.Left)
			}

			if node.Right != nil {
				queue = append(queue, root.Right)
			}
		}

		res = append(res, level)
	}

	return res
}
