package main

import "fmt"

func main() {
	t := &TreeNode{Val: 3}
	t1 := &TreeNode{Val: 9}
	t2 := &TreeNode{Val: 20}
	t3 := &TreeNode{Val: 15}
	t4 := &TreeNode{Val: 7}

	t.Left = t1
	t.Right = t2
	t2.Left = t3
	t2.Right = t4

	fmt.Println(zigzagLevelOrder(t))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	res := [][]int{}
	leftToRight := true

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

		if !leftToRight {
			reverse(level)
		}

		leftToRight = !leftToRight

		res = append(res, level)
	}

	return res
}

func reverse(level []int) {
	left, right := 0, len(level)-1
	for left < right {
		level[left], level[right] = level[right], level[left]
		left++
		right--
	}
}
