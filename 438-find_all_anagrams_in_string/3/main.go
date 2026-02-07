package main

func main() {

}

func findAnagrams(s string, p string) []int {
	res := []int{}
	if len(s) < len(p) || len(p) == 0 {
		return nil
	}

	debth := [26]int{}
	for _, char := range p {
		debth[char-'a']--
	}

	left := 0
	for right, char := range s {
		ascii := char - 'a'
		debth[ascii]++

		for debth[ascii] > 0 {
			debth[s[left]-'a']--
			left++
		}

		if right-left+1 == len(p) {
			res = append(res, left)
		}
	}

	return res
}
