package main

import "sort"

func main() {

}

func accountsMerge(accounts [][]string) [][]string {
	res := [][]string{}
	if len(accounts) == 0 || len(accounts[0]) == 0 {
		return res
	}

	graph := make(map[string][]string)
	emailToName := make(map[string]string)

	for _, acc := range accounts {
		name := acc[0]
		firstEmail := acc[1]
		accs := acc[1:]

		for _, email := range accs {
			graph[firstEmail] = append(graph[firstEmail], email)
			graph[email] = append(graph[email], firstEmail)

			emailToName[email] = name
		}
	}

	visited := make(map[string]bool)

	for email, name := range emailToName {
		if !visited[email] {
			emails := []string{}

			dfs(graph, &emails, visited, email)

			sort.Strings(emails)

			res = append(res, append([]string{name}, emails...))
		}
	}

	return res
}

func dfs(graph map[string][]string, emails *[]string, visited map[string]bool, email string) {
	if visited[email] {
		return
	}

	*emails = append(*emails, email)
	visited[email] = true

	for _, mail := range graph[email] {
		dfs(graph, emails, visited, mail)
	}
}
