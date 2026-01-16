package main

import "fmt"

func main() {
	newTrie := Constructor()
	newTrie.Insert("apple")
	fmt.Println("word exists?", newTrie.Search("apple"))
	fmt.Println("is prefix? ", newTrie.StartsWith("a"))
}

type Trie struct {
	val      rune
	children []*Trie
	isLeaf   bool
}

func Constructor() Trie {
	return Trie{
		children: make([]*Trie, 26),
	}
}

// Big O: O(n) -> n = len(word)
// Space: O(n) -> n = len(word)
func (this *Trie) Insert(word string) {
	if word == "" {
		return
	}

	// pegamos referencia da raiz
	// isso é muito importante, pois a cada iteração, vamos inserir em cima do nó atual
	// criando a palavra
	curr := this

	for i := 0; i < len(word); i++ {
		// idx will be its value in ASCII
		// we need to check if the node exists for the current character in the trie
		idx := int(word[i] - 'a')
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{
				val:      rune(word[i]),
				children: make([]*Trie, 26),
			}
		}

		curr = curr.children[idx]
	}

	curr.isLeaf = true
}

// Big O: O(n) -> sendo n o tamanho da palavra
// Space Complexity -> O(1)
func (this *Trie) Search(word string) bool {
	// importante apontar o current para a raiz
	// vamos atualizar ele, sempre que acharmos um char que dê match na árvore, dessa forma montando a palavra
	// ao final, se acharmos todas buscadas retornamos se é true
	curr := this

	for i := 0; i < len(word); i++ {
		// checa se o node existe para o caractere atual
		idx := int(word[i] - 'a')
		if curr.children[idx] == nil {
			return false
		}

		// movemos o ponteiro curr para o node com o caractere já existente para o caractere existente
		curr = curr.children[idx]
	}

	// se a palavra existir retornamos true
	return curr.isLeaf
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this

	for i := 0; i < len(prefix); i++ {
		// checa se o node existe para o caractere atual
		idx := int(prefix[i] - 'a')
		if curr.children[idx] == nil {
			return false
		}

		// movemos o ponteiro curr para o node com o caractere já existente para o caractere existente
		curr = curr.children[idx]
	}

	return true
}
