package main

import "fmt"

func main() {
	// tree := &TreeNode{
	// 	Val: 3,
	// 	Left: &TreeNode{
	// 		Val:   9,
	// 		Right: nil,
	// 		Left:  nil,
	// 	},
	// 	Right: &TreeNode{
	// 		Val:   20,
	// 		Left:  &TreeNode{Val: 15},
	// 		Right: &TreeNode{Val: 7},
	// 	},
	// }
	// tree := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val:   2,
	// 		Left:  &TreeNode{Val: 4},
	// 		Right: &TreeNode{Val: 5},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 3,
	// 	},
	// }
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

// retornar a ordem de travessia por nível - cada camada
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		lenQ := len(queue)
		level := []int{}

		for range lenQ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level)
	}

	return res
}
