package main

func main() {
	str := "abcd"
	k := 2
	println(reversePrefix(str, k))
}

func reversePrefix(s string, k int) string {
	str := make([]byte, len(s))
	for i := range k {
		str[i] = s[k-1-i]
	}

	for i := k; i < len(s); i++ {
		str[i] = s[i]
	}

	return string(str)
}
