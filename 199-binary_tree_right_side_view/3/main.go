package main

func main() {

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

	q := []*TreeNode{root}
	res := []int{}

	for len(q) > 0 {
		levelSize := len(q)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			curr := q[0]
			q = q[1:]

			if i == levelSize-1 {
				level = append(level, curr.Val)
			}

			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}

		res = append(res, level...)
	}

	return res
}
