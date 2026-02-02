package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	q := []*TreeNode{root}
	avg := []float64{}

	for len(q) > 0 {
		levelSize := len(q)
		sum := 0.0

		for range levelSize {
			curr := q[0]
			q = q[1:]

			sum += float64(curr.Val)

			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}

		avg = append(avg, sum/float64(levelSize))
	}

	return avg
}
