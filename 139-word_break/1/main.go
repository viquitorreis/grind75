package main

import "fmt"

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

type Trie struct {
	val      rune
	children []*Trie
	isLeaf   bool
}

func newTrie() Trie {
	return Trie{
		children: make([]*Trie, 26),
	}
}

func (t *Trie) insert(s string) {
	if s == "" {
		return
	}

	// precisamos armazenar o root, vamos atualizando a medida que andamos
	curr := t

	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{
				val:      rune(s[i]),
				children: make([]*Trie, 26),
			}
		}

		curr = curr.children[idx]
	}

	curr.isLeaf = true
}

func (t *Trie) searchWord(s string) bool {
	curr := t

	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		if curr.children[idx] == nil {
			return false
		}

		curr = curr.children[idx]
	}

	return curr.isLeaf
}

func wordBreak(s string, wordDict []string) bool {
	trie := newTrie()

	// constroi a trie com todas as palavras do dict
	for _, word := range wordDict {
		trie.insert(word)
	}

	// DP
	// vamos tentar criar um array on a posição i significa que consigo formar a substring de zero até i
	// precisamos de len + 1, pois a posição de inicio 0 - é o base case, onde não temos nada
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// se conseguirmos formar até j E a palavra de j até i existe no trie
			// então consigo formar até i, logo a palavra existe
			// dp[j] -> garante que chegamos até a posição j usando palavras válidas
			if dp[j] && trie.searchWord(s[j:i]) {
				dp[i] = true
				break
			}
		}
	}

	// ao final a resposta está no dp na ultima posição, que representa a string inteira
	return dp[n]
}
