package main

import "fmt"

func main() {
	t := &TreeNode{Val: 1}
	t1 := &TreeNode{Val: 2}
	t2 := &TreeNode{Val: 3}
	t3 := &TreeNode{Val: 4}
	t4 := &TreeNode{Val: 5}

	t.Left = t1
	t.Right = t2
	t1.Right = t4
	t2.Right = t3

	fmt.Println(rightSideView(t))
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		levelSize := len(q)

		for i := range levelSize {
			curr := q[0]
			q = q[1:]

			if i == levelSize-1 {
				res = append(res, curr.Val)
			}

			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
	}

	return res
}
