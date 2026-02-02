package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	rootIdx := rootIdx(rootVal, inorder)

	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:rootIdx+1], inorder[:rootIdx]),
		Right: buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:]),
	}
}

func rootIdx(root int, inorder []int) int {
	for i, num := range inorder {
		if num == root {
			return i
		}
	}

	return -1
}
