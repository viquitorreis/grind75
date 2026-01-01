package main

func main() {
	ransomNote := "aa"
	magazine := "aab"
	println("can?", canConstruct(ransomNote, magazine))
}

// return TRUE se ransomNote puder ser construida usando as eltras de magazine
// do contrario FALSE
func canConstruct(ransomNote string, magazine string) bool {
	// todas as letras em ransomNote devem estar contidas em magazine
	// precisa ter a quantia correta
	seenMag := make(map[byte]int, 0)
	seenRansom := make(map[byte]int, 0)
	maxLen := func(mg, rs int) int {
		if mg > rs {
			return mg
		}

		return rs
	}(len(magazine), len(ransomNote))

	for count := range maxLen {
		if count < len(magazine) {
			charMag := magazine[count]
			seenMag[charMag]++
		}

		var charRansom byte
		if count < len(ransomNote) {
			charRansom = ransomNote[count]
			seenRansom[charRansom]++
		}
	}

	count := 0

	for k, v := range seenRansom {
		val, ok := seenMag[k]
		if !ok {
			return false
		}

		if val < v {
			return false
		}

		count += v
	}

	return true && count == len(ransomNote)
}
