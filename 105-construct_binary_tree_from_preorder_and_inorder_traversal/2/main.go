package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// root = 3, preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	if len(preorder) == 0 {
		return nil
	}

	root := preorder[0]
	rootIdx := getRootIdx(root, inorder)

	node := &TreeNode{
		Val:   root,
		Left:  buildTree(preorder[1:rootIdx+1], inorder[:rootIdx]),
		Right: buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:]),
	}

	return node
}

func getRootIdx(root int, inorder []int) int {
	for i, num := range inorder {
		if num == root {
			return i
		}
	}

	return -1
}
