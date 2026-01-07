package main

func main() {
	s := "aacc"
	t := "ccac"
	println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	seenT := make(map[byte]int)
	seenS := make(map[byte]int)

	for sIdx, tIdx := 0, 0; ; sIdx, tIdx = sIdx+1, tIdx+1 {
		if sIdx < len(s) {
			seenS[s[sIdx]]++
		}

		if tIdx < len(t) {
			seenT[t[tIdx]]++
		}

		if sIdx > len(s) && tIdx > len(t) {
			break
		}
	}

	for k, v := range seenS {
		val, ok := seenT[k]
		if !ok {
			return false
		}

		if v != val {
			return false
		}
	}

	return true
}
