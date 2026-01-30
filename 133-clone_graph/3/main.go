package main

func main() {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// precisamos armazenar os nodes já vistos anteriormente para adicionar o node corretamente
	// posteriormente quando for referenciado (vizinho de alguém)
	nodes := make(map[*Node]*Node)
	q := []*Node{node}
	nodes[node] = &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, neighbor := range curr.Neighbors {
			// se nao viu ainda, cria
			_, ok := nodes[neighbor]
			if !ok {
				nodes[neighbor] = &Node{
					Val:       neighbor.Val,
					Neighbors: make([]*Node, 0),
				}
				q = append(q, neighbor)
			}

			// adiciona aos vizinhos do curr
			nodes[curr].Neighbors = append(nodes[curr].Neighbors, nodes[neighbor])
		}
	}

	return nodes[node]
}
