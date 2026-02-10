package main

import (
	"fmt"
)

func main() {
	// fmt.Println(findMinHeightTrees(4, [][]int{
	// 	{1, 0},
	// 	{1, 2},
	// 	{1, 3},
	// }))

	fmt.Println(findMinHeightTrees(6, [][]int{
		{3, 0},
		{3, 1},
		{3, 2},
		{3, 4},
		{5, 4},
	}))

	// fmt.Println(findMinHeightTrees(3, [][]int{
	// 	{0, 1},
	// 	{0, 2},
	// }))
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make(map[int][]int)
	degrees := make([]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)

		degrees[a]++
		degrees[b]++
	}

	leaves := []int{}
	for node, d := range degrees {
		if d == 1 {
			leaves = append(leaves, node)
		}
	}

	remaining := n
	for remaining > 2 {
		numLeaves := len(leaves)
		remaining -= numLeaves
		newLeaves := []int{}

		for _, leaf := range leaves {
			for _, neighbor := range graph[leaf] {
				degrees[neighbor]--

				if degrees[neighbor] == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}

		leaves = newLeaves
	}

	return leaves
}
