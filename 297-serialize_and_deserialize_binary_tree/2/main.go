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

	t.Left = t2
	t.Right = t3

	c := Constructor()
	fmt.Println(c.serialize(t))
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type Codec struct{}

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

		fmt.Println("curr: ", curr)

		if curr == nil {
			s.WriteString("N,")
			continue
		} else {
			s.WriteString(strconv.Itoa(curr.Val))
			s.WriteString(",")
		}

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

	data = strings.TrimSuffix(data, ",")
	nodesStr := strings.Split(data, ",")

	if nodesStr[0] == "N" {
		return nil
	}

	rootVal, _ := strconv.Atoi(nodesStr[0])
	root := &TreeNode{Val: rootVal}

	q := []*TreeNode{root}

	i := 1
	for len(q) > 0 && i < len(data) {
		curr := q[0]
		q = q[1:]

		if i < len(data) && !strings.EqualFold(nodesStr[i], "N") {
			leftVal, _ := strconv.Atoi(nodesStr[i])
			leftNode := &TreeNode{Val: leftVal}
			curr.Left = leftNode
			q = append(q, leftNode)
		}
		i++

		if i < len(data) && !strings.EqualFold(nodesStr[i], "N") {
			rightVal, _ := strconv.Atoi(nodesStr[i])
			rightNode := &TreeNode{Val: rightVal}
			curr.Right = rightNode
			q = append(q, rightNode)
		}
		i++
	}

	return root
}
