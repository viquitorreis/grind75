package main

import "fmt"

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
}

type Trie struct {
	val      rune
	children map[rune]*Trie
	isLeaf   bool
}

func Constructor() Trie {
	return Trie{
		children: make(map[rune]*Trie),
		isLeaf:   false,
	}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	// 1. pega ref da head
	curr := this

	for _, char := range word {
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

func (this *Trie) Search(word string) bool {
	curr := this

	for _, char := range word {
		val, ok := curr.children[char]
		if !ok {
			return false
		}

		curr = val
	}

	return curr.isLeaf
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this

	for _, char := range prefix {
		val, ok := curr.children[char]
		if !ok {
			return false
		}

		curr = val
	}

	return true
}
