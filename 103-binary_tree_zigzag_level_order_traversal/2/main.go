package main

func main() {

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

		res = append(res, level)
		leftToRight = !leftToRight
	}

	return res
}

func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
