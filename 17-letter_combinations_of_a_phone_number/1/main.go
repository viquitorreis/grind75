package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

func digitLetters(d rune) []string {
	combs := map[rune][]string{}
	combs['2'] = []string{"a", "b", "c"}
	combs['3'] = []string{"d", "e", "f"}
	combs['4'] = []string{"g", "h", "i"}
	combs['5'] = []string{"j", "k", "l"}
	combs['6'] = []string{"m", "n", "o"}
	combs['7'] = []string{"p", "q", "r", "s"}
	combs['8'] = []string{"t", "u", "v"}
	combs['9'] = []string{"w", "x", "y", "z"}
	return combs[d]
}

func permute(old, r []string) []string {
	newStr := []string{}
	if len(old) == 0 {
		for _, ch := range r {
			newStr = append(newStr, ch)
		}

		return newStr
	}

	for _, char := range old {
		for _, ch := range r {
			comb := string(char) + string(ch)
			newStr = append(newStr, comb)
		}
	}

	return newStr
}

func letterCombinations(digits string) []string {
	combs := []string{}
	for _, digit := range digits {
		ascii := digit - '0'
		if ascii < 0 || ascii > 9 {
			continue
		}

		combs = permute(combs, digitLetters(digit))
	}
	return combs
}
