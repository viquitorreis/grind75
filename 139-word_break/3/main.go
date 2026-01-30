package main

func main() {

}

func wordBreak(s string, wordDict []string) bool {
	t := NewTrie()

	for _, word := range wordDict {
		t.Insert(word)
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// precisamos testar todas as formas de criar uma palavra até i
			// 1. dp[j] = conseguiu formar a string até j usando palavras válidas
			// 2. t.SearchWord(s[j:i]) = a palavra de j até i existe no dicionário
			if dp[j] && t.SearchWord(s[j:i]) {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

type Trie struct {
	Val      rune
	children map[rune]*Trie
	isLeaf   bool
}

func NewTrie() *Trie {
	return &Trie{
		children: make(map[rune]*Trie),
	}
}

func (t *Trie) Insert(w string) {
	curr := t

	for _, char := range w {
		val, ok := curr.children[char]
		if !ok {
			curr.children[char] = &Trie{
				Val:      char,
				children: make(map[rune]*Trie),
			}

			curr = curr.children[char]
			continue
		}

		curr = val
	}

	curr.isLeaf = true
}

func (t *Trie) SearchWord(w string) bool {
	curr := t

	for _, char := range w {
		val, ok := curr.children[char]
		if !ok {
			return false
		}

		curr = val
	}

	return curr.isLeaf
}
