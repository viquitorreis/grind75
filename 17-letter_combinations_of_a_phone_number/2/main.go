package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}

	for _, ch := range digits {
		ascii := ch - '0'

		if ascii < 2 || ascii > 9 {
			continue
		}

		res = permute(numToLetter(ascii), res)
	}

	return res
}

func numToLetter(n rune) []string {
	toLetters := map[rune][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}

	return toLetters[n]
}

func permute(s, old []string) []string {
	if len(old) == 0 {
		return s
	}

	newStr := []string{}
	for _, ch := range s {
		for _, char := range old {
			word := fmt.Sprintf("%s%s", string(char), string(ch))
			newStr = append(newStr, word)
		}
	}

	return newStr
}
