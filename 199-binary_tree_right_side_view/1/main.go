package main

func main() {
	t := &TreeNode{Val: 1}
	t1 := &TreeNode{Val: 3}
	t2 := &TreeNode{Val: 5}

	t3 := &TreeNode{Val: 2}
	t4 := &TreeNode{Val: 6}
	t5 := &TreeNode{Val: 7}

	t.Right = t1
	t.Left = t3
	t3.Left = t4
	t4.Right = t5
	t1.Left = t2

	rightSideView(t)
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
		levelSize := len(q) // precisamos do tamanho do NÍVEL ATUAL para iterarar nele

		for i := range levelSize {
			curr := q[0]
			q = q[1:]

			// ultimo elemento colcoamos na resposta
			if i == levelSize-1 {
				res = append(res, curr.Val)
			}

			// processamos primeiro os elementos da esquyerda, por precisamos deixar os da direita por ultimo no nível da fila
			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}

	}

	return res
}
