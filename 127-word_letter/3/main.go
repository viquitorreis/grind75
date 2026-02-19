package main

import (
	"fmt"
	"strings"
)

func main() {

}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	patternsMap := make(map[string][]string)
	hasEnd := false
	for _, word := range wordList {
		if strings.EqualFold(word, endWord) {
			hasEnd = !hasEnd
		}

		for i := range word {
			pattern := fmt.Sprintf("%s%s%s", word[:i], "*", word[i+1:])
			patternsMap[pattern] = append(patternsMap[pattern], word)
		}
	}
	if !hasEnd {
		return 0
	}

	visited := make(map[string]bool)
	// precisamos partir de begin word e chegar no endword
	q := []string{beginWord}
	visited[beginWord] = true
	res := 1

	for len(q) > 0 {
		levelSize := len(q)

		for range levelSize {
			curr := q[0]
			q = q[1:]

			if curr == endWord {
				return res
			}

			for _, pattern := range createPatterns(curr) {
				for _, word := range patternsMap[pattern] {
					if !visited[word] {
						visited[word] = true
						q = append(q, word)
					}
				}
			}
		}

		res++
	}

	return 0
}

func createPatterns(word string) []string {
	w := []string{}
	for i := range word {
		w = append(w, fmt.Sprintf("%s%s%s", word[:i], "*", word[i+1:]))
	}
	return w
}
