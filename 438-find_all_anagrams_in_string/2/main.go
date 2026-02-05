package main

func main() {

}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) || len(p) == 0 {
		return nil
	}

	toForm := [26]int{}
	for _, char := range p {
		ascii := char - 'a'
		toForm[ascii]--
	}

	res := []int{}
	left := 0
	for right, char := range s {
		ascii := char - 'a'

		toForm[ascii]++

		for toForm[ascii] > 0 {
			// ainda tem mais da letra atual na janela, portanto precisamos aumentar a janela de busca...
			// logo esse left precisa ir para frente pois ainda não forma anagrama
			// então vamos retornar a dívida para tentar em outra janela
			toForm[s[left]-'a']--
			left++
		}

		if right-left+1 == len(p) {
			res = append(res, left)
		}
	}

	return res
}
