package main

func main() {

}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make(map[int][]int)
	degree := make([]int, n) // grau de cada vertice

	// 1. cria o grafo não direcionado -> tem que ir e voltar
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	for node := 0; node < n; node++ {
		degree[node] = len(graph[node])
	}

	leaves := []int{}
	for node, deg := range degree {
		if deg == 1 {
			leaves = append(leaves, node)
		}
	}

	remaining := n
	for remaining > 2 {
		remaining -= len(leaves)
		newLeafs := []int{}

		for _, leaf := range leaves {
			for _, neighbor := range graph[leaf] {
				degree[neighbor]--
				if degree[neighbor] == 1 {
					newLeafs = append(newLeafs, neighbor)
				}
			}
		}

		leaves = newLeafs
	}

	return leaves
}
