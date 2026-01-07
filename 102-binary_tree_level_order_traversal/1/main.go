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

// COMPLEXITY ANALYSIS
// BIG O:
//		O(n) - onde n é o número de nós. Visitamos um nó só uma vez

// Space Complexity
//		O(n) - no pior caso
//			onde a fila pode conter até n/2 - ultima camada de uma árvore completa
//			o array de resultado também armazena todos os n valores

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
