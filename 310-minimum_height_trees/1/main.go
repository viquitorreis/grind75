package main

import "fmt"

func main() {
	// t := &TreeNode{Val: 1}
	fmt.Println(findMinHeightTrees(3, [][]int{
		{1, 0},
		{1, 2},
		{1, 3},
	}))
}

// n -> nodes
// edges[i] = [ai, bi] -> undirected graph -> ai and bi is undirected edge
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	// graph -> connections
	graph := make(map[int][]int)

	for _, edge := range edges {
		a := edge[0]
		b := edge[1]

		// grafo bidirecionado (não direcionado)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	// precisamos encontrar os nós folhas com grau 1
	leaves := []int{}
	degree := make([]int, n)
	for node := 0; node < n; node++ {
		degree[node] = len(graph[node])
		if degree[node] == 1 {
			leaves = append(leaves, node)
		}
	}

	// precisamos remover camada por camada até sobrar no maximo 2 (que podem ser centro)
	remaining := n // quantos nós ainda estão no grafo
	for remaining > 2 {
		numLeaves := len(leaves)
		remaining -= numLeaves

		newLeaves := []int{}

		for _, leaf := range leaves {
			for _, neighbor := range graph[leaf] {
				degree[neighbor]--

				// se vizinho vira folha adiciona como leave atual
				if degree[neighbor] == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}

		// atualiza leaves para proxima iteração
		leaves = newLeaves
	}

	// vai sobrar 1 ou 2
	return leaves
}
