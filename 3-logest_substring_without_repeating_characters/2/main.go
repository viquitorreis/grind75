package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	// sliding window
	if len(s) == 0 {
		return 0
	}

	elements := make(map[byte]bool)
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		for {
			// sempre que achar um elemento repetido em right, removemos o elementos da esquerda
			if _, ok := elements[s[right]]; ok {
				delete(elements, s[left])
				left++
				continue
			}

			break
		}

		elements[s[right]] = true
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
