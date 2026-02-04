package main

func main() {

}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) || len(p) == 0 {
		return nil
	}

	count := [26]int{}
	for _, ch := range p {
		count[ch-'a']--
	}

	res := []int{}

	// left determina onde se inicia o anagrama
	left := 0
	for right, ch := range s {
		ascii := ch - 'a'
		count[ascii]++ // diminui a divida

		// movemos a janela, pois ainda estamos devendo o caractere na janela atual
		for count[ascii] > 0 {
			count[s[left]-'a']--
			left++
		}

		if right-left+1 == len(p) {
			res = append(res, left)
		}
	}

	return res
}
