package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

// SLIDING WINDOW

// BIG O: O(n)

// Space: O(n)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	chars := make(map[byte]bool)
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		// 1. enquanto s[right] já existe no chars map
		//		- removemos s[left] do chars map
		//		- avançamos left++ -> caso contrário left (onde começa) fica no mesmo lugar
		for {
			if _, ok := chars[s[right]]; ok {
				delete(chars, s[left])
				left++
				continue
			}

			break
		}

		// 2. adicionamos s[right] ao chars map
		chars[s[right]] = true

		// 3. atualizar maxLen semrpe que necessario
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
