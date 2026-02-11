package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	t := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 3}
	t4 := &TreeNode{Val: 4}
	t5 := &TreeNode{Val: 5}

	t.Left = t2
	t.Right = t3
	t3.Left = t4
	t3.Right = t5

	c := Constructor()

	serialized := c.serialize(t)

	fmt.Println(serialized)

	tree := c.deserialize(serialized)

	printTree(tree)
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var s strings.Builder

	q := []*TreeNode{root}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr == nil {
			s.WriteString("N,")
			continue
		}

		s.WriteString(strconv.Itoa(curr.Val))
		s.WriteString(",")

		q = append(q, curr.Left)
		q = append(q, curr.Right)
	}

	return s.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	// se o ultimo caractere for , removemos. Depois separamos a string em um array de strings
	data = strings.TrimSuffix(data, ",")
	treeStrings := strings.Split(data, ",")

	// se o primeiro for nulo raiz é nula...
	if treeStrings[0] == "N" {
		return nil
	}

	rootVal, _ := strconv.Atoi(treeStrings[0])
	root := &TreeNode{Val: rootVal}

	// precisamos da fila para processar em todos os níveis corretamente e não só da raiz
	// dessa forma vamos saber exatamente qual nó pai os filhos vao estar ligados

	q := []*TreeNode{root}

	i := 1
	for len(q) > 0 && i < len(treeStrings) {
		curr := q[0]
		q = q[1:]

		if i < len(data) && !strings.EqualFold(treeStrings[i], "N") {
			leftVal, _ := strconv.Atoi(treeStrings[i])
			curr.Left = &TreeNode{Val: leftVal}
			q = append(q, curr.Left)
		}
		i++

		if i < len(data) && !strings.EqualFold(treeStrings[i], "N") {
			rightVal, _ := strconv.Atoi(treeStrings[i])
			curr.Right = &TreeNode{Val: rightVal}
			q = append(q, curr.Right)
		}
		i++
	}

	return root
}

func printTree(t *TreeNode) {
	if t == nil {
		return
	}

	q := []*TreeNode{t}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		fmt.Printf("%d ", curr.Val)

		if curr.Left != nil {
			q = append(q, curr.Left)
		}

		if curr.Right != nil {
			q = append(q, curr.Right)
		}
	}
}
