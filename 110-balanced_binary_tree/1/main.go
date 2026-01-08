package main

func main() {
	// tree := &TreeNode{Val: 3}
	// tree1 := &TreeNode{Val: 9}
	// tree2 := &TreeNode{Val: 20}
	// tree3 := &TreeNode{Val: 15}
	// tree4 := &TreeNode{Val: 7}

	// tree.Left = tree1
	// tree.Right = tree2
	// tree2.Left = tree3
	// tree2.Right = tree4

	tree := &TreeNode{Val: 3}
	tree1 := &TreeNode{Val: 9}
	tree2 := &TreeNode{Val: 20}
	tree3 := &TreeNode{Val: 15}
	tree4 := &TreeNode{Val: 7}
	// tree5 := &TreeNode{Val: 4}
	// tree6 := &TreeNode{Val: 4}

	tree.Left = tree1
	tree.Right = tree2
	tree2.Left = tree3
	tree2.Right = tree4
	// tree3.Left = tree5
	// tree4.Right = tree6

	println(isBalanced(tree))
}

// Complexity Analsisys

// Time Complexity - Big O
// 		Big O(n²) -> para cada nó, visitamos os de baixo

// Space Complexity
//		O(n) -> RECURSION STACK -> recursão, chamadas na stack

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// para uma árvore estar balanceada, todas as subárvores esquerda e direita não podem ter uma diferença de mais de 1 element
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := getHeight(root.Left)
	r := getHeight(root.Right)

	return l-r < 2 && r-l < 2 && isBalanced(root.Left) && isBalanced(root.Right) // recursão -> garantir que fazemos em todas subárvores
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}
