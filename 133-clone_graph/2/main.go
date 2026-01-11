package main

import "fmt"

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

	fmt.Println(cloneGraph(node1))
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	seen := make(map[*Node]*Node) // vamos retornar o novo grafo do map
	q := []*Node{node}            // fila com nós antigos
	seen[node] = &Node{Val: node.Val, Neighbors: []*Node{}}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, neighbor := range curr.Neighbors {
			// 1. caso de nao existir no map -> cria e poe na fila
			_, ok := seen[neighbor]
			if !ok {
				seen[neighbor] = &Node{Val: neighbor.Val, Neighbors: []*Node{}}
				q = append(q, neighbor)
			}

			// 2. sempre adiciona como vizinho nesse node, existindo ou não.
			seen[curr].Neighbors = append(seen[curr].Neighbors, seen[neighbor])
		}
	}

	return seen[node]
}
