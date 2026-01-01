package main

func main() {
	// str1 := "ab#c"
	// str2 := "ad#c"
	// println(backspaceCompare(str1, str2))

	// str3 := "ab##"
	// str4 := "c#d#"
	// println(backspaceCompare(str3, str4))

	// str5 := "a#c"
	// str6 := "b"
	// println(backspaceCompare(str5, str6))

	str7 := "a##c"
	str8 := "#a#c"
	println(backspaceCompare(str7, str8))
}

// COMPLEXITY ANALYSIS
// Big O: O(n) - linear
//		O(n) - being n the biggest string from both. Since we are iterating on it only one time.
//			All of other ones are only conditionals, variables which takes O(1)
// Space Complexity: O(n) - being n the biggest string
//		O(n) since we are allocating at most the biggest string - asymptotically
//			technically its O(len(s) + len(t)), but we are simplyfying to represent the length of the bigger string

// # -> remove last character
func backspaceCompare(s string, t string) bool {
	biggestLen := 0
	if len(s) > len(t) {
		biggestLen = len(s)
	} else {
		biggestLen = len(t)
	}

	sWord := []byte{}
	tWord := []byte{}

	for i := 0; i < biggestLen; i++ {
		if len(s) > i {
			if s[i] == '#' && len(sWord) > 0 {
				sWord = sWord[:len(sWord)-1]
			} else if s[i] != '#' {
				sWord = append(sWord, s[i])
			}
		}

		if len(t) > i {
			if t[i] == '#' && len(tWord) > 0 {
				tWord = tWord[:len(tWord)-1]
			} else if t[i] != '#' {
				tWord = append(tWord, t[i])
			}
		}
	}

	return string(sWord) == string(tWord)
}
