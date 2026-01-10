package main

import (
	"fmt"
)

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = append(node1.Neighbors, node2)
	node1.Neighbors = append(node1.Neighbors, node4)

	node2.Neighbors = append(node2.Neighbors, node1)
	node2.Neighbors = append(node2.Neighbors, node3)

	node3.Neighbors = append(node3.Neighbors, node2)
	node3.Neighbors = append(node3.Neighbors, node1)

	node4.Neighbors = append(node4.Neighbors, node1)
	node4.Neighbors = append(node4.Neighbors, node2)

	// graph := node1
	fmt.Printf("old: %x\n", node1)

	fmt.Printf("new: %x\n", cloneGraph(node1))
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	q := make([]*Node, 0)
	q = append(q, node)

	// precisamos garantir que vamos apontar o nó original para o nó novo, sempre
	// quando já existir registro usamos, se não criamos um nó novo
	// original -> novo
	clones := make(map[*Node]*Node)
	clones[node] = &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, neighbor := range curr.Neighbors {
			if _, ok := clones[neighbor]; !ok {
				clones[neighbor] = &Node{
					Val:       neighbor.Val,
					Neighbors: []*Node{},
				}

				q = append(q, neighbor)
			}

			clones[curr].Neighbors = append(clones[curr].Neighbors, clones[neighbor])

		}

	}

	return clones[node]
}
