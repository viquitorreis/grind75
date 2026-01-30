package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorder traversal -> Visita o node primeiro, depois a subtree a esquerda, depois a direita
// inorder traversal -> Para cada node visitamos a árvore da esquerda primeiro, depois o node, depois a da direita
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]

	rootIdx := findRootIdx(rootVal, inorder)

	root := &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:rootIdx+1], inorder[:rootIdx]),
		Right: buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:]),
	}

	return root
}

func findRootIdx(root int, inorder []int) int {
	for i, num := range inorder {
		if num == root {
			return i
		}
	}

	return -1
}
