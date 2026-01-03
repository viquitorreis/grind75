package main

import "fmt"

func main() {
	s := "abcd"
	k := 2
	fmt.Println(reversePrefix(s, k))
}

func reversePrefix(s string, k int) string {
	str := make([]byte, len(s))

	for i := 0; i < k; i++ {
		str[i] = s[k-1-i]
	}

	for i := k; i < len(s); i++ {
		str[i] = s[i]
	}

	return string(str)
}
