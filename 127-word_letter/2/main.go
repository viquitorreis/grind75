package main

import (
	"fmt"
	"strings"
)

func main() {

}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	hasEnd := false
	for _, word := range wordList {
		if strings.EqualFold(endWord, word) {
			hasEnd = true
			break
		}
	}
	if !hasEnd {
		return 0
	}

	wordList = append(wordList, beginWord)
	patternMap := make(map[string][]string)
	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			pattern := fmt.Sprintf("%s%s%s", word[:i], "*", word[i+1:])
			patternMap[pattern] = append(patternMap[pattern], word)
		}
	}

	visited := make(map[string]bool)
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

	return 0
}

func createPatterns(w string) []string {
	patterns := []string{}
	for i := range w {
		pattern := fmt.Sprintf("%s%s%s", w[:i], "*", w[i+1:])
		patterns = append(patterns, pattern)
	}
	return patterns
}
