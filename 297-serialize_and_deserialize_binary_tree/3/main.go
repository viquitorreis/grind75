package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	t := &TreeNode{Val: 1}
	t1 := &TreeNode{Val: 2}
	t2 := &TreeNode{Val: 3}
	t3 := &TreeNode{Val: 4}
	t4 := &TreeNode{Val: 5}

	t.Left = t1
	t.Right = t2
	t2.Left = t3
	t2.Right = t4

	c := Constructor()

	fmt.Println(c.serialize(t))
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

	var res strings.Builder
	q := []*TreeNode{root}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr == nil {
			// res.WriteString("N")
			res.WriteString("N,")
			continue
		} else {
			num := strconv.Itoa(curr.Val)
			res.WriteString(num)
		}
		res.WriteString(",")

		q = append(q, curr.Left)
		q = append(q, curr.Right)
	}

	return strings.TrimSuffix(res.String(), ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if strings.EqualFold(data, "") {
		return nil
	}

	words := strings.Split(data, ",")

	i := 0
	rootNum, _ := strconv.Atoi(words[i])
	q := []*TreeNode{{Val: rootNum}}
	root := q[0]
	i++

	for len(q) > 0 && i < len(data) {
		curr := q[0]
		q = q[1:]

		if i < len(data) && !strings.EqualFold(words[i], "N") {
			num, _ := strconv.Atoi(words[i])
			leftNode := &TreeNode{Val: num}
			curr.Left = leftNode
			q = append(q, leftNode)
		}
		i++

		if i < len(data) && !strings.EqualFold(words[i], "N") {
			num, _ := strconv.Atoi(words[i])
			rightNode := &TreeNode{Val: num}
			curr.Right = rightNode
			q = append(q, rightNode)
		}
		i++
	}

	return root
}
