package main

func main() {

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
		levelNums := []int{}

		for range levelSize {
			curr := q[0]
			q = q[1:]

			levelNums = append(levelNums, curr.Val)

			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}

		res = append(res, levelNums)
	}

	ans := make([][]int, 0)

	for i := len(res) - 1; i >= 0; i-- {
		ans = append(ans, res[i])
	}

	return ans
}
