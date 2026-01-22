package main

import (
	"fmt"
	"sort"
)

func main() {
	s := [][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}
	fmt.Println(accountsMerge(s))
}

func accountsMerge(accounts [][]string) [][]string {
	graph := make(map[string][]string)
	emailToName := make(map[string]string)

	for _, acc := range accounts {
		name := acc[0]
		firstEmail := acc[1]

		for i := 1; i < len(acc); i++ {
			email := acc[i]
			emailToName[email] = name

			graph[firstEmail] = append(graph[firstEmail], email)
			graph[email] = append(graph[email], firstEmail)
		}
	}

	visited := make(map[string]bool)
	result := [][]string{}

	for email := range emailToName {
		if !visited[email] {
			emails := []string{}
			dfs(email, graph, visited, &emails)

			sort.Strings(emails)
			result = append(result, append([]string{emailToName[email]}, emails...))
		}
	}

	return result
}

func dfs(email string, graph map[string][]string, visited map[string]bool, emails *[]string) {
	if visited[email] {
		return
	}

	visited[email] = true
	*emails = append(*emails, email)

	for _, neighbor := range graph[email] {
		dfs(neighbor, graph, visited, emails)
	}
}
