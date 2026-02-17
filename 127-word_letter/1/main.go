package main

import (
	"fmt"
	"strings"
)

func main() {

}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// se a palavra não existir em wordList a resposta deve ser 0
	endFound := false
	for _, word := range wordList {
		if strings.EqualFold(word, endWord) {
			endFound = true
			break
		}
	}
	if !endFound {
		return 0
	}

	// o patern vai mapear padrões para palavras. O objetivo é tendo um padrão como "h*t"
	// encontrar todas as palavras que encaixam nele
	wordList = append(wordList, beginWord)
	patternMap := make(map[string][]string)
	for _, word := range wordList {
		for i := range word {
			pattern := fmt.Sprintf("%s%s%s", word[:i], "*", word[i+1:])
			patternMap[pattern] = append(patternMap[pattern], word)
		}
	}

	visited := make(map[string]bool)
	visited[beginWord] = true
	q := []string{beginWord}

	res := 1
	for len(q) > 0 {
		levelSize := len(q)

		for range levelSize {
			curr := q[0]
			q = q[1:]

			if curr == endWord {
				return res
			}

			// precisamos pegar as palavras que encaixam no pattern
			// para cada uma dessas palavras verificar se já foi visitada antes de adicionar a fila
			for _, pattern := range createPaterns(curr) {
				for _, word := range patternMap[pattern] {
					if !visited[word] {
						visited[word] = true
						q = append(q, word)
					}
				}
			}
		}

		res++
	}

	// nao tem transformaçao valida para a palavra.
	return 0
}

func createPaterns(w string) []string {
	words := []string{}
	for i := range w {
		words = append(words, fmt.Sprintf("%s%s%s", w[:i], "*", w[i+1:]))
	}
	return words
}
