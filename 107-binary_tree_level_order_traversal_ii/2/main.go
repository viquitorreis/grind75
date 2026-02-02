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

	left, right := 0, len(res)-1

	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}

	return res
}
