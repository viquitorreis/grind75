package main

import "fmt"

func main() {
	fmt.Println()
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		levelSize := len(q)
		level := []int{}

		for range levelSize {
			curr := q[0]
			q = q[1:]

			level = append(level, curr.Val)

			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}

		res = append(res, level)
	}

	ans := [][]int{}

	for i := len(res) - 1; i >= 0; i-- {
		ans = append(ans, res[i])
	}

	return ans
}
