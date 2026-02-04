package main

import "fmt"

func main() {
	// fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func wordBreak(s string, wordDict []string) bool {
	t := NewTrie()
	for _, word := range wordDict {
		t.InsertWord(word)
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for start := 1; start <= len(s); start++ {
		for j := 0; j <= start; j++ {
			// precisamos verificar se dp[j] é verdadeiro (conseguimos formar tudo até a posição j)
			// e se  a substring s[j:start] está no dicionario
			if dp[j] && t.SearchWord(s[j:start]) {
				dp[start] = true
				fmt.Println("its word", string(s[j:start]), dp[j], dp[start-j])
			}
		}
	}

	fmt.Println("dp", dp)

	return dp[len(s)]
}

type Trie struct {
	Val      rune
	Children map[rune]*Trie
	IsLeaf   bool
}

func NewTrie() *Trie {
	return &Trie{
		Children: make(map[rune]*Trie),
	}
}

func (t *Trie) InsertWord(w string) {
	curr := t

	word := []string{}

	for _, char := range w {
		val, ok := curr.Children[char]
		if !ok {
			curr.Children[char] = &Trie{
				Val:      char,
				Children: make(map[rune]*Trie),
			}

			curr = curr.Children[char]

			word = append(word, string(char))

			continue
		}

		curr = val
	}

	fmt.Println("inserted: ", word)

	curr.IsLeaf = true
}

func (t *Trie) SearchWord(w string) bool {
	fmt.Println("searching:", w)
	curr := t

	for _, char := range w {
		val, ok := curr.Children[char]
		if !ok {
			return false
		}

		curr = val
	}

	return curr.IsLeaf
}
