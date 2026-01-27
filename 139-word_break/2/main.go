package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

func wordBreak(s string, wordDict []string) bool {
	trie := NewTrie()
	for _, word := range wordDict {
		trie.Insert(word)
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// j até i, precisamos testar TODAS as formas possíveis de quebrar a string até a posição i
			// precisamos verificar na posição j (anterior) se temos uma palavra válida, até então..
			if dp[j] && trie.SearchWord(s[j:i]) {
				dp[i] = true
				break
			}
		}
	}

	// string inteira
	return dp[len(s)]
}

type Trie struct {
	val      rune
	children map[rune]*Trie
	isLeaf   bool
}

func NewTrie() *Trie {
	return &Trie{
		children: make(map[rune]*Trie, 26),
	}
}

func (t *Trie) Insert(word string) {
	curr := t

	for _, char := range word {
		// analisar cada char
		val, ok := curr.children[char]
		if !ok {
			curr.children[char] = &Trie{
				val:      char,
				children: make(map[rune]*Trie),
			}

			curr = curr.children[char]

			continue
		}

		curr = val
	}

	curr.isLeaf = true
}

func (t *Trie) SearchWord(word string) bool {
	curr := t

	for _, char := range word {
		val, ok := curr.children[char]
		if !ok {
			return false
		}

		curr = val
	}

	return curr.isLeaf

}
