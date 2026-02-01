package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	res := [][]int{}

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

	resp := [][]int{}

	for i := len(res) - 1; i >= 0; i-- {
		resp = append(resp, res[i])
	}

	return resp
}
