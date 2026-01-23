package main

import "sort"

func main() {

}

func accountsMerge(accounts [][]string) [][]string {
	// precisamos fazer um grafo de componentes conexas

	// firstEmail -> todos emails
	// nesse problema é a unica forma de garantir unicidade
	// quando tem o primeiro email para o nome, garantimos que é unico dele
	graph := make(map[string][]string)
	emailToname := make(map[string]string)

	for _, acc := range accounts {
		name := acc[0]
		firstEmail := acc[1] // precisamos do first email para mapear corretamente

		for i := 1; i < len(acc); i++ {
			email := acc[i]
			emailToname[email] = name

			graph[firstEmail] = append(graph[firstEmail], email)
			// dfs -> precisamos disso para navegar nos dois sentidos, pois
			// de firstEmail conseguimos chegar em email
			// mas de email não vamos conseguir voltar para firstEmail.
			// Ex: graph[john1] = [john2, john3]
			// graph[john2] = [john1]
			// graph[john3] = [john1]
			// assim conseguimos voltar para o primeiro email, a partir de qualquer elemento
			graph[email] = append(graph[email], firstEmail)
		}
	}

	visited := make(map[string]bool)
	result := [][]string{}

	for email := range emailToname {
		if !visited[email] {
			res := []string{}

			dfs(email, graph, &res, visited)

			sort.Strings(res)

			result = append(result, append([]string{emailToname[email]}, res...))
		}
	}

	return result
}

func dfs(email string, graph map[string][]string, emails *[]string, visited map[string]bool) {
	if visited[email] {
		return
	}

	*emails = append(*emails, email)
	visited[email] = true

	for _, neighbor := range graph[email] {
		dfs(neighbor, graph, emails, visited)
	}
}
