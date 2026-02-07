package main

func main() {

}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	// edges[i] = [ai, bi]
	// significa que tem uma aresta não direcionada entre os dois nodes...
	graph := make(map[int][]int)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	// precisamos processar as folhas primeiro
	// as folhas, são os nodes que jamais serao o centro (minimo) pois elas estao muito longe dos outros
	leaves := []int{}
	degree := make([]int, n)
	for node := 0; node < n; node++ {
		degree[node] = len(graph[node])
		// precisamos pegar as folhas, que tem degree == 1
		if degree[node] == 1 {
			leaves = append(leaves, node)
		}
	}

	// processamos até sobrar os elementos centrais (< 2)
	remaining := n
	for remaining > 2 {
		numLeaves := len(leaves)
		// vamos processar todas as folhas nesse nivel então subtraimos na quantidade de grafos para processar
		remaining -= numLeaves
		newLeaves := []int{}

		for _, leaf := range leaves {
			// processamos todos os vizinhos dessa folha
			for _, neighbor := range graph[leaf] {
				degree[neighbor]--

				if degree[neighbor] == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}

		leaves = newLeaves
	}

	return leaves
}
