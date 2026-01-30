package main

import "fmt"

func main() {
	t := &TreeNode{Val: 1}
	t1 := &TreeNode{Val: 2}
	t2 := &TreeNode{Val: 3}
	t3 := &TreeNode{Val: 5}
	t4 := &TreeNode{Val: 4}

	t.Left = t1
	t.Right = t2
	t1.Right = t3
	t2.Right = t4

	fmt.Println("rightSideView:", rightSideView(t))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// bfs -> sempre colocamos o ultimo elemento do nível na resposta
	q := []*TreeNode{root}
	res := []int{}

	for len(q) > 0 {
		levelSize := len(q)

		for i := 0; i < levelSize; i++ {
			curr := q[0]
			q = q[1:]

			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}

			if i == levelSize-1 {
				res = append(res, curr.Val)
			}

		}
	}

	return res
}
