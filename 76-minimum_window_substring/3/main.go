package main

import (
	"math"
	"strings"
)

func main() {

}

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	if len(s) == 1 && len(t) == 1 && strings.EqualFold(s, t) {
		return s
	}

	tFreq := make(map[byte]int)
	for _, char := range t {
		tFreq[byte(char)]++
	}

	// have só é somado quando quando achamos todas ocorrencias do char
	need, have := len(tFreq), 0

	minLen := math.MaxInt
	minStart := 0

	left := 0
	for right := 0; right < len(s); right++ {
		char := s[right]
		if _, ok := tFreq[char]; ok {
			tFreq[char]--
			if tFreq[char] == 0 {
				have++
			}
		}

		for have == need {
			// achamos uma janela menor, atualizamos minLen e minStart
			if right-left+1 < minLen {
				minLen = right - left + 1
				minStart = left
			}

			// quando removemos o caractere pela esquerda
			// precisamos reverter o que fizemos antes de adicionar ele (incrementar sua frequencia) para usar de novo.
			leftChar := s[left]
			if _, ok := tFreq[leftChar]; ok {
				if tFreq[leftChar] == 0 {
					have--
				}

				tFreq[leftChar]++
			}
			left++
		}
	}

	if minLen == math.MaxInt {
		return ""
	}

	return s[minStart : minStart+minLen]
}
