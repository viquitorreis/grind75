package main

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}

	stack := []*TreeNode{}
	curr := root
	count := 0

	for curr != nil || count < k {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// processamos, todos estão ordenados
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++

		// se chegamos onde queremos, retornamos
		if count == k {
			return curr.Val
		}

		curr = curr.Right
	}

	return stack[k].Val
}
