package main

import (
	"fmt"
	"sort"
)

func main() {
	accs := [][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}

	fmt.Println(accountsMerge(accs))
}

func accountsMerge(accounts [][]string) [][]string {
	// esse problema é um union find
	// precisamos de um eleito para cada grupo, que será o representante da componente conexa
	// sabemos que se existe um primeiro email para a conta, aquele email é com certeza absuluta de uma pessoa unica...

	// 1. Precisamos montar o grafo
	// 		como o first email após o nome é único. Vamos mapear o first email para todos os outros emails
	// 		também vamos mapear o email para o nome do user
	graph := make(map[string][]string)
	emailToName := make(map[string]string)

	for _, acc := range accounts {
		name := acc[0]
		firstEmail := acc[1]

		// precisamos mapear todos os emails para o nome
		// e tambem construir o grafo de first email para os emails
		for _, email := range acc[1:] {
			graph[firstEmail] = append(graph[firstEmail], email)
			emailToName[email] = name

			// adicionando bi direcionalidade, precisamos ir do email para o first email também
			graph[email] = append(graph[email], firstEmail)
		}
	}

	visited := make(map[string]bool)
	res := [][]string{}

	// iteramos nos email to name

	for email, name := range emailToName {
		if !visited[email] {
			emails := []string{}

			dfs(graph, visited, email, &emails)

			sort.Strings(emails)

			res = append(res, append([]string{name}, emails...))
		}
	}

	return res
}

func dfs(graph map[string][]string, visited map[string]bool, email string, emails *[]string) {
	if visited[email] {
		return
	}

	visited[email] = true
	*emails = append(*emails, email)

	for _, email := range graph[email] {
		dfs(graph, visited, email, emails)
	}
}
